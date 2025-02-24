package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"snippetbox.netletic.com/internal/assert"
)

func TestInjectSecurityHeaders(t *testing.T) {
	t.Parallel()
	tests := []struct {
		requestHeader string
		want          string
	}{
		{
			requestHeader: "Content-Security-Policy",
			want:          "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com",
		},
		{
			requestHeader: "Referrer-Policy",
			want:          "origin-when-cross-origin",
		},
		{
			requestHeader: "X-Content-Type-Options",
			want:          "nosniff",
		},
		{
			requestHeader: "X-Frame-Options",
			want:          "deny",
		},
		{
			requestHeader: "X-XSS-Protection",
			want:          "0",
		},
	}

	// w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
	// w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
	// w.Header().Set("X-Content-Type-Options", "nosniff")
	// w.Header().Set("X-Frame-Options", "deny")
	// w.Header().Set("X-XSS-Protection", "0")

	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	injectSecurityHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()
	defer rs.Body.Close()

	for _, tt := range tests {
		got := rs.Header.Get(tt.requestHeader)
		assert.Equal(t, tt.want, got)
	}
	assert.Equal(t, rs.StatusCode, http.StatusOK)

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	assert.Equal(t, "OK", string(body))
}
