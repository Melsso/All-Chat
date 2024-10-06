package req_handlers

import (
	"encoding/json"
	"net/http"
	"All-Chat/back-end/datab"
	"All-Chat/back-end/models"
	"All-Chat/back-end/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusNoContent)
		return 
	}

	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/login.html")
		return
	
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var creds models.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := datab.GetUserByEmail(creds.Email)
	if err != nil {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	session, err := utils.Store.Get(r, "auth")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}
	session.Values["authenticated"] = true
	session.Values["user_id"] = user.UserID
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Login successful"}
	utils.JsonResponse(w, http.StatusOK, response)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusNoContent)
		return 
	}

	session, _ := utils.Store.Get(r, "auth")
	session.Values["authenticated"] = false;
	session.Options.MaxAge = -1
	session.Save(r, w)
	response := map[string]string{"message": "Log out successful."}
	utils.JsonResponse(w, http.StatusOK, response)
}
