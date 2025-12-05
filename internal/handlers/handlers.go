package handlers

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read the file", http.StatusInternalServerError)
		return
	}
	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newName := time.Now().UTC().String() + filepath.Ext(handler.Filename)
	out, err := os.Create(newName)
	if err != nil {
		http.Error(w, "failed to create new file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	out.Write([]byte(result))
	w.Write([]byte(result))

}
