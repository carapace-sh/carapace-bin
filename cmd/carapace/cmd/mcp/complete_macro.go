package mcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type mcpCompleteMacroRequest struct {
	Macro      string   `json:"macro,omitempty"`
	Args       []string `json:"args,omitempty"`
	Executable string   `json:"executable,omitempty"`
}

func (s *MCPServer) handleCompleteMacro(args json.RawMessage) (map[string]any, error) {
	var request mcpCompleteMacroRequest
	if len(args) != 0 {
		if err := json.Unmarshal(args, &request); err != nil {
			return nil, fmt.Errorf("invalid complete_macro arguments: %w", err)
		}
	}

	completions, err := s.completeMacro(request)
	if err != nil {
		return mcpTextResult(err.Error(), true), nil
	}
	return mcpTextResult(completions, false), nil
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

	return runCommand(executable, args, os.Environ())
}

func (s *MCPServer) resolveMacroExecutable(executablePath string) (string, error) {
	if executablePath == "" {
		return selfExecutable()
	}
	return resolveExecutable(executablePath)
}
