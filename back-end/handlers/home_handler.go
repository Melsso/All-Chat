package handlers

import (
	"All-Chat/back-end/datab"
	"All-Chat/back-end/models"
	"net/http"
	"All-Chat/back-end/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden not authenticated", http.StatusForbidden)
		return
		
	}
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "Forbidden no user_id", http.StatusForbidden)
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
		"friends": friends,
		"posts":   posts,
		"invite":  invitelist,
	}
	w.Header().Set("Cache-Control", "no-store")
	utils.JsonResponse(w, http.StatusOK, data)
}
