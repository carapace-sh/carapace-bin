package gh

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionCodespaceViewFields completes codespace fields for view
//
//	displayName
//	repository
func ActionCodespaceViewFields() carapace.Action {
	return carapace.ActionValues(
		"billableOwner",
		"createdAt",
		"devcontainerPath",
		"displayName",
		"environmentId",
		"gitStatus",
		"idleTimeoutMinutes",
		"lastUsedAt",
		"location",
		"machineDisplayName",
		"machineName",
		"name",
		"owner",
		"prebuild",
		"recentFolders",
		"repository",
		"retentionExpiresAt",
		"retentionPeriodDays",
		"state",
		"vscsTarget",
	)
}

type codespace struct {
	Name       string
	Owner      string
	Repository string
	State      string
}

type codespacePort struct {
	Label      string
	SourcePort int
	Visibility string
}

// ActionCodespaces completes codespaces
//
//	compact-nightingale-tv87 (carapace-sh - carapace-bin - Active)
func ActionCodespaces() carapace.Action {
	return carapace.ActionExecCommand("gh", "codespace", "list", "--json", "name,owner,repository,state")(func(output []byte) carapace.Action {
		var codespaces []codespace
		if err := json.Unmarshal(output, &codespaces); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, codespace := range codespaces {
			vals = append(vals, codespace.Name, fmt.Sprintf("%v - %v - %v", codespace.Owner, codespace.Repository, codespace.State))
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("codespaces")
}

type CodespaceOpts struct {
	Codespace string
}

func (o CodespaceOpts) Default() CodespaceOpts {
	return o
}

// ActionCodespacePorts completes codespace ports
//
//	80 (nginx - private)
func ActionCodespacePorts(opts CodespaceOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"codespace", "ports", "--json", "label,sourcePort,visibility"}
		if opts.Codespace != "" {
			args = append(args, "--codespace", opts.Codespace)
		}
		return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
			var ports []codespacePort
			if err := json.Unmarshal(output, &ports); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, port := range ports {
				vals = append(vals, strconv.Itoa(port.SourcePort), fmt.Sprintf("%v - %v", port.Label, port.Visibility))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

type CodespacePathOpts struct {
	Codespace string
	Expand    bool
}

func (o CodespacePathOpts) Default() CodespacePathOpts {
	return o
}

// ActionCodespacePath completes paths in a codespace
//
//	remote:/home/codespace/bin
func ActionCodespacePath(opts CodespacePathOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path := filepath.Dir(c.Value)

		if !opts.Expand {
			// scp treats each filename argument as a shell expression,
			// subjecting it to expansion of environment variables, braces,
			// tilde, backticks, globs and so on. Because these present a
			// security risk (see https://lwn.net/Articles/835962/), we
			// disable them by shell-escaping the argument unless the user
			// provided the -e flag.
			path = `'` + strings.ReplaceAll(path, `'`, `'\''`) + `'`
		}

		args := []string{"codespace", "ssh", "--codespace", opts.Codespace, "--", "ls", "-1", "-p", path}
		if splitted := strings.Split(c.Value, "/"); strings.HasPrefix(splitted[len(splitted)-1], ".") {
			args = append(args, "-a") // show hidden
		}
		return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[:len(lines)-1]...)
			})
		}).StyleF(style.ForPathExt)
	})
}

// ActionCodespaceFields completes codespace fields
//
//	gitStatus
//	state
func ActionCodespaceFields() carapace.Action {
	return carapace.ActionValues(
		"createdAt",
		"gitStatus",
		"lastUsedAt",
		"name",
		"owner",
		"repository",
		"state",
	)
}

// ActionCodespacePortFields completes codespace port fields
//
//	sourcePort
//	visibility
func ActionCodespacePortFields() carapace.Action {
	return carapace.ActionValues( // carapace:curated
		"sourcePort",
		// "destinationPort", // TODO(mislav): this appears to always be blank?
		"visibility",
		"label",
		"browseUrl",
	)
}
