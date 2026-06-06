package mcp

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func resolveExecutable(path string) (string, error) {
	if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err != nil {
			return "", err
		}
		return path, nil
	}

	resolved, err := exec.LookPath(path)
	if err != nil {
		return "", err
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

func runCommand(name string, args []string, env []string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Env = env
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w\n%s", err, strings.TrimSpace(string(output)))
	}
	return strings.TrimRight(string(output), "\r\n"), nil
}

func containsNUL(args []string) bool {
	for _, arg := range args {
		if strings.ContainsRune(arg, 0) {
			return true
		}
	}
	return false
}

func toJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}
