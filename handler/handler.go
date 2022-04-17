package handler

import (
	"golang-web/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome Home"))

	tempt, err :=  template.ParseFiles(path.Join("views", "index.html"),  path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "error bro", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title": "Home",
	// 	"name": "Jery Hardianto",
	// 	"age": 20,
	// 	"address": "Janti",
	// }
	
	// data := entity.Product{ID: 1, Name: "Product 1", Price: 10000, Stock: 10}
	data := []entity.Product{
		{ID: 1, Name: "Product 1", Price: 10000, Stock: 3},
		{ID: 2, Name: "Product 2", Price: 20000, Stock: 8},
		{ID: 3, Name: "Product 3", Price: 30000, Stock: 1},
		
	}

	err = tempt.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "error bro", http.StatusInternalServerError)
		return
	}
}

func Sayhello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World and Golang"))
}

func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}

func Product(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNum, err := strconv.Atoi(id)
	if err != nil || idNum < 1 {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}


	tempt, err :=  template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "error bro", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNum,
	}

	err = tempt.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "error bro", http.StatusInternalServerError)
		return
	}
}

//* POST & GET
func PostGet(w http.ResponseWriter, r *http.Request) {
	//Cek method yang digunakan
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		w.Write([]byte("Method tidak dikenali"))
	}


}

func PostForm(w http.ResponseWriter, r *http.Request) {
	//Cek method yang digunakan
	if r.Method == "GET" {
		tempt, err :=  template.ParseFiles(path.Join("views", "post.html"), path.Join("views", "layout.html"))
		if err != nil { 
			log.Println(err)
			http.Error(w, "error bro", http.StatusInternalServerError)
			return
		}

	
		err = tempt.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "error bro", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Method tidak dikenali", http.StatusMethodNotAllowed)
}

func Process(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			log.Println(err)
			http.Error(w, "error bro", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		pesan := r.Form.Get("pesan")


		data := map[string]interface{}{
			"name": name,
			"pesan": pesan,
		}
		templ, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "error bro", http.StatusInternalServerError)
			return
		}

		err = templ.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "error bro", http.StatusInternalServerError)
			return
		}

		return


	}
}


























