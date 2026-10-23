package hooks

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestOnSuccessLogsResultOnlyAtDebug(t *testing.T) {
	result := map[string]string{"data": "payload"}
	resultJSON := toJSON(result)

	tests := []struct {
		name       string
		level      slog.Level
		wantResult bool
	}{
		{name: "info", level: slog.LevelInfo, wantResult: false},
		{name: "debug", level: slog.LevelDebug, wantResult: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			prev := slog.Default()
			slog.SetDefault(slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: tt.level})))
			defer slog.SetDefault(prev)

			h := NewLoggerHooks()
			for _, fn := range h.OnSuccess {
				fn(context.Background(), 1, mcp.MethodToolsCall, nil, result)
			}

			var entry map[string]any
			if err := json.Unmarshal(buf.Bytes(), &entry); err != nil {
				t.Fatalf("cannot parse log line %q: %s", buf.String(), err)
			}
			if got := entry["result_bytes"]; got != float64(len(resultJSON)) {
				t.Errorf("result_bytes = %v, want %d", got, len(resultJSON))
			}
			_, hasResult := entry["result"]
			if hasResult != tt.wantResult {
				t.Errorf("result present = %v, want %v", hasResult, tt.wantResult)
			}
		})
	}
}
