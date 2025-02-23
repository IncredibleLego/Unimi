//Autore: Francesco Corrado

package main

import(
	"fmt"
	"net/http"
	"math"
	_"os"
	"image"
	"image/png"
	"github.com/holizz/terrapin"
)

func pippoPage(w http.ResponseWriter, r *http.Request) { //La richiesta deve essere fatta così
//	w.Write([]byte("<!doctype html>\n<title>Pippo</title>\n<h1>Pippo</h1>\n<p>Ciao sono Pippo!"))
	w.Write([]byte(`
		<!doctype html>
		 	<title>Pippo</title>
		 	<h1>Pippo</h1>
		 	<p>Ciao sono Pippo2!
			<p> <img src="https://cialdecaffecapsule.files.wordpress.com/2013/01/elefante.jpg" alt="Mia Immagine">
		`))
}

func plutoPage(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<!doctype html>
				 <title>Pluto</title>
				 <h1>Pluto</h1>
				 <p>Ciao sono Pluto
				<p> <img src="http://localhost:3000/nostraimmagine.png">
			`))
}

func imgPage(w http.ResponseWriter, r *http.Request) {
//	f, _:=os.Open("/home/francesco/Documenti/GO/Esercizi_Classe/28_novembre/serverHTML/immagine.png")
	image:= image.NewRGBA(image.Rect(0, 0, 500, 500))
	t := terrapin.NewTerrapin(image, terrapin.Position{250.0, 450.0})
	t.Forward(100)
	t.Right(math.Pi/2)
	t.Forward(100)
	t.Right(math.Pi/2 + math.Pi/4)
	t.Forward(math.Sqrt(100*100+100*100))
	png.Encode(w, image)
}

func natalePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<!doctype html>
			 <title>Natale</title>
			 <h1>Natale</h1>
			 <p>Buon Natale con il fiocco di Koch!
			<p> <img src="http://localhost:3000/curvaKoch.png">
		`))
}

/* Questa funzione usa la tartaruga t per disegnare una curva di koch
   di livello dato, con segmenti di lung */


func koch(t *terrapin.Terrapin, lung float64, livello int) {
	if livello == 0{
		t.Forward(lung)
	}else{
		koch(t, lung, livello-1)
		t.Left(math.Pi/ 3.0)
		koch(t, lung, livello-1)
		t.Right(math.Pi - mat)
	}
}

func curvaKoch(w http.ResponseWriter, r *http.Request) {
		image:= image.NewRGBA(image.Rect(0, 0, 500, 500))
		t := terrapin.NewTerrapin(image, terrapin.Position{250.0, 250.0})
		koch(t, 10.0, 3)
		png.Encode(w, image)
	}



func main(){
	http.HandleFunc("/pippo", pippoPage) //Se la richiesta è /pippo, rispondi con la funzione pippoPage
	http.HandleFunc("/pluto", plutoPage)
	http.HandleFunc("/nostraimmagine.png", imgPage)
	http.HandleFunc("/natale", natalePage)
	http.HandleFunc("/curvaKoch", curvaKoch)
	fmt.Println("Listening on http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}