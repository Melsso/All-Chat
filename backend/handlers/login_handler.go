package handlers

import (
	"encoding/json"
	"net/http"

	"playground/datab"
	"playground/models"
	"playground/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	session, _ := utils.Store.Get(r, "session")
	session.Values["user_id"] = user.UserID
	session.Values["authenticated"] = true;
	session.Save(r, w)

	response := map[string]string{"message": "Login successful"}
	utils.JsonResponse(w, http.StatusOK, response)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := utils.Store.Get(r, "session")

    session.Options.MaxAge = -1
    session.Save(r, w)
	w.Header().Set("Cache-Control", "no-store")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")
    http.Redirect(w, r, "/login.html", http.StatusSeeOther)
}