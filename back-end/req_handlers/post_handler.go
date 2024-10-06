package req_handlers

import (
	"All-Chat/back-end/datab"
	"All-Chat/back-end/models"
	"All-Chat/back-end/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusNoContent)
		return 
	}

	session, _ := utils.Store.Get(r, "auth")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden Create Post: Not authenticated", http.StatusForbidden)
		return
	}
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "Forbidden Create Post: No user_id", http.StatusForbidden)
		return
	}
	
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.UserID = userID
	post.CreatedAt = time.Now()
	if err := datab.CreatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JsonResponse(w, http.StatusCreated, post)
}

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusNoContent)
		return 
	}

	session, _ := utils.Store.Get(r, "auth")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden Like Post: Not authenticated", http.StatusForbidden)
		return
	}
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "Forbidden Like Post: No user_id", http.StatusForbidden)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var requestData struct {
		PostID string `json:"post_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(requestData.PostID)
	if err != nil {
		http.Error(w, "Invalid post_id format", http.StatusBadRequest)
		return
	}
	err = datab.LikePost(userID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JsonResponse(w, http.StatusOK, map[string]string{"status": "liked"})
}

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusNoContent)
		return 
	}
	
	session, _ := utils.Store.Get(r, "auth")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden Comment: Not authenticated", http.StatusForbidden)
		return
	}
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "Forbidden Comment: No user_id", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodGet {
		postID := r.URL.Query().Get("post_id")
		if postID == "" {
			http.Error(w, "Missing post_id parameter", http.StatusBadRequest)
			return
		}
		pid, err := strconv.Atoi(postID)
		if err != nil {
			http.Error(w, "Invalid post_id format", http.StatusBadRequest)
			return
		}
		comments, err := datab.GetComments(pid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"comments": comments,
		}
		utils.JsonResponse(w, http.StatusOK, data)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var requestData struct {
		PostID  string `json:"post_id"`
		Content string `json:"content"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postID, err := strconv.Atoi(requestData.PostID)
	if err != nil {
		http.Error(w, "Invalid post_id format", http.StatusBadRequest)
		return
	}
	err = datab.CommentPost(userID, postID, requestData.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JsonResponse(w, http.StatusOK, map[string]string{"status": "commented"})
}
