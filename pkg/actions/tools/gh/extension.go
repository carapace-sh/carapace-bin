package gh

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
)

// _extensionsDir is the gh extensions data path following the
// gh CLI directory precedence (XDG_DATA_HOME, LocalAppData on windows, HOME).
func extensionsDir() string {
	var path string
	if a := os.Getenv("XDG_DATA_HOME"); a != "" {
		path = filepath.Join(a, "gh")
	} else if b := os.Getenv("LocalAppData"); runtime.GOOS == "windows" && b != "" {
		path = filepath.Join(b, "GitHub CLI")
	} else {
		c, _ := os.UserHomeDir()
		path = filepath.Join(c, ".local", "share", "gh")
	}
	return path
}

// Extensions returns names of installed gh extensions.
func Extensions() ([]string, error) {
	files, err := os.ReadDir(extensionsDir() + "/extensions")
	if err != nil {
		return nil, err
	}

	vals := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			vals = append(vals, strings.TrimPrefix(file.Name(), "gh-"))
		}
	}
	return vals, nil
}

// ActionExtensions completes installed extensions
//
//	extension-a
//	extension-b
func ActionExtensions() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		extensions, err := Extensions()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionValues(extensions...)
	}).Tag("extensions")
}

// ActionTopExtensions completes popular extensions
//
//	extension-a (12 stars - description)
func ActionTopExtensions(opts HostOpts) carapace.Action {
	args := []string{"api", "--preview", "mercy",
		fmt.Sprintf(`search/repositories?per_page=100&q=%v`, url.QueryEscape("topic:gh-extension sort:stars-desc")),
	}
	if opts.Host != "" {
		args = append(args, "--hostname", opts.Host)
	}
	return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
		var queryResult struct {
			Items []struct {
				FullName        string `json:"full_name"`
				Description     string
				StargazersCount int `json:"stargazers_count"`
			}
		}
		if err := json.Unmarshal(output, &queryResult); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, item := range queryResult.Items {
			vals = append(vals, item.FullName, fmt.Sprintf("%v - %v", item.StargazersCount, item.Description))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
