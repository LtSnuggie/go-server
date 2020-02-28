package server

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func ReturnSuccess(w http.ResponseWriter) {
	ReturnSuccessMessage(w, []byte("success"))
}

func ReturnSuccessMessage(w http.ResponseWriter, message []byte) {
	w.Write(message)
	// w.WriteHeader(200)
}

func ReturnError(w http.ResponseWriter, err error, msg string) {
	errID := uuid.Must(uuid.NewRandom()).String()
	log.SetReportCaller(true)
	log.WithFields(log.Fields{
		"ts":    time.Now(),
		"err":   err,
		"msg":   msg,
		"errID": errID,
	}).Error()
	log.SetReportCaller(false)
	w.WriteHeader(500)
	w.Write([]byte(errID))
}
