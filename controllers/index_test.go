package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIndexHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Chdir("..")

	html, err := os.ReadFile("views/index.html")
	if err != nil {
		t.Fatalf("read index page: %v", err)
	}
	curlText, err := os.ReadFile("views/awacha.txt")
	if err != nil {
		t.Fatalf("read curl page: %v", err)
	}

	tests := []struct {
		name      string
		userAgent string
		wantBody  string
	}{
		{
			name:      "browser receives HTML",
			userAgent: "Mozilla/5.0",
			wantBody:  string(html),
		},
		{
			name:      "curl receives plain text",
			userAgent: "curl/8.15.0",
			wantBody:  string(curlText),
		},
		{
			name:      "curl must be at the start of the user agent",
			userAgent: "wrapper curl/8.15.0",
			wantBody:  string(html),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.LoadHTMLFiles("views/index.html")
			router.GET("/", IndexHandler())

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			request.Header.Set("User-Agent", tt.userAgent)
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			if response.Code != http.StatusOK {
				t.Errorf("status code = %d, want %d", response.Code, http.StatusOK)
			}
			if response.Body.String() != tt.wantBody {
				t.Errorf("response body did not match %q", tt.userAgent)
			}
		})
	}
}
