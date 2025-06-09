package main

import (
	"encoding/json"
	"go-server/internal/config"
	"go-server/internal/logger"
	"net/http"
)

func main() {
	// load logger
	logger.InitLogger()
	defer logger.SyncLogger()

	// load env
	config.LoadConfig()
	// create router
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "It works",
		}
		json.NewEncoder(w).Encode(response)
	})
	logger.Log.Info("Server is started at https://localhost:" + config.AppConfig.Port)
	http.ListenAndServe(":"+config.AppConfig.Port, mux)
}
