package devenv

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

type task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// TaskOpts filters which tasks are completed.
type TaskOpts struct {
	GlobalOpts

	// Oneshot includes tasks that run to completion
	Oneshot bool

	// Process includes tasks that manage a long-running process
	Process bool
}

func (o TaskOpts) Default() TaskOpts {
	o.Oneshot = true
	o.Process = true
	return o
}

func (o TaskOpts) cacheKey() key.Key {
	return key.String(
		strconv.FormatBool(o.Oneshot),
		strconv.FormatBool(o.Process),
	)
}

// ActionTasks completes tasks.
// Oneshot and Process select which task types are included.
//
//	devenv:enterShell (Runs when entering the shell)
//	myproj:setup (Prepare the workspace)
func ActionTasks(opts TaskOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		file, err := configFile(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		return carapace.ActionExecCommand("devenv", opts.command("tasks", "list", "--json")...)(parseTasks(func(tasks []task) carapace.Action {
			vals := make([]string, 0)
			for _, t := range tasks {
				switch t.Type {
				case "process":
					if !opts.Process {
						continue
					}
				default:
					if !opts.Oneshot {
						continue
					}
				}
				vals = append(vals, t.Name, t.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		})).Cache(5*time.Minute, opts.cacheKey(), opts.GlobalOpts.cacheKey(), key.FileStats(file))
	}).Tag("tasks")
}

func parseTasks(f func(tasks []task) carapace.Action) func(output []byte) carapace.Action {
	return func(output []byte) carapace.Action {
		var tasks []task
		if err := json.Unmarshal(output, &tasks); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(tasks)
	}
}
