FROM golang:1.14 as builder
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM scratch as final
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /src/app /app
USER 65534:65534
ENTRYPOINT ["/app"]