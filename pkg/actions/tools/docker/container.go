package docker

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionContainers completes container names
//
//	agitated_engelbart (alpine (Up 6 seconds))
//	crazy_satoshi (alpine (Up 4 seconds))
func ActionContainers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "container", "ls", "--all", "--format", "{{.Names}}\t{{.Image}}\t{{.Status}}", "--filter", "name="+c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				if splitted := strings.Split(line, "\t"); len(splitted) == 3 {
					var s string
					switch status, _, _ := strings.Cut(splitted[2], " "); status {
					case "Up":
						s = style.Carapace.KeywordPositive
					case "Exited":
						s = style.Carapace.KeywordNegative
					default:
						s = style.Carapace.KeywordAmbiguous
					}
					vals = append(vals, splitted[0], fmt.Sprintf("%v (%v)", splitted[1], splitted[2]), s)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Tag("containers").UidF(Uid("container"))
}

// ActionContainerIds completes container names
//
//	c84ca01b41f1 (alpine (Up 6 seconds))
//	1c3cf2aeee96 (alpine (Up 4 seconds))
func ActionContainerIds() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "container", "ls", "--format", "{{.ID}}\t{{.Image}}\t{{.Status}}")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				if splitted := strings.Split(line, "\t"); len(splitted) == 3 {
					var s string
					switch status, _, _ := strings.Cut(splitted[2], " "); status {
					case "Up":
						s = style.Carapace.KeywordPositive
					case "Exited":
						s = style.Carapace.KeywordNegative
					default:
						s = style.Carapace.KeywordAmbiguous
					}
					vals = append(vals, splitted[0], fmt.Sprintf("%v (%v)", splitted[1], splitted[2]), s)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Tag("container ids").UidF(Uid("container"))
}

// ActionContainerPath completes container names and their file system separately
//
//	agitated_engelbart:/bin/echo
//	crazy_satoshi:/usr/lib/engines-1.1/afalg.so
func ActionContainerPath() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionContainers().Suffix(":")
		default:
			container := c.Parts[0]
			path := filepath.Dir(c.Value)

			args := []string{"exec", container, "ls", "-1", "-p", path}
			if splitted := strings.Split(c.Value, "/"); strings.HasPrefix(splitted[len(splitted)-1], ".") {
				args = append(args, "-a") // show hidden
			}
			return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
				return carapace.ActionExecCommand("docker", args...)(func(output []byte) carapace.Action {
					lines := strings.Split(string(output), "\n")

					vals := make([]string, 0)
					for _, path := range lines[:len(lines)-1] {
						vals = append(vals, path, style.ForPathExt(path, c))
					}
					return carapace.ActionStyledValues(vals...)
				})
			})
		}
	})
}
