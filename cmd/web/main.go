package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorlog *log.Logger
	infolog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorlog: errorLog,
		infolog:  infoLog,
	}

	mux := http.NewServeMux()
	http.FileServer(http.Dir("./ui/static/"))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	svr := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorlog,
		Handler:  mux,
	}
	infoLog.Printf("Start server on %s\n", *addr)
	err := svr.ListenAndServe()
	errorLog.Fatal(err)
}
