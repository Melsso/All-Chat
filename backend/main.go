package main

import (
	"fmt"
	"net/http"

	"playground/datab"
	"playground/handlers"
	"playground/utils"
)

func main() {
	datab.InitDB()
	defer datab.CloseDB()
	utils.Init()
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register handlers for specific routes
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/create-post", handlers.CreatePostHandler)
	http.HandleFunc("/like-post", handlers.LikePostHandler)
	http.HandleFunc("/add-comment", handlers.AddCommentHandler) 
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/add-friend", handlers.AddFriendHandler)
	http.HandleFunc("/accept-friend", handlers.AcceptFriendHandler)
	http.HandleFunc("/lookup-user", handlers.LookUpFriendHandler)
	http.HandleFunc("/delete-friend", handlers.DeleteFriendHandler)
	http.HandleFunc("/invite-list", handlers.InviteListHandler)
	http.HandleFunc("/messages", handlers.MessageHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/login.html", http.StatusFound)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fmt.Println("Started server on port 8080")
}
