package web

import (
	"fmt"
	"github.com/duruyao/gochat/server/conf"
	mlog "github.com/duruyao/gochat/server/log"
	"golang.org/x/net/http2"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var files []string
	if rand.Int()%2 == 1 { // TODO: check if the current user is logged in by Cookie
		files = []string{
			HtmlPath("layout"),
			HtmlPath("login.header"),
			HtmlPath("index.content"),
		}
	} else {
		files = []string{
			HtmlPath("layout"),
			HtmlPath("logout.header"),
			HtmlPath("index.content"),
		}
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.ExecuteTemplate(w, "layout.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if "GET" == r.Method {
		files := []string{
			HtmlPath("layout"),
			HtmlPath("logout.header"),
			HtmlPath("login.content"),
		}
		tmpl := template.Must(template.ParseFiles(files...))
		if err := tmpl.ExecuteTemplate(w, "layout.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if "POST" == r.Method {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(fmt.Sprint(r.Form))); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if "GET" == r.Method {
		files := []string{
			HtmlPath("layout"),
			HtmlPath("logout.header"),
			HtmlPath("sign_up.content"),
		}
		tmpl := template.Must(template.ParseFiles(files...))
		if err := tmpl.ExecuteTemplate(w, "layout.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if "POST" == r.Method {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(fmt.Sprint(r.Form))); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GoRunWebApp() {
	mux := http.NewServeMux()
	resHandler := http.FileServer(http.Dir(ResourceDir()))
	mux.Handle("/static/", http.StripPrefix("/static/", resHandler))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/sign_up", signUpHandler)

	app := http.Server{
		Addr:              conf.Addr(),
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       7 * time.Second,
		ReadHeaderTimeout: 7 * time.Second,
		WriteTimeout:      7 * time.Second,
		IdleTimeout:       7 * time.Second,
		MaxHeaderBytes:    1024 * 1024,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          mlog.ErrorLogger,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	if conf.HttpsEnable() {
		if err := http2.ConfigureServer(&app, &http2.Server{}); err != nil {
			mlog.FatalLn(err)
		}
		mlog.FatalLn(app.ListenAndServeTLS(TLSCertPath(), TLSKeyPath()))
	} else {
		mlog.FatalLn(app.ListenAndServe())
	}
}
