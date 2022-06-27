package main

import (
	stdlog "log"
	"net/http"
	"os"

	admissioncontrol "github.com/elithrar/admission-control"
	log "github.com/go-kit/kit/log"
	"github.com/joaquinalvarezdev/entropy/controls"
)

func main() {
	router := http.NewServeMux()

	//Routes
	router.HandleFunc("/API/entropy", controls.AnalyzeFile)

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)

	// Create an instance of our LoggingMiddleware with our configured logger
	loggingMiddleware := admissioncontrol.LoggingMiddleware(logger)
	loggedRouter := loggingMiddleware(router)

	http.ListenAndServe(":8000", loggedRouter)
}
