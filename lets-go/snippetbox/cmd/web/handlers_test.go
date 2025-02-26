package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"snippetbox.netletic.com/internal/assert"
)

func TestPing_HandlerReturnsStatus200AndOK(t *testing.T) {
	t.Parallel()

	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(rr, r)

	rs := rr.Result()
	defer rs.Body.Close()
	want := http.StatusOK
	got := rs.StatusCode
	assert.Equal(t, want, got)

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	assert.Equal(t, "OK", string(body))
}

func TestPing_EndToEndReturnsStatus200AndOK(t *testing.T) {
	t.Parallel()
	app := newTestApplication(t)
	ts := newTestServer(t, app.Routes())
	defer ts.Close()
	code, _, body := ts.get(t, "/ping")
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "OK", string(body))
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.Routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "An old silent pond...",
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/snippet/view/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlPath:  "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/snippet/view/foo",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlPath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, _, gotBody := ts.get(t, tt.urlPath)

			assert.Equal(t, tt.wantCode, gotCode)

			if tt.wantBody != "" {
				assert.StringContains(t, tt.wantBody, gotBody)
			}
		})
	}
}
