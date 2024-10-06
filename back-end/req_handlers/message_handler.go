package req_handlers

import (
	"All-Chat/back-end/datab"
	"fmt"
	"All-Chat/back-end/models"
	"net/http"
	"strconv"
	"encoding/json"
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
			fmt.Println("friend id error: ", err)
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
			fmt.Println("getmsg error: ", err)
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
		var msgConv models.MsgConversation
		err := json.NewDecoder(r.Body).Decode(&msgConv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return 
		}
		
		er := datab.SendMessage(msgConv.ConversationID, userID, msgConv.Content)
		if er != nil {
			fmt.Println("Could not save msg on db: ", er.Error())
			http.Error(w, er.Error(), http.StatusInternalServerError)
			return 	
		}

		response := map[string]interface{}{
			"status": "sent",
		}
		utils.JsonResponse(w, http.StatusOK, response)

	} else {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}
