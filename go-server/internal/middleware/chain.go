package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(mainHandler http.Handler, middleware ...Middleware) http.Handler{

for i := 0; i <= len(middleware) - 1; i++ {
	mainHandler = middleware[i](mainHandler)
}

return mainHandler

}