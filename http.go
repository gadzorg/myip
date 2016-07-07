package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, r.RemoteAddr)
 	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}

