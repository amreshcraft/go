package main

import (
	"encoding/json"
	"go-server/internal/config"
	"go-server/internal/cube"
	"go-server/internal/logger"
	"go-server/internal/middleware"
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
	mux.HandleFunc("/cube", cube.CalculateCubeHandler)
	middlewareChain := middleware.Chain(mux,middleware.LoggerMiddleware,middleware.RecoveryMiddleware)
	logger.Log.Info("Server is started at http://localhost:" + config.AppConfig.Port)
	http.ListenAndServe(":"+config.AppConfig.Port, middlewareChain)
}
