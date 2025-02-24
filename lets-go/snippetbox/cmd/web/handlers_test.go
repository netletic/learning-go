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
