package req_handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"All-Chat/back-end/datab"
	"All-Chat/back-end/models"
	"All-Chat/back-end/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusNoContent)
		return 
	}

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
	
	utils.JsonResponse(w, http.StatusOK, response)
}
