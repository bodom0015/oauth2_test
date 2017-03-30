// Simple http request dumper to test nginx/oauth2_proxy/upstream request handling
package main

// Stdlib
import "fmt"
import "log"
import "flag"
import "time"
import "net/http"
import "net/http/httputil"

func RequestDumpHandler(resp http.ResponseWriter, req *http.Request) {
    request_dump, err := httputil.DumpRequest(req, true)

    t := time.Now()
    fmt.Printf("[%s] %s, with dump:\n", t.Format(time.StampMilli), req.URL.String())

    if err == nil {
        fmt.Printf("---\n")
        fmt.Printf("%s", request_dump)
        fmt.Printf("---\n")
        resp.Write([]byte("Request received\n"))
    } else {
        fmt.Printf("Request unable to be output\n")
        resp.Write([]byte("Could not handle request\n"))
    }
}

func main() {
    addr := flag.String("address", "", "the address to bind to. Default is 0.0.0.0")
    port := flag.String("port", "8080", "the local port to bind to")

    flag.Parse()

    http.HandleFunc("/", RequestDumpHandler)

    fmt.Printf("httpecho starting up on %s:%s...\n", *addr, *port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *addr, *port), nil))
}
