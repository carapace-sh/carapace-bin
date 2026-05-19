package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "run MCP server",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMCPServer(os.Stdin, os.Stdout, completeWithCarapace)
	},
}

var mcpServerVersion = "dev"

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

type mcpCompleter func(mcpCompleteRequest) (string, error)

func runMCPServer(input io.Reader, output io.Writer, complete mcpCompleter) error {
	decoder := json.NewDecoder(input)
	encoder := json.NewEncoder(output)

	for {
		var message json.RawMessage
		if err := decoder.Decode(&message); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		if err := writeMCPResponse(encoder, message, complete); err != nil {
			return err
		}
	}
}

func writeMCPResponse(encoder *json.Encoder, message json.RawMessage, complete mcpCompleter) error {
	switch firstNonSpace(message) {
	case '[':
		var requests []mcpRequest
		if err := json.Unmarshal(message, &requests); err != nil {
			return err
		}
		responses := make([]mcpResponse, 0, len(requests))
		for _, request := range requests {
			response, ok := handleMCPRequest(request, complete)
			if ok {
				responses = append(responses, response)
			}
		}
		if len(responses) == 0 {
			return nil
		}
		return encoder.Encode(responses)
	default:
		var request mcpRequest
		if err := json.Unmarshal(message, &request); err != nil {
			return err
		}
		response, ok := handleMCPRequest(request, complete)
		if !ok {
			return nil
		}
		return encoder.Encode(response)
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

func handleMCPRequest(request mcpRequest, complete mcpCompleter) (mcpResponse, bool) {
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
				"version": mcpServerVersion,
			},
		}
	case "tools/list":
		response.Result = map[string]any{
			"tools": []mcpTool{
				{
					Name: "complete",
					Description: "Return context‑aware, dynamic completions for shell commands.\n" +
						"Use it:\n" +
						"• Whenever you need to discover or list the available flags/options for a command (e.g., `git show`).\n" +
						"• Provide interactive autocompletion for user‑entered commands in the CLI, ensuring correct spelling and flag syntax.\n" +
						"• When building scripts or tools that require prompt, reliable completion data without hard‑coding flag lists.\n",
					InputSchema: map[string]any{
						"type": "object",
						"properties": map[string]any{
							"args": map[string]any{
								"type": "array",
								"items": map[string]any{
									"type": "string",
								},
								"description": "command line arguments. use `\"\"` as last argument to complete positionals and `-` for flags",
							},
						},
						"required":             []string{"args"},
						"additionalProperties": false,
					},
				},
			},
		}
	case "tools/call":
		result, err := handleMCPToolCall(request.Params, complete)
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

func handleMCPToolCall(params json.RawMessage, complete mcpCompleter) (map[string]any, error) {
	var call mcpToolCallParams
	if err := json.Unmarshal(params, &call); err != nil {
		return nil, fmt.Errorf("invalid tool call parameters: %w", err)
	}
	if call.Name != "complete" {
		return nil, fmt.Errorf("unknown tool %q", call.Name)
	}

	var request mcpCompleteRequest
	if len(call.Arguments) != 0 {
		if err := json.Unmarshal(call.Arguments, &request); err != nil {
			return nil, fmt.Errorf("invalid complete arguments: %w", err)
		}
	}

	completions, err := complete(request)
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

func completeWithCarapace(request mcpCompleteRequest) (string, error) {
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
	args := []string{command, "export", command}
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
