package logging

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)

// responseWriter wraps http.ResponseWriter to capture status code and size
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

// Middleware creates HTTP logging middleware
func (l *Logger) Middleware(next http.Handler) http.Handler {
	if !l.enabled {
		return next
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Generate request ID
		requestID := generateRequestID()

		// Extract session ID from header or query param
		sessionID := r.Header.Get("Mcp-Session-Id")
		if sessionID == "" {
			sessionID = r.URL.Query().Get("sessionId")
		}

		// Log request start
		l.Info("HTTP request started",
			"request_id", requestID,
			"session_id", sessionID,
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
		)

		// Wrap response writer
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Call next handler
		next.ServeHTTP(wrapped, r)

		// Log request completion
		duration := time.Since(start)
		l.Info("HTTP request completed",
			"request_id", requestID,
			"session_id", sessionID,
			"method", r.Method,
			"path", r.URL.Path,
			"status", wrapped.statusCode,
			"size", wrapped.size,
			"duration_ms", duration.Milliseconds(),
		)
	})
}

// generateRequestID generates a random request ID
func generateRequestID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
