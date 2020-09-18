package main

import (
	"fmt"
)

import (
	"os"
 	"time"
//     "log"
//     "net/http"
    model "github.com/evaizee/go-docker/models"
)

func main() {
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    // })

    // http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
    //     fmt.Fprintf(w, "Hi")
    // })

    // http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request){
    //     fmt.Fprintf(w, "hansa rostock")
    // })

    // http.HandleFunc("/rust", func(w http.ResponseWriter, r *http.Request){
    //     fmt.Fprintf(w, "hansa bastard")
    // })

    // http.HandleFunc("/getMatches", func(w http.ResponseWriter, r *http.Request){
    //     model.GetByLeagueAndDate(524, time.Now().Format("2006-11-22"), "Asia/Jakarta")
    // })

    // log.Fatal(http.ListenAndServe(":9100", nil))
    fmt.Println("March On")
    if len(os.Args) >= 2 {
        args := os.Args
        if args[1] == "rest" {
            model.GetByLeagueAndDate(2790, time.Now().Format("2006-01-02"), "Asia/Jakarta")
        } else {
            fmt.Println("bas bas")
        }
    }
}