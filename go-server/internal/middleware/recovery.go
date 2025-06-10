package middleware

import (
	"go-server/internal/logger"
	"go-server/internal/response"
	"net/http"
	"go.uber.org/zap"
)


func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil { 
                logger.Log.Error("Panic recovered",
                    zap.Any("error", err),             
                    zap.String("method", r.Method),     
                    zap.String("path", r.URL.Path),     
                )

                response.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
            }
        }()

        next.ServeHTTP(w, r)
    })
}
