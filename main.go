// Minimal Client/Server AJAX Communication using golang web-server and JQuery
// Visit: http://127.0.0.1:8088
package main

import (
    "fmt"
    "html/template"
    "net/http"
    "path"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    data := Data{"World!"}

    fp := path.Join("templates", "ajax.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
       data := r.FormValue("name")
       fmt.Fprintf(w, "Hi, %s!", data)
    }
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/ajax", ajaxHandler)
    http.ListenAndServe(":8088", nil)
}
