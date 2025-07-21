package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, "Ошибка парсинга формы", http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	response, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusBadRequest)
		return
	}

	convert, err := service.Convert(response)
	if err != nil {
		http.Error(w, "Ошибка конвертации файла", http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("20250713T150405Z")
	ext := filepath.Ext(header.Filename)
	localFileName := timestamp + ext

	localFile, err := os.Create(localFileName)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	_, err = localFile.WriteString(convert)
	if err != nil {
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(convert))

}
