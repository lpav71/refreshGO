package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"refresh/controllers"
	"refresh/models"
)

func StartRouter() {
	router := mux.NewRouter()
	routerApi := mux.NewRouter()

	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/shop", controllers.Shop)
	router.HandleFunc("/map", controllers.Maps)
	router.HandleFunc("/warehouse", controllers.Warehouse)
	router.HandleFunc("/finduser", controllers.Users)

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
	routerApi.HandleFunc("/api/tasks/save/edit", models.Task{}.SaveEditModal)
	routerApi.HandleFunc("/api/shop/get", models.Store{}.GetAll)
	routerApi.HandleFunc("/api/client/all", models.ClientTable{}.GetClients)
	routerApi.HandleFunc("/api/shop/find/client", models.ClientTable{}.GetByLogin)
	routerApi.HandleFunc("/api/computer/all", models.Map{}.GetAllComputers)
	routerApi.HandleFunc("/api/store/operation/type", models.StoreOperationType{}.GetStoreOparationType)
	routerApi.HandleFunc("/api/client/find", models.ClientTable{}.UsersFind)
	routerApi.HandleFunc("/api/age-filter", models.ClientTable{}.AgeFilter)

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

	// Разрешаем загружать любые файлы из папки images
	http.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		// Получаем запрошенный путь после "/images/"
		imagePath := r.URL.Path[len("/images/"):]

		// Получаем абсолютный путь к изображению
		absPath, err := filepath.Abs(filepath.Join("views/images", imagePath))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Проверяем, существует ли файл
		info, err := os.Stat(absPath)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Проверяем, является ли файл обычным файлом
		if info.Mode().IsRegular() {
			// Открываем файл
			file, err := os.Open(absPath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// Передаем файл в ответ
			http.ServeContent(w, r, info.Name(), info.ModTime(), file)
			return
		}

		http.NotFound(w, r)
	})
}
