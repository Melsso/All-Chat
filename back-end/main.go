package main

import (
	"fmt"
	"net/http"

	"All-Chat/back-end/datab"
	"All-Chat/back-end/req_handlers"
	"All-Chat/back-end/utils"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	datab.InitDB()
	defer datab.CloseDB()
	utils.Init()
	
	r := mux.NewRouter()

	r.HandleFunc("/login", req_handlers.LoginHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/register", req_handlers.RegisterHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/home", req_handlers.HomeHandler)
	r.HandleFunc("/create-post", req_handlers.CreatePostHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/like-post", req_handlers.LikePostHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/add-comment", req_handlers.AddCommentHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/logout", req_handlers.LogoutHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/add-friend", req_handlers.AddFriendHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/accept-friend", req_handlers.AcceptFriendHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/lookup-user", req_handlers.LookUpFriendHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/delete-friend", req_handlers.DeleteFriendHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/invite-list", req_handlers.InviteListHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/messages", req_handlers.MessageHandler).Methods("GET", "POST", "OPTIONS")
	
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"https://localhost:443", "https://localhost"}),
		handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"}),
		handlers.AllowCredentials(),
	)

	handler := corsHandler(r)
	// Start the server
	fmt.Println("Started server on port 8443...")
	http.ListenAndServeTLS(":8443", "localhost.pem", "localhost-key.pem", handler)

}
