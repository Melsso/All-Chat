package handlers

import (
    "net/http"
    "strconv"
	"fmt"
    "playground/datab"
    "playground/utils"
    "playground/models"
    /*"time"*/
)

func MessageHandler(w http.ResponseWriter, r * http.Request) {
	session, _ := utils.Store.Get(r, "session")
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
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