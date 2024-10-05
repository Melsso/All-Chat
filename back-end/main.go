// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))


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

	r.HandleFunc("/login", req_handlers.LoginHandler)
	r.HandleFunc("/register", req_handlers.RegisterHandler)
	r.HandleFunc("/home", req_handlers.HomeHandler)
	r.HandleFunc("/create-post", req_handlers.CreatePostHandler)
	r.HandleFunc("/like-post", req_handlers.LikePostHandler)
	r.HandleFunc("/add-comment", req_handlers.AddCommentHandler)
	r.HandleFunc("/logout", req_handlers.LogoutHandler)
	r.HandleFunc("/add-friend", req_handlers.AddFriendHandler)
	r.HandleFunc("/accept-friend", req_handlers.AcceptFriendHandler)
	r.HandleFunc("/lookup-user", req_handlers.LookUpFriendHandler)
	r.HandleFunc("/delete-friend", req_handlers.DeleteFriendHandler)
	r.HandleFunc("/invite-list", req_handlers.InviteListHandler)
	r.HandleFunc("/messages", req_handlers.MessageHandler)
	
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:80", "http://localhost"}),
		handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"}),
		handlers.AllowCredentials(),
	)

	handler := corsHandler(r)
	// Start the server
	fmt.Println("Started server on port 8000")
	http.ListenAndServe(":8000", handler)

}
