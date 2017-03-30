package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/gorilla/securecookie"
    "log"
    "net/http"
    "net/http/httputil"
)

var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))

var router = mux.NewRouter()

func main() {
    router.HandleFunc("/cauth/auth", checkHandler).Methods("GET")
    router.HandleFunc("/cauth/sign_in", indexPageHandler)
    router.HandleFunc("/cauth/login", loginHandler).Methods("POST")
    router.HandleFunc("/cauth/logout", logoutHandler).Methods("GET")

    http.Handle("/", router)
    log.Print("Listening on 8081")
    http.ListenAndServe(":8081", nil)
}

const indexPage = `
 <h1>Login</h1>
 <form method="post" action="/cauth/login">
     <label for="name">User name</label>
     <input type="text" id="name" name="name">
     <label for="password">Password</label>
     <input type="password" id="password" name="password">
     <button type="submit">Login</button>
 </form>
 `

func dump(req *http.Request) {
    log.Print("dump")
    request_dump, _ := httputil.DumpRequest(req, true)
    log.Print(string(request_dump))
}

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
    log.Print("indexHandler")
    dump(request)
    fmt.Fprintf(response, indexPage)
}

func checkHandler(response http.ResponseWriter, request *http.Request) {
    log.Print("checkHandler")
    dump(request)
    username := getUserName(request)
    if username != "" {
        log.Print("checkHandler already logged in")
        log.Print("checkHandler returning 200")
        response.WriteHeader(http.StatusOK)
        return
    }
    log.Print("checkHandler returning 401")
    response.WriteHeader(http.StatusUnauthorized)
    return
}

func loginHandler(response http.ResponseWriter, request *http.Request) {
    log.Print("loginHandler")
    dump(request)

    //fmt.Fprintf(response, indexPage)
    name := request.FormValue("name")
    pass := request.FormValue("password")
    if name == "test" && pass == "200" {

        log.Print("login OK")
        setSession(name, response)
        //target := request.Header.Get("X-Target")
        //log.Print(target)
        log.Print("loginHandler redirecting to /")
        http.Redirect(response, request, "/", http.StatusFound)
        return
    } else if name == "test" && pass == "403" {
        log.Print("login OK, but unauthorized")
        log.Print("loginHandler returning 403")
        response.WriteHeader(http.StatusForbidden)
        return
    }
    log.Print("loginHandler returning 401")
    response.WriteHeader(http.StatusUnauthorized)
    return
}

func logoutHandler(response http.ResponseWriter, request *http.Request) {
    log.Print("logoutHandler")
    dump(request)
    clearSession(response)
    response.WriteHeader(http.StatusOK)
    return
}

func setSession(userName string, response http.ResponseWriter) {
    value := map[string]string{
        "name": userName,
    }
    if encoded, err := cookieHandler.Encode("session", value); err == nil {
        cookie := &http.Cookie{
            Name:   "session",
            Value:  encoded,
            Domain: ".ndslabs.org",
            Path:   "/",
        }
        http.SetCookie(response, cookie)
    } else {
        log.Print(err.Error())
    }
}

func getUserName(request *http.Request) (userName string) {
    if cookie, err := request.Cookie("session"); err == nil {
        cookieValue := make(map[string]string)
        if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
            userName = cookieValue["name"]
            log.Print(userName)
        } else {
            log.Print(err.Error())
        }
    } else {
        log.Print(err.Error())
    }
    return userName
}

func clearSession(response http.ResponseWriter) {
    cookie := &http.Cookie{
        Name:   "session",
        Value:  "",
        Domain: ".ndslabs.org",
        Path:   "/",
        MaxAge: -1,
    }
    http.SetCookie(response, cookie)
}
