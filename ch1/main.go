package main

import (
	"fmt"
	//"log"
	"net/http"
	"log"
	"sync"
	"text/template"
	"path/filepath"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t * templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates",
			t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {

	fmt.Println("start!")

	http.Handle("/", &templateHandler{filename: "chat.html"})
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(`
	//		<html>
	//			<head>
	//				<title>Chat</title>
	//			</head>
	//			<body>
	//				Chat!
	//			</body>
	//		</html>
	//	`))
	//})
	// webサーバを開始します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

