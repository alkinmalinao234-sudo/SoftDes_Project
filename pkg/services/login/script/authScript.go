package AUTHscript

import (
	"fmt"

	DBConnection "template_school/pkg/middleware/databaseConnection"
	AUTHmodels "template_school/pkg/services/login/models"

	"golang.org/x/crypto/bcrypt"
)

func AuthScriptLogin(req AUTHmodels.RequestBody) (string, error) {

	// 1. Validate input
	if req.Login == "" || req.Password == "" {
		return "", fmt.Errorf("login and password are required")
	}

	// 2. Get DB
	db := DBConnection.GetDB()
	if db == nil {
		return "", fmt.Errorf("database not connected")
	}

	// 3. Get user from database (email OR username supported)
	var user struct {
		ID           int
		Username     string
		Email        string
		PasswordHash string
	}

	err := db.Raw(
		"SELECT id, username, email, password_hash FROM users WHERE email = ? OR username = ?",
		req.Login,
		req.Login,
	).Scan(&user).Error

	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	// 4. Check password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return "", fmt.Errorf("wrong password")
	}

	// 5. Success response
	return "success login", nil
}
