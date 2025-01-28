package transport

import (
	"fmt"
	//"log"
	"net/http"
	"strconv"
)

func (s *APIHandlers) MyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (s *APIHandlers) GetLastHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем значение строки запроса count
	countParam := r.URL.Query().Get("count") // api/transactions?count=7
	//log.Fatal(countParam)

	// Пытаемся преобразовать в число (опционально)
	count, err := strconv.Atoi(countParam)
	if err != nil || count <= 0 {
		http.Error(w, "Invalid count parameter", http.StatusBadRequest)
		return
	}

	// Обработка логики (в данном случае просто ответ)
	fmt.Println(countParam)
	w.Write([]byte(fmt.Sprintf("Fetching %d transactions\n", count)))

}
