// tlsFileServer is simple HTTP/TLS FileServer, written in go.
// 
//  tlsFileServer.go -port=8443 -root=/usr/local/share/
// above command will serve the given root-dir (/usr/local/share/) 
// on port 8443 via https.
// 
// use `generate_cert.go` to get `cert.pem` and `key.pem` like this:
//  go run generate_cert.go -host=127.0.0.1,localhost
//
// tlsFileServer is written by: Florian 'scusi' Walther, <flw@posteo.de>
//
package main

import (
    "flag"
    "fmt"
    "net/http"
    _ "os"
    "log"
    _ "expvar"
)

// global variables must declared here...
var port string
var root string
var cert string
var key string

// init function initializes our command line flags with sane defaults
func init() {
    flag.StringVar(&port, "port", "8443", "port to listen on")
    flag.StringVar(&root, "root", "/usr/local/share", "root directory to serve files from")
    flag.StringVar(&cert, "cert", "cert.pem", "pem encoded cert to use")
    flag.StringVar(&key, "key",   "key.pem", "pem encoded key to use")
}

// main function parses flags, print flags and starts the server,
// fatal errors will be logged to STDOUT
func main() {
    flag.Parse()
    fmt.Println("port has value ", port)
    fmt.Println("root has value ", root)
    fmt.Println("cert has value ", cert)
    fmt.Println("key has value ", key)
    // todo: proper input validation
    // BUG(scusi) input validation missing
    path := root
    fmt.Printf("Listening on port '%s', serving directory '%s'\n", port, path)
    log.Fatal(http.ListenAndServeTLS(":" + port, cert, key, http.FileServer(http.Dir(path))))
}
