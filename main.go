package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/jeypc/go-auth/controllers"
	"github.com/jeypc/go-auth/controllers/catalogcontroller"
	"github.com/jeypc/go-auth/controllers/ditributorcontroller"
	"github.com/jeypc/go-auth/controllers/filecontroller"
	"github.com/jeypc/go-auth/controllers/usercontroller"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)

	// user
	http.HandleFunc("/user", usercontroller.Index)
	http.HandleFunc("/user/store", usercontroller.Store)
	http.HandleFunc("/user/delete", usercontroller.Delete)

	//catalog
	http.HandleFunc("/catalog", catalogcontroller.Index)

	//ditributor
	http.HandleFunc("/distributor", ditributorcontroller.Index)
	http.HandleFunc("/distributor/store", ditributorcontroller.Store)
	http.HandleFunc("/distributor/update", ditributorcontroller.Update)
	http.HandleFunc("/distributor/delete", ditributorcontroller.Delete)

	//file
	http.HandleFunc("/upload", filecontroller.Index)
	http.HandleFunc("/file/store", filecontroller.Store)

	fmt.Println("Server jalan di: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
