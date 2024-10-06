package req_handlers

import (
	"All-Chat/back-end/datab"
	"encoding/json"
	"fmt"
	"All-Chat/back-end/models"
	"net/http"
	"strconv"
	"All-Chat/back-end/utils"
)

func AddFriendHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Forbidden AddFriend: Not authenticated", http.StatusForbidden)
		return
	}
	
	if r.Method == http.MethodGet {
		// do stuff
		return
	}

	var req map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	
	userID, okUserID := session.Values["user_id"].(int)
	friendIDs, okFriendID := req["friend_id"].(string)
	friendID, err := strconv.Atoi(friendIDs)
	if !okUserID || !okFriendID || err != nil {
		http.Error(w, "Invalid data types", http.StatusBadRequest)
		return
	}
	response := map[string]string{}

	status, err := datab.Addfriend(userID, friendID)
	if err != nil {
		fmt.Println("error : ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status == "pending" {
		response["message"] = "Friend Request Already Sent!"
	} else if status == "accepted" {
		response["message"] = "You Are Already Friends!"
	} else {
		response["message"] = "Friend Request Sent!"
	}

	utils.JsonResponse(w, http.StatusOK, response)
}

func AcceptFriendHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Forbidden AcceptFR: Not authenticated", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodGet {
		// do stuff
		return
	}

	var req map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	userID, okUserID := session.Values["user_id"].(int)
	friendIDs, okFriendID := req["friend_id"].(string)
	choice, choiceOk := req["action"].(string)

	friendID, err := strconv.Atoi(friendIDs)
	if !okUserID || !okFriendID || !choiceOk || err != nil {
		http.Error(w, "Invalid data types", http.StatusBadRequest)
		return
	}
	response := map[string]string{"message": ""}
	if choice == "y" {
		err = datab.Acceptfriend(userID, friendID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response["message"] = "Friend Request Accepted!"
	} else {
		// here remove the invite from the database
		response["message"] = "Friend Request Refused"
	}
	utils.JsonResponse(w, http.StatusOK, response)
}

func DeleteFriendHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Forbidden DeleteFriend: Not authenticated", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodGet {
		// do stuff
		return
	}

	var req map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	userID, okUserID := session.Values["user_id"].(int)
	friendIDs, okFriendID := req["friend_id"].(string)
	friendID, err := strconv.Atoi(friendIDs)
	if !okUserID || !okFriendID || err != nil {
		http.Error(w, "Invalid data types", http.StatusBadRequest)
		return
	}
	err = datab.Deletefriend(userID, friendID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Friend deleted!"}
	utils.JsonResponse(w, http.StatusOK, response)
}

func LookUpFriendHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Forbidden LFriend: Not authenticated", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodGet {
		// do stuff
		return
	}

	var req map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	target, okTarget := req["user_name"].(string)
	if !okTarget {
		http.Error(w, "Invalid data types", http.StatusBadRequest)
		return
	}

	users, err := datab.LookupUser(target)
	if err != nil {
		fmt.Println("db error, ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if users == nil {
		users = []models.User{}
	}

	response := map[string]interface{}{
		"user_list": users,
	}
	utils.JsonResponse(w, http.StatusOK, response)
}

func InviteListHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Forbidden InvLst: Not authenticated", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodGet {
		userid, _ := session.Values["user_id"].(int)
		users, err := datab.GetInvites(userid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if users == nil {
			users = []models.User{}
		}
		response := map[string]interface{}{
			"user_list": users,
		}
		utils.JsonResponse(w, http.StatusOK, response)
	}

}
