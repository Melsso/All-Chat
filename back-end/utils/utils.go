package utils

import (
	"encoding/json"
	"net/http"
    "log"
    "os"
    "github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func JsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "https://localhost")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

func Init() {
    sessionKey := os.Getenv("SESSION_KEY")
    if sessionKey == "" {
        log.Fatal("SESSION_KEY environment variable is not set")
    }

    Store = sessions.NewCookieStore([]byte(sessionKey))
    Store.Options = &sessions.Options{
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteNoneMode,
    }
}