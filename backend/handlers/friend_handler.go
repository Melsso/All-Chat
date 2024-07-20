package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"playground/datab"
	"playground/models"
	"playground/utils"
	"strconv"
)

func AddFriendHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session")
	_, ok := session.Values["user_id"].(int)	
	if (!ok) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return 
	}

	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
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
	if !okUserID || ! okFriendID || err != nil{
		http.Error(w, "Invalid data types", http.StatusBadRequest)
		return 
	}
	err = datab.Addfriend(userID, friendID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	response := map[string]string{"message": "Friend Request Sent"}
	utils.JsonResponse(w, http.StatusOK, response)
}

func AcceptFriendHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session")
	_, ok := session.Values["user_id"].(int)	
	if (!ok) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return 
	}
	
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
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
	if !okUserID || !okFriendID || !choiceOk || err != nil{
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
	session, _ := utils.Store.Get(r, "session")
	_, ok := session.Values["user_id"].(int)	
	if (!ok) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return 
	}

	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
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
	if !okUserID || ! okFriendID  || err != nil {
		http.Error(w, "Invalid data types", http.StatusBadRequest)
		return 
	}
	err = datab.Deletefriend(userID, friendID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	response := map[string]string{"message": "Friend Request Sent"}
	utils.JsonResponse(w, http.StatusOK, response)
}

func LookUpFriendHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session")
	_, ok := session.Values["user_id"].(int)	
	if (!ok) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return 
	}

	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
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
	session, _ := utils.Store.Get(r, "session")
	_, ok := session.Values["user_id"].(int)	
	if (!ok) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return 
	}

	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return 
	}

	// get invite list through Get method, post method will involve accepting/refusing friend request
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

