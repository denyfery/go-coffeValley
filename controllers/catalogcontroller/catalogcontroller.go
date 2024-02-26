package catalogcontroller

import (
	"html/template"
	"net/http"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/models"
)

var catalogModel = models.NewCatalogModel()

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			catalog, _ := catalogModel.FindAll()

			data := map[string]interface{}{
				"catalogs":     catalog,
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, err := template.ParseFiles("views/catalog/index.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, data)
		}
	}
}
