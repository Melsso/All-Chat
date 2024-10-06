package req_handlers

import (
	"All-Chat/back-end/datab"
	"fmt"
	"All-Chat/back-end/models"
	"net/http"
	"strconv"
	"All-Chat/back-end/utils"
	/*"time"*/)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Forbidden Message: Not authenticated", http.StatusForbidden)
		return
	}
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "Forbidden Message: No user_id", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodGet {
		friendIDs := r.URL.Query().Get("friend_id")
		if friendIDs == "" {
			http.Error(w, "Missing friend_id", http.StatusBadRequest)
			return
		}
		friend_id, err := strconv.Atoi(friendIDs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		convID, err := datab.StartOrGetConversation(userID, friend_id)
		if err != nil {
			fmt.Println("datab error startorget: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages, err := datab.GetMessages(convID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if messages == nil {
			messages = []models.Message{}
		}
		response := map[string]interface{}{
			"conversation_id": convID,
			"messages":        messages,
		}
		utils.JsonResponse(w, http.StatusOK, response)

	} else if r.Method == http.MethodPost {
		// handle new messages sent, does nothing for now
	} else {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}
