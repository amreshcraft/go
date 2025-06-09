package response
import (
    "encoding/json"
    "net/http"
)

type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}
func SendErrorResponse(w http.ResponseWriter, code int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(ErrorResponse{
        Code:    code,
        Message: message,
    })
}