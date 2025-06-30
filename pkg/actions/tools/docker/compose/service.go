package compose

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type config struct {
	Services map[string]struct {
		Image string
		Build struct {
			Context    string
			Dockerfile string
		}
	}
	Volumes map[string]struct {
		Name string
	}
}

func actionExecCompose(files []string, arg ...string) func(f func(output []byte) carapace.Action) carapace.Action {
	return func(f func(output []byte) carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			a := []string{"compose"}
			for _, file := range files {
				a = append(a, "--file", file)
			}
			a = append(a, arg...)
			return carapace.ActionExecCommand("docker", a...)(f)
		})
	}
}

func actionConfig(files []string, f func(c config) carapace.Action) carapace.Action {
	return actionExecCompose(files, "config", "--format", "json")(func(output []byte) carapace.Action {
		var c config
		if err := json.Unmarshal(output, &c); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(c)
	})
}

// ActionServices completes services
//
//	elvish (ghcr.io/carapace-sh/carapace:v0.12.4)
//	fish (ghcr.io/carapace-sh/carapace:v0.12.4)
func ActionServices(files ...string) carapace.Action {
	return actionConfig(files, func(c config) carapace.Action {
		vals := make([]string, 0)
		for name, service := range c.Services {
			description := service.Image
			if strings.TrimSpace(description) == "" {
				description = fmt.Sprintf("%v/%v", service.Build.Context, service.Build.Dockerfile)
			}
			vals = append(vals, name, description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

type container struct {
	Name  string
	State string
}

type ContainerOpts struct {
	Files      []string
	Paused     bool
	Restarting bool
	Removing   bool
	Running    bool
	Dead       bool
	Created    bool
	Exited     bool
}

func (o *ContainerOpts) Default() {
	o.Paused = true
	o.Restarting = true
	o.Removing = true
	o.Running = true
	o.Dead = true
	o.Created = true
	o.Exited = true
}

func (o ContainerOpts) includeState(s string) bool {
	switch s {
	case "paused":
		return o.Paused
	case "restarting":
		return o.Restarting
	case "removing":
		return o.Removing
	case "running":
		return o.Running
	case "dead":
		return o.Dead
	case "created":
		return o.Created
	case "exited":
		return o.Exited
	default:
		return false
	}
}

// ActionContainers completes containers
//
//	carapace-bin-bash-1 (exited)
//	carapace-bin-elvish-1 (running)
func ActionContainers(opts ContainerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return actionExecCompose(opts.Files, "ps", "--format", "json", "--all")(func(output []byte) carapace.Action {
			var containers []container
			if err := json.Unmarshal(output, &containers); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, container := range containers {
				if opts.includeState(container.State) {
					vals = append(vals, container.Name, container.State, style.ForKeyword(container.State, c))
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
