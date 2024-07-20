package handlers

import (
    "net/http"
    "strconv"
	"encoding/json"
    "playground/datab"
    "playground/utils"
    "playground/models"
    "time"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
    // Get user ID from session (you'll need to implement this part)
    session, _ := utils.Store.Get(r, "session")
    userID, ok := session.Values["user_id"].(int)
    if !ok {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    // Assuming you're sending the post content in the request body as JSON
    var post models.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Set other fields for the post
    post.UserID = userID
    post.CreatedAt = time.Now()
    // Save the post to the database
    if err := datab.CreatePost(&post); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // Respond with the newly created post data
    utils.JsonResponse(w, http.StatusCreated, post)
}

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := utils.Store.Get(r, "session")
    userID, ok := session.Values["user_id"].(int)
    if !ok {
        http.Error(w, "Forbidden", http.StatusForbidden)
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
    session, _ := utils.Store.Get(r, "session")
    userID, ok := session.Values["user_id"].(int)
    if !ok {
        http.Error(w, "Forbidden", http.StatusForbidden)
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