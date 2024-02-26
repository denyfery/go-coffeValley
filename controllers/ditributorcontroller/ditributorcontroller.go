package ditributorcontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
	"github.com/jeypc/go-auth/libraries"
	"github.com/jeypc/go-auth/models"
)

var validation = libraries.NewValidation()
var distributorModel = models.NewDistributorModel()

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			distributor, _ := distributorModel.FindAll()

			data := map[string]interface{}{
				"distributors": distributor,
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, err := template.ParseFiles("views/distributor/index.html")
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
				temp, err := template.ParseFiles("views/distributor/add.html")
				if err != nil {
					panic(err)
				}
				temp.Execute(w, nil)
			} else if r.Method == http.MethodPost {

				r.ParseForm()

				var distributor entities.Distributor
				distributor.Name = r.Form.Get("name")
				distributor.City = r.Form.Get("city")
				distributor.Region = r.Form.Get("region")
				distributor.Country = r.Form.Get("country")
				distributor.Phone = r.Form.Get("phone")
				distributor.Email = r.Form.Get("email")

				var data = make(map[string]interface{})

				vErrors := validation.Struct(distributor)

				if vErrors != nil {
					data["distributor"] = distributor
					data["validation"] = vErrors
				} else {
					data["pesan"] = "Data distributor berhasil disimpan"
					distributorModel.Create(distributor)
				}

				temp, _ := template.ParseFiles("views/distributor/add.html")
				temp.Execute(w, data)
			}
		}
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			if r.Method == http.MethodGet {
				queryString := r.URL.Query()
				id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

				var distributor entities.Distributor
				distributorModel.Find(id, &distributor)

				data := map[string]interface{}{
					"distributor": distributor,
				}

				temp, err := template.ParseFiles("views/distributor/update.html")
				if err != nil {
					panic(err)
				}
				temp.Execute(w, data)
			} else if r.Method == http.MethodPost {
				r.ParseForm()

				var distributor entities.Distributor
				distributor.Id, _ = strconv.ParseInt(r.Form.Get("id"), 10, 64)
				distributor.Name = r.Form.Get("name")
				distributor.City = r.Form.Get("city")
				distributor.Region = r.Form.Get("region")
				distributor.Country = r.Form.Get("country")
				distributor.Phone = r.Form.Get("phone")
				distributor.Email = r.Form.Get("email")

				var data = make(map[string]interface{})

				vErrors := validation.Struct(distributor)

				if vErrors != nil {
					data["distributor"] = distributor
					data["validation"] = vErrors
				} else {
					data["pesan"] = "Data distributor berhasil diperbarui"
					distributorModel.Update(distributor)
				}

				temp, _ := template.ParseFiles("views/distributor/update.html")
				temp.Execute(w, data)
			}
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	distributorModel.Delete(id)

	http.Redirect(w, r, "/distributor", http.StatusSeeOther)
}
