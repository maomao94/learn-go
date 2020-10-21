package main

import (
	"learn-go/errorhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter,
	r *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		defer func() {
			r := recover()
			log.Printf("Panic: %v", r)
			http.Error(writer,
				http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error Handling request: %s",
				err.Error())
			code := http.StatusOK
			switch {
			case os.IsExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code),
				code)
		}
	}
}

func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
