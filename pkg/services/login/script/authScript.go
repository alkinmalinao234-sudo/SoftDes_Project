package AUTHscript

import (
	"fmt"

	DBConnection "template_school/pkg/middleware/databaseConnection"
	AUTHmodels "template_school/pkg/services/login/models"
)

func AuthScriptLogin(req AUTHmodels.RequestBody) (string, error) {

	// 1. Validate input
	if req.Login == "" || req.Password == "" {
		return "", fmt.Errorf("login and password are required")
	}

	// 2. DB connection
	db := DBConnection.GetDB()
	if db == nil {
		return "", fmt.Errorf("database not connected")
	}

	// 3. Get user (no bcrypt needed)
	var user struct {
		ID       int
		Username string
		Email    string
		Password string
	}

	err := db.Raw(
		"SELECT id, username, email, password_hash FROM users WHERE email = ? OR username = ?",
		req.Login,
		req.Login,
	).Scan(&user).Error

	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	// 4. SIMPLE PASSWORD CHECK (NO HASH)
	if user.Password != req.Password {
		return "", fmt.Errorf("wrong password")
	}

	return "success login", nil
}
