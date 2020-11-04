import {Component} from '@angular/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Network Buddy';

  domainName = "internal-ingress.cluster-infra.svc.cluster.local";
  result = "";
  error = "";
  port = 8080;
  loading = false;
  ngOnInit(): void {

  }

  public nsLookup(): void {
    this.setResult("");
    this.error = "";
    this.loading = true;
    setTimeout(() => {
      this.remoteCall("lookup", this.domainName);
    }, 500)
  }

  public probe(): void {
    this.setResult("");
    this.error = "";
    this.loading = true;
    setTimeout(() => {
      this.remoteCall("probe", this.domainName + ":" + this.port);
    }, 500)
  }

  public remoteCall(o, v): void {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/network-buddy/api", true);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.send(JSON.stringify({
      operation: o,
      value: v,
    }));
    xhr.onload = () => {
      const data = xhr.responseText;
      let result = JSON.parse(data);
      this.setResult(result.body);
      this.error = result.error;
      this.loading = false;
    }
  }

  setResult(s: string): void {
    this.result = s;
  }

}
