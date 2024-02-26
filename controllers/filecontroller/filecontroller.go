package filecontroller

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
	"github.com/jeypc/go-auth/libraries"
	"github.com/jeypc/go-auth/models"
)

var fileModel = models.NewFileModel()
var validation = libraries.NewValidation()

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			file, _ := fileModel.FindAll()

			data := map[string]interface{}{
				"files":        file,
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, err := template.ParseFiles("views/file/index.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, data)
		}
	}
}

func Store(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			if r.Method == http.MethodGet {
				temp, err := template.ParseFiles("views/file/index.html")
				if err != nil {
					panic(err)
				}
				temp.Execute(w, nil)
			} else if r.Method == http.MethodPost {
				if err := r.ParseMultipartForm(1024); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				alias := r.FormValue("title")

				uploadedFile, handler, err := r.FormFile("file")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				defer uploadedFile.Close()

				dir, err := os.Getwd()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				filename := handler.Filename
				if alias != "" {
					filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
				}
				fileLocation := filepath.Join(dir, "files", filename)
				targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				defer targetFile.Close()

				if _, err := io.Copy(targetFile, uploadedFile); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// w.Write([]byte("done"))

				r.ParseForm()

				var file entities.File
				file.Title = alias
				file.File = filename
				file.Author = r.Form.Get("author")

				var data = make(map[string]interface{})

				vErrors := validation.Struct(file)

				if vErrors != nil {
					data["file"] = file
					data["validation"] = vErrors
				} else {
					data["pesan"] = "Data file berhasil disimpan"
					fileModel.Create(file)
				}

				temp, _ := template.ParseFiles("views/file/index.html")
				temp.Execute(w, data)

			}
		}
	}
}
