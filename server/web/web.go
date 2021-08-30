package web

import (
	"github.com/duruyao/gochat/server/key"
	mlog "github.com/duruyao/gochat/server/log"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type PageData struct {
	HttpAddr string
	Rid      string
}

func makeHtmlHandler(tmpl *template.Template, filename string) http.HandlerFunc {
	if "join_room.html" == filename {
		return func(w http.ResponseWriter, r *http.Request) {
			rid := r.FormValue("rid")
			if len(rid) < 1 {
				http.Error(w, "no rid specified", http.StatusBadRequest)
				return
			}
			data := PageData{Rid: rid}
			if err := tmpl.ExecuteTemplate(w, filename, data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		data := PageData{}
		if err := tmpl.ExecuteTemplate(w, filename, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func makeSourceHandler(filepath string) http.HandlerFunc {
	content, _ := ioutil.ReadFile(filepath)
	return func(w http.ResponseWriter, r *http.Request) {
		if len(content) < 1 {
			http.Error(w, "empty content", http.StatusInternalServerError)
		}
		if _, err := w.Write(content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GoHttpService(addr string) {
	for _, path := range AllSourcePaths() {
		http.Handle("/"+filepath.Base(path), makeSourceHandler(path))
	}
	tmpl := template.Must(template.ParseFiles(AllHtmlPaths()...))
	http.Handle("/", makeHtmlHandler(tmpl, "sign_in.html"))
	for _, path := range AllHtmlPaths() {
		http.Handle("/"+filenameWithoutExt(path), makeHtmlHandler(tmpl, filepath.Base(path)))
	}
	mlog.FatalLn(http.ListenAndServe(addr, nil))
}

func GoHttpsService(addr string) {
	for _, path := range AllSourcePaths() {
		http.Handle("/"+filepath.Base(path), makeSourceHandler(path))
	}
	tmpl := template.Must(template.ParseFiles(AllHtmlPaths()...))
	http.Handle("/", makeHtmlHandler(tmpl, "sign_in.html"))
	for _, path := range AllHtmlPaths() {
		http.Handle("/"+filenameWithoutExt(path), makeHtmlHandler(tmpl, filepath.Base(path)))
	}
	mlog.FatalLn(http.ListenAndServeTLS(addr, key.Path("crt"), key.Path("key"), nil))
}
