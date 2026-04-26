package models

import (
	"crypto/rand"
	"encoding/hex"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type AdminSession struct {
	passwordHash []byte
	sessionToken string
	mu           sync.Mutex
}

var admin = &AdminSession{}

// InitAdmin sets the admin password hash from a plaintext password.
func InitAdmin(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.mu.Lock()
	defer admin.mu.Unlock()
	admin.passwordHash = hash
	return nil
}

// AdminLogin checks the password and returns a session token if valid.
func AdminLogin(password string) (string, bool) {
	admin.mu.Lock()
	defer admin.mu.Unlock()

	if err := bcrypt.CompareHashAndPassword(admin.passwordHash, []byte(password)); err != nil {
		return "", false
	}

	token := make([]byte, 32)
	rand.Read(token)
	admin.sessionToken = hex.EncodeToString(token)
	return admin.sessionToken, true
}

// AdminLogout clears the session.
func AdminLogout() {
	admin.mu.Lock()
	defer admin.mu.Unlock()
	admin.sessionToken = ""
}

// ValidateSession checks if the provided token matches the active session.
func ValidateSession(token string) bool {
	admin.mu.Lock()
	defer admin.mu.Unlock()
	return token != "" && token == admin.sessionToken
}
