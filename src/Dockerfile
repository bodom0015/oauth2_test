FROM alpine:latest

RUN  apk update  && apk add wget 
ADD  https://github.com/craig-willis/oauth2_test/releases/download/v0.1/login-amd64-linux /
RUN chmod +x /login-amd64-linux


CMD ["/login-amd64-linux"]
