# oauth2_test

Download oauth2_proxy binary 

```
./oauth2_proxy --config oauth2_proxy.cfg

docker run --net=host -p 80:80 -v `pwd`/nginx/default.conf:/etc/nginx/conf.d/default.conf nginx

docker run -d -p 8081:8080 gcr.io/google_containers/echoserver:1.4
```


Notes:
* The main problem is we want the upstream to receive the oauth user information (email, name, token, etc). The only way I can get this to work is to use proxy_pass to the oauth proxy for each location (in this example /).
* Maybe we can find a way to override auth_request to handle both authentication and authorization
* Otherwise, if we need to use proxy_pass, this breaks the current Kubernetes ingress model -- so we'll need to customize.
