package main

import (
	"log"
	"net/http"
	"test-memcached/controllers"
	"test-memcached/models"
)

func main() {
	//fmt.Println("Hello, World!")
	addr := ":8086"

	//подключаемся к БД
	models.ConnectDatabase()

	//запускаем миграцию
	models.DBMigrate()

	//создаем роутер
	mux := http.NewServeMux()

	// В хендлере будем использовать параметры пути ({id}). Введены в go 1.22.
	// Можно посмотреть видео: https://www.youtube.com/watch?v=H7tbjKFSg58&t=48s
	mux.HandleFunc("/blogs/{id}", controllers.BlogsShow)

	//запускаем сервер
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))

}
