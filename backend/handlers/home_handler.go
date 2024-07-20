package handlers

import (
	"net/http"
	"playground/datab"
	"playground/utils"
	"playground/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session")
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return 

	}
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/home.html")
		return
	}
	friends, err := datab.GetFriends(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	posts, err := datab.GetPosts(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	invitelist, err := datab.GetInvites(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	if friends == nil {
		friends = []models.User{}
	}
	if posts == nil {
		posts = []models.Post{}
	}
	if invitelist == nil {
		invitelist = []models.User{}
	}
	data := map[string]interface{}{
		"friends":	friends,
		"posts":	posts,
		"invite":	invitelist,
	}
	w.Header().Set("Cache-Control", "no-store")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")
	utils.JsonResponse(w, http.StatusOK, data)
}
