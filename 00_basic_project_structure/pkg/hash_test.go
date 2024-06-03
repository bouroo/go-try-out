package pkg_test

import (
	"go-try-out/00_basic_project_structure/pkg"
	"testing"
)

func TestHashPassword(t *testing.T) {
	// Test case: password is empty
	hashed, err := pkg.HashPassword("")
	if err == nil {
		t.Errorf("Expected error for empty password, got nil")
	}
	if hashed != "" {
		t.Errorf("Expected empty hashed password for empty password, got %s", hashed)
	}

	// Test case: password is not empty
	password := "mypassword"
	hashed, err = pkg.HashPassword(password)
	if err != nil {
		t.Errorf("Expected no error for valid password, got %v", err)
	}
	if hashed == "" {
		t.Errorf("Expected non-empty hashed password for valid password, got empty")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	// Test case: correct password
	password := "mypassword"
	hashed, _ := pkg.HashPassword(password)
	if !pkg.CheckPasswordHash(password, hashed) {
		t.Errorf("Expected true for correct password, got false")
	}

	// Test case: incorrect password
	if pkg.CheckPasswordHash("wrongpassword", hashed) {
		t.Errorf("Expected false for incorrect password, got true")
	}

	// Test case: empty password
	if pkg.CheckPasswordHash("", hashed) {
		t.Errorf("Expected false for empty password, got true")
	}

	// Test case: empty hashed password
	if pkg.CheckPasswordHash(password, "") {
		t.Errorf("Expected false for empty hashed password, got true")
	}
}
