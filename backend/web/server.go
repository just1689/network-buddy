package web

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"network-buddy/backend/metrics"
	"os"
	"time"
)

type Event struct {
	Operation string `json:"operation"`
	Value     string `json:"value"`
}
type Result struct {
	Body  string `json:"body"`
	Error string `json:"error"`
}

func ListenAndServe(listen string) {
	http.HandleFunc("/network-buddy/api", func(writer http.ResponseWriter, request *http.Request) {
		logrus.Println("+>>", request.URL.Path)
		metrics.HTTPCalls.Inc()
		b, _ := ioutil.ReadAll(request.Body)
		e := &Event{}
		err := json.Unmarshal(b, e)
		if err != nil {
			logrus.Errorln(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if e.Operation == "lookup" {
			logrus.Infoln("> Going to look up hostname:", e.Operation)
			addresses, err := net.LookupHost(e.Value)
			result := Result{}
			if err != nil {
				result.Error = err.Error()
				b, _ = json.Marshal(result)
				writer.Write(b)
				return
			}
			for _, a := range addresses {
				if result.Body != "" {
					result.Body += ", "
				}
				result.Body += a
			}
			b, _ = json.Marshal(result)
			writer.Write(b)
			return

		} else if e.Operation == "probe" {
			logrus.Infoln("> Going to probe address:", e.Operation)
			open := tcpTest(e.Value)
			result := Result{}
			if open {
				result.Body = "open"
			} else {
				result.Error = "closed"
			}
			b, _ = json.Marshal(result)
			writer.Write(b)
			return

		}

		fmt.Println(string(b))
		writer.Write([]byte("result from server!"))
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		logrus.Println(">>>", request.URL.Path)
		metrics.HTTPCalls.Inc()
		getStaticURLRP().ServeHTTP(writer, request)
	})
	http.Handle("/metrics", promhttp.Handler())
	logrus.Infoln("Starting server on ", listen)
	http.ListenAndServe(listen, nil)
}

func getStaticURL() string {
	staticURL := os.Getenv("staticURL")
	if staticURL == "" {
		staticURL = "http://localhost:4200"
	}
	return staticURL
}

func getStaticURLRP() *httputil.ReverseProxy {
	u, _ := url.Parse(getStaticURL())
	rp := httputil.NewSingleHostReverseProxy(u)
	return rp
}

func tcpTest(address string) bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Println("Connecting error:", err)
	}
	if conn != nil {
		defer conn.Close()
		return true
	}
	return false
}
