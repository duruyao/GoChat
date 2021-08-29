package web

import (
	"github.com/duruyao/gochat/server/key"
	mlog "github.com/duruyao/gochat/server/log"
	"html/template"
	"net/http"
	"sync"
)

func joinRoomHandle(w http.ResponseWriter, r *http.Request, tmpl *template.Template, filename string) {
	data := struct {
		Rid string
	}{r.FormValue("rid")}
	if err := tmpl.ExecuteTemplate(w, filename, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func simpleHandle(w http.ResponseWriter, r *http.Request, tmpl *template.Template, filename string) {
	if err := tmpl.ExecuteTemplate(w, filename, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(handleFunc func(http.ResponseWriter, *http.Request, *template.Template, string), tmpl *template.Template, filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleFunc(w, r, tmpl, filename)
	}
}

var httpServiceOnce sync.Once

func GoHttpService(addr string) {
	httpServiceOnce.Do(func() {
		tmpl := template.Must(template.ParseFiles(AllHtmlPaths()...))
		http.HandleFunc("/sign_in", makeHandler(simpleHandle, tmpl, "sign_in.html"))
		http.HandleFunc("/change_pwd", makeHandler(simpleHandle, tmpl, "change_pwd.html"))
		http.HandleFunc("/add_admin", makeHandler(simpleHandle, tmpl, "add_admin.html"))
		http.HandleFunc("/add_room", makeHandler(simpleHandle, tmpl, "add_room.html"))
		http.HandleFunc("/join_room", makeHandler(joinRoomHandle, tmpl, "join_room.html"))
		mlog.FatalLn(http.ListenAndServe(addr, nil))
	})
}

var httpServicesOnce sync.Once

func GoHttpsService(addr string) {
	httpServicesOnce.Do(func() {
		tmpl := template.Must(template.ParseFiles(AllHtmlPaths()...))
		http.HandleFunc("/sign_in", makeHandler(simpleHandle, tmpl, "sign_in.html"))
		http.HandleFunc("/change_pwd", makeHandler(simpleHandle, tmpl, "change_pwd.html"))
		http.HandleFunc("/add_admin", makeHandler(simpleHandle, tmpl, "add_admin.html"))
		http.HandleFunc("/add_room", makeHandler(simpleHandle, tmpl, "add_room.html"))
		http.HandleFunc("/join_room", makeHandler(joinRoomHandle, tmpl, "join_room.html"))
		mlog.FatalLn(http.ListenAndServeTLS(addr, key.Path("crt"), key.Path("key"), nil))
	})
}
