package devenv

import (
	"errors"
	"os/exec"
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionProcesses completes processes defined in devenv.nix
//
//	web
//	worker
func ActionProcesses(opts GlobalOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		file, err := configFile(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		return carapace.ActionExecCommand("devenv", opts.command("tasks", "list", "--json")...)(parseTasks(func(tasks []task) carapace.Action {
			vals := make([]string, 0)
			for _, t := range tasks {
				if t.Type == "process" {
					vals = append(vals, strings.TrimPrefix(t.Name, "devenv:processes:"), t.Description)
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})).Cache(5*time.Minute, opts.cacheKey(), key.FileStats(file))
	}).Tag("processes")
}

// ActionRunningProcesses completes processes managed by the running process manager
//
//	web (ready)
//	worker (running)
func ActionRunningProcesses(opts GlobalOpts) carapace.Action {
	return carapace.ActionExecCommandE("devenv", opts.command("processes", "list")...)(func(output []byte, err error) carapace.Action {
		if err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) && len(exitErr.Stderr) > 0 {
				return carapace.ActionMessage(strings.SplitN(strings.TrimSpace(string(exitErr.Stderr)), "\n", 2)[0])
			}
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for line := range strings.SplitSeq(string(output), "\n") {
			fields := strings.Fields(line)
			if len(fields) < 2 {
				continue
			}
			vals = append(vals, fields[0], fields[1], styleForPhase(fields[1]))
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("running processes")
}

func styleForPhase(s string) string {
	switch s {
	case "completed", "ready", "running":
		return style.Carapace.KeywordPositive
	case "error", "fatal", "terminating":
		return style.Carapace.KeywordNegative
	case "launching", "pending", "restarting":
		return style.Carapace.KeywordAmbiguous
	default:
		return style.Carapace.KeywordUnknown
	}
}
