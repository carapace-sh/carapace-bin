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
	if !ok || len(tools) != 3 {
		t.Fatalf("expected 3 tools, got %#v", result["tools"])
	}
	tool, ok := tools[0].(map[string]any)
	if !ok || tool["name"] != "complete_command" {
		t.Fatalf("expected complete_command tool, got %#v", tools[0])
	}
	tool1, ok := tools[1].(map[string]any)
	if !ok || tool1["name"] != "list_macros" {
		t.Fatalf("expected list_macros tool, got %#v", tools[1])
	}
	tool2, ok := tools[2].(map[string]any)
	if !ok || tool2["name"] != "codegen" {
		t.Fatalf("expected codegen tool, got %#v", tools[2])
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

func TestMCPCodegenMissingPath(t *testing.T) {
	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"codegen","arguments":{}}}`)

	var output bytes.Buffer
	s := NewMCPServer("", input, &output)
	if err := s.Run(); err != nil {
		t.Fatal(err)
	}

	var resp map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &resp); err != nil {
		t.Fatal(err)
	}
	result, ok := resp["result"].(map[string]any)
	if !ok {
		t.Fatalf("expected result, got: %#v", resp)
	}
	if result["isError"] != true {
		t.Fatalf("expected isError for missing path, got: %#v", result)
	}
}

func TestMCPCompleteCommandRequiresArgs(t *testing.T) {
	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"complete_command","arguments":{"args":[]}}}`)

	var output bytes.Buffer
	s := NewMCPServer("", input, &output)
	if err := s.Run(); err != nil {
		t.Fatal(err)
	}

	var resp map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &resp); err != nil {
		t.Fatal(err)
	}
	result, ok := resp["result"].(map[string]any)
	if !ok {
		t.Fatalf("expected result, got: %#v", resp)
	}
	if result["isError"] != true {
		t.Fatalf("expected error for empty args, got: %#v", result)
	}
}

func TestMCPCompleteCommandExecutableWithoutBridge(t *testing.T) {
	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"complete_command","arguments":{"args":["git",""],"executable":"/usr/bin/git"}}}`)

	var output bytes.Buffer
	s := NewMCPServer("", input, &output)
	if err := s.Run(); err != nil {
		t.Fatal(err)
	}

	var resp map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &resp); err != nil {
		t.Fatal(err)
	}
	result, ok := resp["result"].(map[string]any)
	if !ok {
		t.Fatalf("expected result, got: %#v", resp)
	}
	if result["isError"] != true {
		t.Fatalf("expected error for executable without bridge, got: %#v", result)
	}
	content := result["content"].([]any)
	text := content[0].(map[string]any)["text"].(string)
	if !strings.Contains(text, "bridge is required") {
		t.Fatalf("expected 'bridge is required' error, got: %v", text)
	}
}

func TestMCPCompleteCommandUnknownBridge(t *testing.T) {
	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"complete_command","arguments":{"args":["git",""],"bridge":"nonexistent"}}}`)

	var output bytes.Buffer
	s := NewMCPServer("", input, &output)
	if err := s.Run(); err != nil {
		t.Fatal(err)
	}

	var resp map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &resp); err != nil {
		t.Fatal(err)
	}
	result, ok := resp["result"].(map[string]any)
	if !ok {
		t.Fatalf("expected result, got: %#v", resp)
	}
	if result["isError"] != true {
		t.Fatalf("expected error for unknown bridge, got: %#v", result)
	}
}

func TestMCPCompleteCommandOldNameRejected(t *testing.T) {
	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"complete","arguments":{"args":["git",""]}}}`)

	var output bytes.Buffer
	s := NewMCPServer("", input, &output)
	if err := s.Run(); err != nil {
		t.Fatal(err)
	}

	var resp map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &resp); err != nil {
		t.Fatal(err)
	}
	if resp["error"] == nil {
		t.Fatalf("expected error for old tool name, got: %#v", resp)
	}
}

func TestMCPWithPathPrepended(t *testing.T) {
	env := []string{"PATH=/usr/bin:/bin", "HOME=/home/user"}
	result := withPathPrepended(env, "/custom/dir")

	found := false
	for _, e := range result {
		if strings.HasPrefix(e, "PATH=") {
			if e != "PATH=/custom/dir:/usr/bin:/bin" {
				t.Fatalf("expected PATH with prepended dir, got: %s", e)
			}
			found = true
		}
	}
	if !found {
		t.Fatal("PATH not found in result")
	}
}

// TODO: re-enable once complete_macro is properly implemented
// func TestMCPCompleteMacroRequiresMacro(t *testing.T) {
// 	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"complete_macro","arguments":{"args":[""]}}}`)
//
// 	var output bytes.Buffer
// 	s := NewMCPServer("", input, &output)
// 	if err := s.Run(); err != nil {
// 		t.Fatal(err)
// 	}
//
// 	var resp map[string]any
// 	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &resp); err != nil {
// 		t.Fatal(err)
// 	}
// 	result, ok := resp["result"].(map[string]any)
// 	if !ok {
// 		t.Fatalf("expected result, got: %#v", resp)
// 	}
// 	if result["isError"] != true {
// 		t.Fatalf("expected error for missing macro, got: %#v", result)
// 	}
// }
//
// func TestMCPCompleteMacroRequiresArgs(t *testing.T) {
// 	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"complete_macro","arguments":{"macro":"tools.git.Refs"}}}`)
//
// 	var output bytes.Buffer
// 	s := NewMCPServer("", input, &output)
// 	if err := s.Run(); err != nil {
// 		t.Fatal(err)
// 	}
//
// 	var resp map[string]any
// 	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &resp); err != nil {
// 		t.Fatal(err)
// 	}
// 	result, ok := resp["result"].(map[string]any)
// 	if !ok {
// 		t.Fatalf("expected result, got: %#v", resp)
// 	}
// 	if result["isError"] != true {
// 		t.Fatalf("expected error for missing args, got: %#v", result)
// 	}
// }
//
// func TestMCPCompleteMacroExecutableNotFound(t *testing.T) {
// 	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"complete_macro","arguments":{"macro":"tools.git.Refs","args":[""],"executable":"/nonexistent/path/carapace"}}}`)
//
// 	var output bytes.Buffer
// 	s := NewMCPServer("", input, &output)
// 	if err := s.Run(); err != nil {
// 		t.Fatal(err)
// 	}
//
// 	var resp map[string]any
// 	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &resp); err != nil {
// 		t.Fatal(err)
// 	}
// 	result, ok := resp["result"].(map[string]any)
// 	if !ok {
// 		t.Fatalf("expected result, got: %#v", resp)
// 	}
// 	if result["isError"] != true {
// 		t.Fatalf("expected error for nonexistent executable, got: %#v", result)
// 	}
// }

func TestMCPListMacrosExecutableNotFound(t *testing.T) {
	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"list_macros","arguments":{"executable":"/nonexistent/path/carapace"}}}`)

	var output bytes.Buffer
	s := NewMCPServer("", input, &output)
	if err := s.Run(); err != nil {
		t.Fatal(err)
	}

	var resp map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &resp); err != nil {
		t.Fatal(err)
	}
	result, ok := resp["result"].(map[string]any)
	if !ok {
		t.Fatalf("expected result, got: %#v", resp)
	}
	if result["isError"] != true {
		t.Fatalf("expected error for nonexistent executable, got: %#v", result)
	}
}
