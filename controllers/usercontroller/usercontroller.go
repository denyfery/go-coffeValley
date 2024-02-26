package usercontroller

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
	"github.com/jeypc/go-auth/libraries"
	"github.com/jeypc/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

var userModel = models.NewUserModel()
var validation = libraries.NewValidation()

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"data":         template.HTML(GetData()),
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, _ := template.ParseFiles("views/user.html")
			temp.Execute(w, data)
		}
	}
}

func GetData() string {
	buffer := &bytes.Buffer{}
	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/user/data.html")

	users, _ := userModel.FindAll()

	data := map[string]interface{}{
		"users": users,
	}
	temp.ExecuteTemplate(buffer, "data.html", data)
	return buffer.String()
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/user/tambah.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		r.ParseForm()

		user := entities.User{
			NamaLengkap: r.Form.Get("nama_lengkap"),
			Email:       r.Form.Get("email"),
			Username:    r.Form.Get("username"),
			Password:    r.Form.Get("password"),
			Cpassword:   r.Form.Get("cpassword"),
		}

		errorMessages := validation.Struct(user)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
				"user":       user,
			}

			temp, _ := template.ParseFiles("views/user/tambah.html")
			temp.Execute(w, data)
		} else {

			// hashPassword
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			// insert ke database
			userModel.Create(user)

			data := map[string]interface{}{
				"pesan": "Tambah data berhasil",
				"data":  template.HTML(GetData()),
			}
			temp, _ := template.ParseFiles("views/user.html")
			temp.Execute(w, data)
		}
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		queryString := r.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var user entities.User
		userModel.Find(id)

		data := map[string]interface{}{
			"user": user,
		}

		temp, err := template.ParseFiles("views/user/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		var user entities.User
		user.Id, _ = strconv.ParseInt(r.Form.Get("id"), 10, 64)
		user.NamaLengkap = r.Form.Get("nama_lengkap")
		user.Email = r.Form.Get("email")
		user.Username = r.Form.Get("username")
		user.Password = r.Form.Get("password")
		user.Cpassword = r.Form.Get("cpassword")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(user)

		if vErrors != nil {
			data["user"] = user
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data user berhasil diperbarui"
			userModel.Update(user)
		}

		temp, _ := template.ParseFiles("views/user/edit.html")
		temp.Execute(w, data)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	userModel.Delete(id)

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}
