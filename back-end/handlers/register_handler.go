package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"All-Chat/back-end/datab"
	"All-Chat/back-end/models"
	"All-Chat/back-end/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/register.html")
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var regForm models.RegistrationForm

	err := json.NewDecoder(r.Body).Decode(&regForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	emailExists, err := datab.CheckEmailExists(regForm.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if emailExists {
		http.Error(w, "Email already in use", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regForm.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	regForm.Password = string(hashedPassword)

	dob, err := time.Parse("2006-01-02", regForm.DateOfBirth)
	if err != nil {
		http.Error(w, "Invalid date format for date of birth", http.StatusBadRequest)
		return
	}
	regForm.DateOfBirth = dob.Format("2006-01-02")

	result, err := datab.InsertUser(regForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"message": "Registration successful",
		"user_id": userID,
	}

	str := "session-name" + strconv.FormatInt(userID, 10)

	session, err := utils.Store.Get(r, str)
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}
	session.Values["authenticated"] = true
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}
	utils.JsonResponse(w, http.StatusOK, response)
}
