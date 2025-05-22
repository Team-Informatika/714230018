package main

import (
	"io"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func setupApp() *fiber.App {
	app := fiber.New()
	app.Get("/staffform", StaffForm)
	app.Post("/staffform", StaffFormSubmit)
	return app
}

func TestStaffFormGet(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/staffform", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send GET /staffform request: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestStaffFormPost(t *testing.T) {
	app := setupApp()

	form := url.Values{}
	form.Set("name", "John Doe")
	form.Set("position", "Manager")
	form.Set("email", "john@example.com")
	form.Set("phone", "1234567890")

	req := httptest.NewRequest("POST", "/staffform", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send POST /staffform request: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	bodyStr := string(body)

	if !strings.Contains(bodyStr, "John Doe") ||
		!strings.Contains(bodyStr, "Manager") ||
		!strings.Contains(bodyStr, "john@example.com") ||
		!strings.Contains(bodyStr, "1234567890") {
		t.Fatalf("Response body does not contain submitted data")
	}
}
