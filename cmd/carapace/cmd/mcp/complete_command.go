package mcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
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

	if containsNUL(request.Args) {
		return "", errors.New("arguments must not contain NUL bytes")
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
	executable, err := os.Executable()
	if err != nil {
		return "", err
	}
	return runCommand(executable, exportArgs(request.Args[0], "", request.Args), os.Environ())
}

func (s *MCPServer) completeBridge(request mcpCompleteRequest) (string, error) {
	bridgeName := resolveBridgeName(request.Bridge)

	if _, ok := bridge.Get(bridgeName); !ok {
		return "", fmt.Errorf("unknown bridge: %s", bridgeName)
	}

	executable, err := os.Executable()
	if err != nil {
		return "", err
	}
	return runCommand(executable, exportArgs(request.Args[0], "/"+bridgeName, request.Args), os.Environ())
}

func (s *MCPServer) completeExecutable(request mcpCompleteRequest) (string, error) {
	resolved, err := resolveExecutable(request.Executable)
	if err != nil {
		return "", err
	}

	bridgeBase, bridgeVariant, hasVariant := strings.Cut(request.Bridge, "/")
	if bridgeBase == "carapace-bin" {
		suffix := ""
		if hasVariant && bridgeVariant != "" {
			suffix = "/" + bridgeVariant
		}
		return runCommand(resolved, exportArgs(request.Args[0], suffix, request.Args), os.Environ())
	}

	bridgeName := resolveBridgeName(request.Bridge)
	actionFactory, ok := bridge.Get(bridgeName)
	if !ok {
		return "", fmt.Errorf("unknown bridge: %s", bridgeName)
	}

	action := actionFactory(resolved)
	c := carapace.NewContext(request.Args[1:]...)
	invoked := action.Invoke(c)
	output, err := json.Marshal(invoked)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func resolveBridgeName(bridgeName string) string {
	if name, variant, ok := strings.Cut(bridgeName, "/"); ok && name == "carapace-bin" {
		return variant
	}
	return bridgeName
}

func exportArgs(command string, suffix string, args []string) []string {
	result := []string{command + suffix, "export"}
	result = append(result, args...)
	return result
}
