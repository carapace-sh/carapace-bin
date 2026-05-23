package mcp

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestMCPInitializeAndListTools(t *testing.T) {
	input := strings.NewReader(strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}`,
	}, "\n"))

	var output bytes.Buffer
	if err := runMCPServer(input, &output, func(mcpCompleteRequest) (string, error) {
		t.Fatal("complete should not be called")
		return "", nil
	}); err != nil {
		t.Fatal(err)
	}

	responses := decodeMCPResponses(t, output.String())
	if len(responses) != 2 {
		t.Fatalf("expected 2 responses, got %d", len(responses))
	}
	if responses[0]["error"] != nil {
		t.Fatalf("initialize returned error: %#v", responses[0]["error"])
	}

	result, ok := responses[1]["result"].(map[string]any)
	if !ok {
		t.Fatalf("tools/list result missing")
	}
	tools, ok := result["tools"].([]any)
	if !ok || len(tools) != 1 {
		t.Fatalf("expected one tool, got %#v", result["tools"])
	}
	tool, ok := tools[0].(map[string]any)
	if !ok || tool["name"] != "complete" {
		t.Fatalf("expected complete tool, got %#v", tools[0])
	}
}

func TestMCPCompleteToolCall(t *testing.T) {
	input := strings.NewReader(`{"jsonrpc":"2.0","id":"call-1","method":"tools/call","params":{"name":"complete","arguments":{"args":["git", "checko"]}}}`)

	var output bytes.Buffer
	if err := runMCPServer(input, &output, func(request mcpCompleteRequest) (string, error) {
		if expected := []string{"git", "checko"}; !reflect.DeepEqual(expected, request.Args) {
			t.Fatalf("expected %#v checko, got %#v", expected, request.Args)
		}
		return `[{"value":"checkout","description":"Switch branches or restore working tree files"}]`, nil
	}); err != nil {
		t.Fatal(err)
	}

	responses := decodeMCPResponses(t, output.String())
	if len(responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(responses))
	}
	result, ok := responses[0]["result"].(map[string]any)
	if !ok {
		t.Fatalf("tools/call result missing")
	}
	content, ok := result["content"].([]any)
	if !ok || len(content) != 1 {
		t.Fatalf("expected one content item, got %#v", result["content"])
	}
	item, ok := content[0].(map[string]any)
	if !ok || !strings.Contains(item["text"].(string), "checkout") {
		t.Fatalf("unexpected tool content: %#v", content[0])
	}
}

func TestMCPBatchSkipsNotifications(t *testing.T) {
	input := strings.NewReader(`[{"jsonrpc":"2.0","method":"notifications/initialized"},{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}]`)

	var output bytes.Buffer
	if err := runMCPServer(input, &output, func(mcpCompleteRequest) (string, error) {
		t.Fatal("complete should not be called")
		return "", nil
	}); err != nil {
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

func decodeMCPResponses(t *testing.T, output string) []map[string]any {
	t.Helper()

	lines := strings.Split(strings.TrimSpace(output), "\n")
	responses := make([]map[string]any, 0, len(lines))
	for _, line := range lines {
		var response map[string]any
		if err := json.Unmarshal([]byte(line), &response); err != nil {
			t.Fatalf("invalid response %q: %v", line, err)
		}
		responses = append(responses, response)
	}
	return responses
}
