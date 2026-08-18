package pkg

import (
	"bytes"
	"fmt"
	"net/http"
)

func CheckMethodGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		NotFoundResponse(w)
		return false
	}

	return true
}

func CheckMethodPost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		NotFoundResponse(w)
		return false
	}

	return true
}

func NotFoundResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func GetErrorResponse(w http.ResponseWriter, handlerName string, err error, statusCode int) {
	w.WriteHeader(statusCode)

	buf := bytes.NewBufferString(handlerName)
	buf.WriteString(": ")
	buf.WriteString(err.Error())
	fmt.Println(buf.String())
	buf.WriteString("\n")

	_, _ = w.Write(buf.Bytes())
}

func GetSuccessResponseWithBody(w http.ResponseWriter, body []byte, statusCode int) {
	w.WriteHeader(statusCode)

	_, _ = w.Write(body)
}

func GetSuccessResponse(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}
