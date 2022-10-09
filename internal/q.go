package q

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

type ResultDetails struct {
	Success  bool
	Username string
	Song     string
	Response string
}

//go:embed assets
var content embed.FS

func Web(port string) {
	filesys := fs.FS(content)
	tmpl := template.Must(template.ParseFS(filesys, "assets/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		username := r.FormValue("username")
		fmt.Println("Request Username: " + username)
		song := r.FormValue("song")
		fmt.Println("Request Song: " + song)
		response := "test"
		fmt.Println("Response: " + response)

		result := ResultDetails{
			Success:  true,
			Username: username,
			Song:     song,
			Response: response,
		}
		tmpl.Execute(w, result)
	})

	http.ListenAndServe(":"+port, nil)
}
