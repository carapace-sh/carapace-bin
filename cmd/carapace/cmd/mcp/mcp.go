package mcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace-bin/pkg/actions"
)

type MCPServer struct {
	version string
	stdin   io.Reader
	stdout  io.Writer
}

func NewMCPServer(version string, stdin io.Reader, stdout io.Writer) *MCPServer {
	return &MCPServer{
		version: version,
		stdin:   stdin,
		stdout:  stdout,
	}
}

type mcpRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type mcpResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Result  any             `json:"result,omitempty"`
	Error   *mcpError       `json:"error,omitempty"`
}

type mcpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type mcpTool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}

type mcpToolCallParams struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments,omitempty"`
}

type mcpCompleteRequest struct {
	Args []string `json:"args,omitempty"`
}

func (s *MCPServer) Run() error {
	decoder := json.NewDecoder(s.stdin)
	encoder := json.NewEncoder(s.stdout)

	for {
		var message json.RawMessage
		if err := decoder.Decode(&message); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		result, err := s.processMessage(message)
		if err != nil {
			return err
		}
		if result != nil {
			if err := encoder.Encode(result); err != nil {
				return err
			}
		}
	}
}

func (s *MCPServer) processMessage(message json.RawMessage) (any, error) {
	switch firstNonSpace(message) {
	case '[':
		var requests []mcpRequest
		if err := json.Unmarshal(message, &requests); err != nil {
			return nil, err
		}
		responses := make([]mcpResponse, 0, len(requests))
		for _, request := range requests {
			response, ok := s.handleRequest(request)
			if ok {
				responses = append(responses, response)
			}
		}
		if len(responses) == 0 {
			return nil, nil
		}
		return responses, nil
	default:
		var request mcpRequest
		if err := json.Unmarshal(message, &request); err != nil {
			return nil, err
		}
		response, ok := s.handleRequest(request)
		if !ok {
			return nil, nil
		}
		return response, nil
	}
}

func firstNonSpace(message []byte) byte {
	for _, b := range message {
		switch b {
		case ' ', '\n', '\r', '\t':
			continue
		default:
			return b
		}
	}
	return 0
}

func (s *MCPServer) handleRequest(request mcpRequest) (mcpResponse, bool) {
	if len(request.ID) == 0 {
		return mcpResponse{}, false
	}

	response := mcpResponse{
		JSONRPC: "2.0",
		ID:      request.ID,
	}

	switch request.Method {
	case "initialize":
		response.Result = map[string]any{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]any{
				"tools": map[string]any{},
			},
			"serverInfo": map[string]any{
				"name":    "carapace",
				"version": s.version,
			},
		}
	case "tools/list":
		response.Result = map[string]any{
			"tools": []mcpTool{
				{
					Name:        "complete",
					Description: "Return context‑aware, dynamic completions for shell commands.",
					InputSchema: map[string]any{
						"type": "object",
						"properties": map[string]any{
							"args": map[string]any{
								"type": "array",
								"items": map[string]any{
									"type": "string",
								},
								"description": "command line arguments",
							},
						},
						"required":             []string{"args"},
						"additionalProperties": false,
					},
				},
				{
					Name:        "list_macros",
					Description: "List available macros and their signatures.",
					InputSchema: map[string]any{
						"type":                 "object",
						"properties":           map[string]any{},
						"additionalProperties": false,
					},
				},
			},
		}
	case "tools/call":
		result, err := s.handleToolCall(request.Params)
		if err != nil {
			response.Error = &mcpError{Code: -32602, Message: err.Error()}
			return response, true
		}
		response.Result = result
	default:
		response.Error = &mcpError{Code: -32601, Message: "method not found"}
	}

	return response, true
}

func (s *MCPServer) handleToolCall(params json.RawMessage) (map[string]any, error) {
	var call mcpToolCallParams
	if err := json.Unmarshal(params, &call); err != nil {
		return nil, fmt.Errorf("invalid tool call parameters: %w", err)
	}
	if call.Name != "complete" && call.Name != "list_macros" {
		return nil, fmt.Errorf("unknown tool %q", call.Name)
	}

	if call.Name == "list_macros" {
		return s.handleListMacros()
	}

	var request mcpCompleteRequest
	if len(call.Arguments) != 0 {
		if err := json.Unmarshal(call.Arguments, &request); err != nil {
			return nil, fmt.Errorf("invalid complete arguments: %w", err)
		}
	}

	completions, err := s.complete(request)
	if err != nil {
		return mcpTextResult(err.Error(), true), nil
	}
	return mcpTextResult(completions, false), nil
}

func mcpTextResult(text string, isError bool) map[string]any {
	result := map[string]any{
		"content": []map[string]string{
			{
				"type": "text",
				"text": text,
			},
		},
	}
	if isError {
		result["isError"] = true
	}
	return result
}

func (s *MCPServer) complete(request mcpCompleteRequest) (string, error) {
	switch len(request.Args) {
	case 0:
		return "", errors.New("command is required")
	case 1:
		return "", errors.New("argument to complete is required")
	}

	command := strings.TrimSpace(request.Args[0])
	if strings.HasPrefix(command, "-") {
		return "", errors.New("command must be a completer name")
	}

	for _, arg := range request.Args {
		if strings.ContainsRune(arg, 0) {
			return "", errors.New("arguments must not contain NUL bytes")
		}
	}
	args := []string{command, "export"}
	args = append(args, request.Args...)

	executable, err := os.Executable()
	if err != nil {
		return "", err
	}

	cmd := exec.Command(executable, args...)
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("completion failed: %w\n%s", err, strings.TrimSpace(string(output)))
	}
	return strings.TrimRight(string(output), "\r\n"), nil
}

func (s *MCPServer) handleListMacros() (map[string]any, error) {
	var names []string
	for name := range actions.Macros {
		names = append(names, name)
	}
	sort.Strings(names)

	macros := make([]map[string]string, 0, len(names))
	for _, name := range names {
		m := actions.Macros[name]
		sig := m.Signature()
		if sig == "" {
			sig = "—"
		}
		macros = append(macros, map[string]string{
			"name":        "carapace." + name,
			"signature":   sig,
			"description": m.Description,
			"reference":   m.Function,
		})
	}

	return map[string]any{
		"content": []map[string]any{
			{
				"type": "text",
				"text": toJSON(macros),
			},
		},
	}, nil
}

func toJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}
