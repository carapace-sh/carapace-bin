package mcp

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestMCPInitializeAndListTools(t *testing.T) {
	input := strings.NewReader(strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}`,
	}, "\n"))

	var output bytes.Buffer
	s := NewMCPServer("", input, &output)
	if err := s.Run(); err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 responses, got %d", len(lines))
	}

	var resp1 map[string]any
	if err := json.Unmarshal([]byte(lines[0]), &resp1); err != nil {
		t.Fatal(err)
	}
	if resp1["error"] != nil {
		t.Fatalf("initialize returned error: %#v", resp1["error"])
	}

	var resp2 map[string]any
	if err := json.Unmarshal([]byte(lines[1]), &resp2); err != nil {
		t.Fatal(err)
	}
	result, ok := resp2["result"].(map[string]any)
	if !ok {
		t.Fatalf("tools/list result missing")
	}
	tools, ok := result["tools"].([]any)
	if !ok || len(tools) != 2 {
		t.Fatalf("expected 2 tools, got %#v", result["tools"])
	}
	tool, ok := tools[0].(map[string]any)
	if !ok || tool["name"] != "complete" {
		t.Fatalf("expected complete tool, got %#v", tools[0])
	}
	tool1, ok := tools[1].(map[string]any)
	if !ok || tool1["name"] != "list_macros" {
		t.Fatalf("expected list_macros tool, got %#v", tools[1])
	}
}

func TestMCPBatchSkipsNotifications(t *testing.T) {
	input := strings.NewReader(`[{"jsonrpc":"2.0","method":"notifications/initialized"},{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}]`)

	var output bytes.Buffer
	s := NewMCPServer("", input, &output)
	if err := s.Run(); err != nil {
		t.Fatal(err)
	}

	var responses []map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &responses); err != nil {
		t.Fatalf("invalid batch response: %v", err)
	}
	if len(responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(responses))
	}
	if responses[0]["id"].(float64) != 1 {
		t.Fatalf("unexpected response id: %#v", responses[0]["id"])
	}
}
