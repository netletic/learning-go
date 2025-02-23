package main

type contextKey string

const (
	traceIDContextKey         = contextKey("TraceID")
	isAuthenticatedContextKey = contextKey("isAuthenticated")
)
