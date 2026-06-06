package mcp

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace-bin/pkg/actions"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
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

type mcpCodegenRequest struct {
	Path string `json:"path"`
}

type mcpCompleteRequest struct {
	Args       []string `json:"args,omitempty"`
	Executable string   `json:"executable,omitempty"`
	Bridge     string   `json:"bridge,omitempty"`
}

type mcpCompleteMacroRequest struct {
	Macro      string   `json:"macro,omitempty"`
	Args       []string `json:"args,omitempty"`
	Executable string   `json:"executable,omitempty"`
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
					Name:        "complete_command",
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
							"executable": map[string]any{
								"type":        "string",
								"description": "path to the executable providing the completion (requires bridge)",
							},
							"bridge": map[string]any{
								"type":        "string",
								"description": "bridge providing the completion (e.g. carapace-bin, cobra, zsh, fish, bash, argcomplete, click)",
							},
						},
						"required":             []string{"args"},
						"additionalProperties": false,
					},
				},
				{
					Name:        "complete_macro",
					Description: "Return context‑aware, dynamic completions for a carapace macro.",
					InputSchema: map[string]any{
						"type": "object",
						"properties": map[string]any{
							"macro": map[string]any{
								"type":        "string",
								"description": "macro name (e.g. tools.git.Refs)",
							},
							"args": map[string]any{
								"type": "array",
								"items": map[string]any{
									"type": "string",
								},
								"description": "arguments to complete",
							},
							"executable": map[string]any{
								"type":        "string",
								"description": "path to the carapace executable providing the macro",
							},
						},
						"required":             []string{"macro", "args"},
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
				{
					Name:        "codegen",
					Description: "Generate Go code from a YAML spec file.",
					InputSchema: map[string]any{
						"type": "object",
						"properties": map[string]any{
							"path": map[string]any{
								"type":        "string",
								"description": "Path to the YAML spec file",
							},
						},
						"required":             []string{"path"},
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
	if call.Name != "complete_command" && call.Name != "complete_macro" && call.Name != "list_macros" && call.Name != "codegen" {
		return nil, fmt.Errorf("unknown tool %q", call.Name)
	}

	if call.Name == "list_macros" {
		return s.handleListMacros()
	}

	if call.Name == "codegen" {
		return s.handleCodegen(call.Arguments)
	}

	if call.Name == "complete_macro" {
		var request mcpCompleteMacroRequest
		if len(call.Arguments) != 0 {
			if err := json.Unmarshal(call.Arguments, &request); err != nil {
				return nil, fmt.Errorf("invalid complete_macro arguments: %w", err)
			}
		}

		completions, err := s.completeMacro(request)
		if err != nil {
			return mcpTextResult(err.Error(), true), nil
		}
		return mcpTextResult(completions, false), nil
	}

	var request mcpCompleteRequest
	if len(call.Arguments) != 0 {
		if err := json.Unmarshal(call.Arguments, &request); err != nil {
			return nil, fmt.Errorf("invalid complete_command arguments: %w", err)
		}
	}

	completions, err := s.completeCommand(request)
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

func (s *MCPServer) completeCommand(request mcpCompleteRequest) (string, error) {
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

	switch {
	case request.Executable == "" && request.Bridge == "":
		return s.completeDefault(request)
	case request.Executable == "" && request.Bridge != "":
		return s.completeBridge(request)
	case request.Executable != "" && request.Bridge == "":
		return "", errors.New("bridge is required when executable is set")
	default:
		return s.completeExecutable(request)
	}
}

func (s *MCPServer) completeDefault(request mcpCompleteRequest) (string, error) {
	args := []string{request.Args[0], "export"}
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

func (s *MCPServer) completeBridge(request mcpCompleteRequest) (string, error) {
	bridgeName := request.Bridge
	command := request.Args[0]

	// support carapace-bin/<bridge> syntax to use explicit bridge with carapace-bin
	if name, variant, ok := strings.Cut(bridgeName, "/"); ok && name == "carapace-bin" {
		bridgeName = variant
	}

	// verify bridge exists
	if _, ok := bridge.Get(bridgeName); !ok {
		return "", fmt.Errorf("unknown bridge: %s", bridgeName)
	}

	args := []string{command + "/" + bridgeName, "export"}
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

func (s *MCPServer) completeExecutable(request mcpCompleteRequest) (string, error) {
	executablePath := request.Executable
	bridgeName := request.Bridge
	command := request.Args[0]

	// resolve executable path
	resolved, err := resolveExecutable(executablePath)
	if err != nil {
		return "", err
	}

	// carapace-bin bridge: invoke the executable directly as a carapace binary
	// supports carapace-bin and carapace-bin/<bridge> syntax
	bridgeBase, bridgeVariant, hasVariant := strings.Cut(bridgeName, "/")
	if bridgeBase == "carapace-bin" {
		return s.completeExecutableCarapaceBin(request, resolved, hasVariant, bridgeVariant)
	}

	// verify bridge exists
	if _, ok := bridge.Get(bridgeName); !ok {
		return "", fmt.Errorf("unknown bridge: %s", bridgeName)
	}

	// for other bridges, invoke carapace with the bridge syntax
	// but prepend the executable's directory to PATH so the bridge can find it
	args := []string{command + "/" + bridgeName, "export"}
	args = append(args, request.Args...)

	carapaceExe, err := os.Executable()
	if err != nil {
		return "", err
	}

	cmd := exec.Command(carapaceExe, args...)
	cmd.Env = withPathPrepended(os.Environ(), filepath.Dir(resolved))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("completion failed: %w\n%s", err, strings.TrimSpace(string(output)))
	}
	return strings.TrimRight(string(output), "\r\n"), nil
}

func (s *MCPServer) completeExecutableCarapaceBin(request mcpCompleteRequest, resolved string, hasVariant bool, variant string) (string, error) {
	command := request.Args[0]

	if hasVariant && variant != "" {
		// carapace-bin/<bridge>: use the given executable but with explicit bridge
		args := []string{command + "/" + variant, "export"}
		args = append(args, request.Args...)

		cmd := exec.Command(resolved, args...)
		cmd.Env = os.Environ()
		output, err := cmd.CombinedOutput()
		if err != nil {
			return "", fmt.Errorf("completion failed: %w\n%s", err, strings.TrimSpace(string(output)))
		}
		return strings.TrimRight(string(output), "\r\n"), nil
	}

	// carapace-bin without variant: use default choice
	args := []string{command, "export"}
	args = append(args, request.Args...)

	cmd := exec.Command(resolved, args...)
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("completion failed: %w\n%s", err, strings.TrimSpace(string(output)))
	}
	return strings.TrimRight(string(output), "\r\n"), nil
}

func (s *MCPServer) completeMacro(request mcpCompleteMacroRequest) (string, error) {
	if request.Macro == "" {
		return "", errors.New("macro is required")
	}

	for _, arg := range request.Args {
		if strings.ContainsRune(arg, 0) {
			return "", errors.New("arguments must not contain NUL bytes")
		}
	}

	args := []string{"_carapace", "macro", request.Macro}
	args = append(args, request.Args...)

	executable, err := s.resolveMacroExecutable(request.Executable)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(executable, args...)
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("macro completion failed: %w\n%s", err, strings.TrimSpace(string(output)))
	}
	return strings.TrimRight(string(output), "\r\n"), nil
}

func (s *MCPServer) resolveMacroExecutable(executablePath string) (string, error) {
	if executablePath == "" {
		exe, err := os.Executable()
		if err != nil {
			return "", err
		}
		return exe, nil
	}
	return resolveExecutable(executablePath)
}

func resolveExecutable(path string) (string, error) {
	if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err != nil {
			return "", fmt.Errorf("executable not found: %s", path)
		}
		return path, nil
	}

	resolved, err := exec.LookPath(path)
	if err != nil {
		return "", fmt.Errorf("executable not found: %s", path)
	}
	return resolved, nil
}

func withPathPrepended(env []string, dir string) []string {
	result := make([]string, 0, len(env))
	found := false
	for _, e := range env {
		if rest, ok := strings.CutPrefix(e, "PATH="); ok {
			result = append(result, "PATH="+dir+":"+rest)
			found = true
		} else {
			result = append(result, e)
		}
	}
	if !found {
		result = append(result, "PATH="+dir)
	}
	return result
}

func (s *MCPServer) handleCodegen(args json.RawMessage) (map[string]any, error) {
	var request mcpCodegenRequest
	if err := json.Unmarshal(args, &request); err != nil {
		return nil, fmt.Errorf("invalid codegen arguments: %w", err)
	}
	if request.Path == "" {
		return nil, errors.New("path is required")
	}

	absPath := request.Path
	if !filepath.IsAbs(absPath) {
		abs, err := filepath.Abs(absPath)
		if err != nil {
			return mcpTextResult(err.Error(), true), nil
		}
		absPath = abs
	}

	executable, err := os.Executable()
	if err != nil {
		return mcpTextResult(err.Error(), true), nil
	}

	// Run codegen and capture the printed filenames
	cmd := exec.Command(executable, "--codegen", absPath)
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if err != nil {
		return mcpTextResult(fmt.Sprintf("codegen failed: %v\n%s", err, output), true), nil
	}

	var filenames []string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasSuffix(line, ".go") {
			filenames = append(filenames, line)
		}
	}

	if len(filenames) == 0 {
		return mcpTextResult("codegen completed but no files were generated", false), nil
	}

	sort.Strings(filenames)
	return map[string]any{
		"content": []map[string]any{
			{
				"type": "text",
				"text": toJSON(filenames),
			},
		},
	}, nil
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
