package main
import (
        "fmt"
        "log"
        "net/http"
        "os"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", hello)
    port := os.Getenv("PORT")
    if port == "" {
            port = "8080"
}
log.Printf("Server Listining on port %s" , port)
log.Fatal(http.ListenAndServe(":"+port, mux))
}
func hello (w http.ResponseWriter, r *http.Request) {
        log.Printf("Serving request:  %s" , r.URL.Path)
        host, _ := os.Hostname()
        fmt.Fprintf(w, "Hello, World!\n")
        fmt.Fprintf(w, "Version: 1.0.0\n")
        fmt.Fprintf(w, "Hostname: %s\n", host)

}
