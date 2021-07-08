package docker

import (
	"path/filepath"

	"strings"

	"github.com/rsteube/carapace"
)

func ActionConfigs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "config", "ls", "--format", "{{.Name}}\nupdated {{.UpdatedAt}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

func ActionContainers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "container", "ls", "--format", "{{.Names}}\n{{.Image}} ({{.Status}})")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

func ActionRepositories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "images", "--format", "{{.Repository}})")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValues(vals[:len(vals)-1]...)
		})
	})
}

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
		})
	})
}

func ActionContainerPath() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionExecCommand("docker", "container", "ls", "--format", "{{.Names}}\n{{.Image}} ({{.Status}})")(func(output []byte) carapace.Action {
				vals := strings.Split(string(output), "\n")
				return carapace.ActionValuesDescribed(vals[:len(vals)-1]...).Invoke(c).Suffix(":").ToA()
			})
		default:
			container := c.Parts[0]
			path := filepath.Dir(c.CallbackValue)
			return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
				return carapace.ActionExecCommand("docker", "exec", container, "ls", "-1", "-p", path)(func(output []byte) carapace.Action {
					lines := strings.Split(string(output), "\n")
					return carapace.ActionValues(lines[:len(lines)-1]...)
				})
			})
		}
	})
}

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

func ActionDetachKeys() carapace.Action {
	return carapace.ActionValues("{a-z}", `ctrl-\`, "ctrl-@", "ctrl-[", "ctrl-]", "ctrl-^", "ctrl-_", "ctrl-{a-z}")
}

func ActionContexts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "context", "ls", "--format", "{{.Name}}\n{{.Description}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

func ActionNetworks() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "network", "ls", "--format", "{{.Name}}\n{{.Driver}}/{{.Scope}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

func ActionNodes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "node", "ls", "--format", "{{.ID}}\n{{.Hostname}} {{.ManagerStatus}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

func ActionPlugins() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "plugin", "ls", "--format", "{{.Name}}\n{{.Description}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

func ActionSecrets() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "secret", "ls", "--format", "{{.Name}}\nupdated {{.UpdatedAt}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

func ActionServices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "service", "ls", "--format", "{{.Name}}\n{{.Image}} {{.Mode}} {{.Replicas}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

func ActionVolumes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "volume", "ls", "--format", "{{.Name}}\n{{.Driver}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}

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
