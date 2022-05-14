// package docker contains docker related actions
package docker

import (
	"path/filepath"

	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionConfigs completes config names
//   another (updated 4 seconds ago)
//   example (updated 23 seconds ago)
func ActionConfigs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "config", "ls", "--format", "{{.Name}}\nupdated {{.UpdatedAt}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

// ActionContainers completes container names
//   agitated_engelbart (alpine (Up 6 seconds))
//   crazy_satoshi (alpine (Up 4 seconds))
func ActionContainers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "container", "ls", "--format", "{{.Names}}\n{{.Image}} ({{.Status}})", "--filter", "name="+c.CallbackValue)(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...).Style(styles.Docker.Container)
		})
	})
}

// ActionContainerIds completes container names
//   c84ca01b41f1 (alpine (Up 6 seconds))
//   1c3cf2aeee96 (alpine (Up 4 seconds))
func ActionContainerIds() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "container", "ls", "--format", "{{.ID}}\n{{.Image}} ({{.Status}})")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

// ActionRepositories completes repository names
//   alpine
//   bash
func ActionRepositories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "images", "--format", "{{.Repository}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValues(vals[:len(vals)-1]...)
		})
	})
}

// ActionRepositoryTags completes repository names and tags separately
//   alpine:latest
//   bash:latest
func ActionRepositoryTags() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "image", "ls", "--format", "{{.Repository}}:{{.Tag}}", "--filter", "dangling=false")(func(output []byte) carapace.Action {
			// TODO add checks to not cause an out of bounds error
			lines := strings.Split(string(output), "\n")
			switch len(c.Parts) {
			case 0:
				reposWithSuffix := make([]string, len(lines)-1)
				for index, val := range lines[:len(lines)-1] {
					if val != " " {
						reposWithSuffix[index] = strings.SplitAfter(val, ":")[0]
					}
				}
				return carapace.ActionValues(reposWithSuffix...)
			case 1:
				tags := make([]string, 0)
				for _, val := range lines[:len(lines)-1] {
					if strings.HasPrefix(val, c.Parts[0]) {
						tag := strings.Split(val, ":")[1]
						tags = append(tags, tag)
					}
				}
				return carapace.ActionValues(tags...)
			default:
				return carapace.ActionValues()
			}
		}).Style(styles.Docker.Image)
	})
}

// ActionContainerPath completes container names and their file system separately
//   agitated_engelbart:/bin/echo
//   crazy_satoshi:/usr/lib/engines-1.1/afalg.so
func ActionContainerPath() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionExecCommand("docker", "container", "ls", "--format", "{{.Names}}\n{{.Image}} ({{.Status}})")(func(output []byte) carapace.Action {
				vals := strings.Split(string(output), "\n")
				return carapace.ActionValuesDescribed(vals[:len(vals)-1]...).Invoke(c).Suffix(":").ToA().Style(styles.Docker.Container)
			})
		default:
			container := c.Parts[0]
			path := filepath.Dir(c.CallbackValue)

			args := []string{"exec", container, "ls", "-1", "-p", path}
			if splitted := strings.Split(c.CallbackValue, "/"); strings.HasPrefix(splitted[len(splitted)-1], ".") {
				args = append(args, "-a") // show hidden
			}
			return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
				return carapace.ActionExecCommand("docker", args...)(func(output []byte) carapace.Action {
					lines := strings.Split(string(output), "\n")

					vals := make([]string, 0)
					for _, path := range lines[:len(lines)-1] {
						vals = append(vals, path, style.ForPathExt(path))
					}
					return carapace.ActionStyledValues(vals...)
				})
			})
		}
	})
}

// ActionLogDrivers completes log drivers
//   splunk (Writes log messages to splunk using the HTTP Event Collector.)
//   syslog (Writes logging messages to the syslog facility. The syslog daemon must be run...)
func ActionLogDrivers() carapace.Action {
	return carapace.ActionValuesDescribed(
		"none", "No logs are available for the container and docker logs does not return any output.",
		"local", "Logs are stored in a custom format designed for minimal overhead.",
		"json-file", "The logs are formatted as JSON. The default logging driver for Docker.",
		"syslog", "Writes logging messages to the syslog facility. The syslog daemon must be running on the host machine.",
		"journald", "Writes log messages to journald. The journald daemon must be running on the host machine.",
		"gelf", "Writes log messages to a Graylog Extended Log Format (GELF) endpoint such as Graylog or Logstash.",
		"fluentd", "Writes log messages to fluentd (forward input). The fluentd daemon must be running on the host machine.",
		"awslogs", "Writes log messages to Amazon CloudWatch Logs.",
		"splunk", "Writes log messages to splunk using the HTTP Event Collector.",
		"etwlogs", "Writes log messages as Event Tracing for Windows (ETW) events. Only available on Windows platforms.",
		"gcplogs", "Writes log messages to Google Cloud Platform (GCP) Logging.",
		"logentries", "Writes log messages to Rapid7 Logentries.",
	)
}

// ActionDetachKeys completes detach keys
//   ctrl-@
//   ctrl-[
func ActionDetachKeys() carapace.Action {
	return carapace.ActionValues("{a-z}", `ctrl-\`, "ctrl-@", "ctrl-[", "ctrl-]", "ctrl-^", "ctrl-_", "ctrl-{a-z}")
}

// ActionContexts completes context names
//   default (Current DOCKER_HOST based configuration)
//   example (custom context)
func ActionContexts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "context", "ls", "--format", "{{.Name}}\n{{.Description}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

// ActionNetworks completes network names
//   bridge (bridge/local)
//   host (host/local)
func ActionNetworks() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "network", "ls", "--format", "{{.Name}}\n{{.Driver}}/{{.Scope}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...).Style(styles.Docker.Network)
		})
	})
}

// ActionNodes completes node ids
//   r08tybjkcyar8vdglerxxxxxx
//   r08tybjkcyar8vdgleryyyyyy
func ActionNodes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "node", "ls", "--format", "{{.ID}}\n{{.Hostname}} {{.ManagerStatus}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...).Style(styles.Docker.Node)
		})
	})
}

// ActionPlugins completes plugins
//  TODO example
func ActionPlugins() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "plugin", "ls", "--format", "{{.Name}}\n{{.Description}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

// ActionSecrets completes secrets
//   another (updated 6 seconds ago)
//   example (updated 11 seconds ago)
func ActionSecrets() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "secret", "ls", "--format", "{{.Name}}\nupdated {{.UpdatedAt}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}).Style(styles.Docker.Secret)
	})
}

// ActionServices completes services
//   funny_robinson (alpine:latest replicated 0/1)
//   vigilant_mccarthy (bash:latest replicated 0/1)
func ActionServices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "service", "ls", "--format", "{{.Name}}\n{{.Image}} {{.Mode}} {{.Replicas}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}).Style(styles.Docker.Service)
	})
}

// ActionVolumes completes volume names
//   carapace-bin_go (local)
//   d35b4ebbab7bd0e9c155dbc9c75361150658ca525793af1d8fcbf97d058b905b (local)
func ActionVolumes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "volume", "ls", "--format", "{{.Name}}\n{{.Driver}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}).Style(styles.Docker.Volume)
	})
}

// ActionStacks completes stacks
//   // TODO example
func ActionStacks(orchestrator string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		arguments := []string{"stack", "ls", "--format", "{{.Name}}\n{{.Services}} on {{.Orchestrator}}"}
		if orchestrator == "swarm" || orchestrator == "kubernetes" {
			arguments = append(arguments, "--orchestrator", orchestrator)
		}
		return carapace.ActionExecCommand("docker", arguments...)(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}
