package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetAsset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath := r.PathValue("path") // Get the file path from the URL

		// Construct the absolute path to the file
		absolutePath := filePath
		log.Println(absolutePath)

		// Open the file
		file, err := os.Open(absolutePath)
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Файл не найден"}, http.StatusBadRequest).Respond(w)
			log.Println(err)
			return
		}
		defer file.Close()

		// Get file info
		fileInfo, err := file.Stat()
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Файл не найден"}, http.StatusBadRequest).Respond(w)
			return
		}

		// Set the content type based on the file extension
		contentType := getContentType(filePath)
		w.Header().Set("Content-Type", contentType)

		// Set the content length
		w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

		// Copy the file to the response
		_, err = io.Copy(w, file)
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Нет возможности отправить файл"}, http.StatusBadRequest).Respond(w)
		}
	}
}

func getContentType(filePath string) string {
	switch filepath.Ext(filePath) {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	case ".txt":
		return "text/plain"
	case ".csv":
		return "text/csv"
	case ".pdf":
		return "application/pdf"
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".xls":
		return "application/vnd.ms-excel"
	case ".xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case ".ppt":
		return "application/vnd.ms-powerpoint"
	case ".pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case ".mp3":
		return "audio/mpeg"
	case ".wav":
		return "audio/wav"
	case ".ogg":
		return "audio/ogg"
	case ".mp4":
		return "video/mp4"
	case ".avi":
		return "video/x-msvideo"
	case ".wmv":
		return "video/x-ms-wmv"
	case ".png":
		return "image/png"
	case ".jpeg", ".jpg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".bmp":
		return "image/bmp"
	case ".svg":
		return "image/svg+xml"
	case ".eot":
		return "application/vnd.ms-fontobject"
	case ".otf":
		return "font/otf"
	case ".ttf":
		return "font/ttf"
	case ".woff":
		return "font/woff"
	case ".woff2":
		return "font/woff2"
	default:
		return "application/octet-stream"
	}
}
