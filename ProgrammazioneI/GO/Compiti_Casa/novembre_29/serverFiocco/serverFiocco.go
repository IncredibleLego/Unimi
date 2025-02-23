//Autore: Francesco Corrado

package main

import(
  "fmt"
  "net/http" //Pacchetto per gestire pagine web
)

func homePage(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(`
    <!doctype html>
      <title>HomePage</title>
      <h1>Benvenuto</h1>
      <p>Seleziona la pagina alla quale vuoi andare`))
}

func CFUPage(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(`
    <!doctype html>
      <title>WARNING</title>
      <h1>BRO CHE FAI</h1>
      <h3>Stai provando ad avere CFU gratis? Non si può bisogna studiare</h3>
      <p<Ti meriti questo rickroll`))
}

func main(){
  http.HandleFunc("/home", homePage)
  http.HandleFunc("/CFUgratis", CFUPage)
  fmt.Println("Listening on http://localhost:400/")
  http.ListenAndServe(":4000", nil)
}