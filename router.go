package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"refresh/controllers"
	"refresh/models"
)

func StartRouter() {
	router := mux.NewRouter()
	routerApi := mux.NewRouter()

	router.HandleFunc("/", controllers.Home)

	//API
	routerApi.HandleFunc("/api/users/get-all", models.GetAllUsers)
	routerApi.HandleFunc("/api/tasks/tasks", models.Task{}.Tasks)
	routerApi.HandleFunc("/api/booking/all", models.Booking{}.GetAllClients)
	routerApi.HandleFunc("/api/saveMapPosition", models.Map{}.SaveMapPosition)
	routerApi.HandleFunc("/api/user/getpermissions", models.UserAccess{}.GetPermisions)
	routerApi.HandleFunc("/api/get-map-position", models.GetMapPosition)
	routerApi.HandleFunc("/api/getSmena", models.Finance{}.GetSmena)
	routerApi.HandleFunc("/api/saveNewPosition", models.Map{}.SaveNewPosition)
	routerApi.HandleFunc("/api/client/getlogin", models.ClientTable{}.GetLogin)
	routerApi.HandleFunc("/api/tasks/users", models.User{}.Users)
	routerApi.HandleFunc("/api/zone/add", models.Zone{}.AddZone)

	http.Handle("/", router)
	http.Handle("/api/", routerApi)

	// Настройка обработчика для статических файлов
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})
	http.HandleFunc("/main.bundle.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/js/main.bundle.js")
	})
	// Задаем путь к папке с ресурсами
	img := http.FileServer(http.Dir("views/img"))

	// Создаем обработчик маршрута для загрузки ресурсов
	http.Handle("/img/", http.StripPrefix("/img/", img))

	css := http.FileServer(http.Dir("views/css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	fonts := http.FileServer(http.Dir("views/fonts"))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", fonts))
}
