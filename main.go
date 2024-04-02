package main

import (
	"fmt"
	"net/http"
)

func main() {

	var respuesta string = "1234"

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "form.html")
		case "POST":
			r.ParseForm()
			data := r.Form.Get("InputAcceso")

			if data == respuesta {
				http.ServeFile(w, r, "test.html")
			} else {
				fmt.Fprintf(w, "Acceso No Autorizado")
			}
		default:
			fmt.Fprintln(w, "Algo ha pasado")
			fmt.Fprintf(w, "Cierre y vuelva a intentarlo más tarde")
		}

	})

	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
