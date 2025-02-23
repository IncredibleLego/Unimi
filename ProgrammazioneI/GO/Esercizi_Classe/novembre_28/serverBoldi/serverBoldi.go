package main

import (
  "fmt"
  "net/http"
)

func pippoPage(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(`
          <!doctype html>
            <title>Pippo</title>
            <h1>Pippo</h1>

            <p>Ciao sono Pippo!

            <p><img src="https://cialdecaffecapsule.files.wordpress.com/2013/01/elefante.jpg?w=240">`))
}

func plutoPage(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(`
          <!doctype html>
            <title>Pluto</title>
            <h1>Pluto</h1>

            <p>Ciao sono Pluto!

            <p><img src="https://citynews-bresciatoday.stgy.ovh/~media/horizontal-hi/68870303716911/ificmwo-2.jpeg">`))
}

func main() {
  http.HandleFunc("/pippo", pippoPage)
  http.HandleFunc("/pluto", plutoPage)
  fmt.Println("Listening on http://localhost:3000/")
  http.ListenAndServe(":3000", nil)
}
