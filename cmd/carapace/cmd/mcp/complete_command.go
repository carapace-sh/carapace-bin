package mcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
)

type mcpCompleteRequest struct {
	Args       []string `json:"args,omitempty"`
	Executable string   `json:"executable,omitempty"`
	Bridge     string   `json:"bridge,omitempty"`
}

func (s *MCPServer) handleCompleteCommand(args json.RawMessage) (map[string]any, error) {
	var request mcpCompleteRequest
	if len(args) != 0 {
		if err := json.Unmarshal(args, &request); err != nil {
			return nil, fmt.Errorf("invalid complete_command arguments: %w", err)
		}
	}

	completions, err := s.completeCommand(request)
	if err != nil {
		return mcpTextResult(err.Error(), true), nil
	}
	return mcpTextResult(completions, false), nil
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

	executable, err := selfExecutable()
	if err != nil {
		return "", err
	}

	return runCommand(executable, args, os.Environ())
}

func (s *MCPServer) completeBridge(request mcpCompleteRequest) (string, error) {
	bridgeName := request.Bridge
	command := request.Args[0]

	if name, variant, ok := strings.Cut(bridgeName, "/"); ok && name == "carapace-bin" {
		bridgeName = variant
	}

	if _, ok := bridge.Get(bridgeName); !ok {
		return "", fmt.Errorf("unknown bridge: %s", bridgeName)
	}

	args := []string{command + "/" + bridgeName, "export"}
	args = append(args, request.Args...)

	executable, err := selfExecutable()
	if err != nil {
		return "", err
	}

	return runCommand(executable, args, os.Environ())
}

func (s *MCPServer) completeExecutable(request mcpCompleteRequest) (string, error) {
	bridgeName := request.Bridge
	command := request.Args[0]

	resolved, err := resolveExecutable(request.Executable)
	if err != nil {
		return "", err
	}

	bridgeBase, bridgeVariant, hasVariant := strings.Cut(bridgeName, "/")
	if bridgeBase == "carapace-bin" {
		return s.completeExecutableCarapaceBin(request, resolved, hasVariant, bridgeVariant)
	}

	if _, ok := bridge.Get(bridgeName); !ok {
		return "", fmt.Errorf("unknown bridge: %s", bridgeName)
	}

	args := []string{command + "/" + bridgeName, "export"}
	args = append(args, request.Args...)

	carapaceExe, err := selfExecutable()
	if err != nil {
		return "", err
	}

	return runCommand(carapaceExe, args, withPathPrepended(os.Environ(), filepath.Dir(resolved)))
}

func (s *MCPServer) completeExecutableCarapaceBin(request mcpCompleteRequest, resolved string, hasVariant bool, variant string) (string, error) {
	command := request.Args[0]

	if hasVariant && variant != "" {
		args := []string{command + "/" + variant, "export"}
		args = append(args, request.Args...)
		return runCommand(resolved, args, os.Environ())
	}

	args := []string{command, "export"}
	args = append(args, request.Args...)
	return runCommand(resolved, args, os.Environ())
}
