package erros

import (
	"fmt"
	"log"
	"net/http"
)

func HttpError(w http.ResponseWriter, modulo string, err error) {
  http.Error(w, fmt.Sprintf("%s: %v", modulo, err.Error()), http.StatusInternalServerError)
}

func LogError(modulo string, err error) {
  log.Println(modulo, ": " ,err.Error())
}
