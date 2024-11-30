package controllers

import (
	"fmt"
	"mime" // 追加
	"net/http"
	"text/template"
	"todo_app/app/models"
	"todo_app/config"
)

func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// sessionId関数を削除し、session関数のみを使用
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUid: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

func StartMainServer() error {
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".js", "application/javascript")

	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todo", index)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
