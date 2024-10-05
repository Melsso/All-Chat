package req_handlers

import (
	"encoding/json"
	"net/http"
	"log"
	"All-Chat/back-end/datab"
	"All-Chat/back-end/models"
	"All-Chat/back-end/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
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

	session, err := utils.Store.Get(r, "session")
	if err != nil {
		log.Println("Failed to get session:", err)
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
	// After saving the session
	log.Printf("Session values after login: %v", session.Values)

	response := map[string]string{"message": "Login successful"}
	utils.JsonResponse(w, http.StatusOK, response)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.WriteHeader(http.StatusNoContent)
		return 
	}

	session, _ := utils.Store.Get(r, "session")
	session.Values["authenticated"] = false;
	session.Options.MaxAge = -1
	session.Save(r, w)
	w.Header().Set("Cache-Control", "no-store")
	http.Redirect(w, r, "/login.html", http.StatusSeeOther)
}
