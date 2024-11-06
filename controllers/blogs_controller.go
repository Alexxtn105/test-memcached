package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"test-memcached/models"
	"time"
)

func BlogsShow(w http.ResponseWriter, r *http.Request) {

	// Извлекаем ИД из параметров пути
	//idStr := r.URL.Path[len("/blogs/"):] // Старый вариант  (до версии 1.22)
	idStr := r.PathValue("id") //сразу берем параметр из пути (стало доступно в go версии 1.22)
	// конвертим строку в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Ivalid blog ID", http.StatusBadRequest)
		return
	}
	log.Println(idStr)
	// сперва посмотрим, есть ли статья в кэше
	// если в кэше нет, кэшируем данные
	data := models.CacheData("blog:"+idStr, 60, func() []byte {
		// вытаскиваем блог из БД
		blog := models.BlogsFind(uint64(id))

		// преобразуем в слайс байтов
		blogBytes, _ := json.Marshal(blog)

		// имитируем длительный процесс
		time.Sleep(time.Second * 2)
		return blogBytes
	})

	// посылаем ответ пользователю
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}
