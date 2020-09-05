package action

import (
	"os/exec"
	"path/filepath"

	"strings"

	"github.com/rsteube/carapace"
)

func ActionConfigs() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("docker", "config", "ls", "--format", "{{.Name}}\nupdated {{.UpdatedAt}}").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}
	})
}

func ActionContainers() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("docker", "container", "ls", "--format", "{{.Names}}\n{{.Image}} ({{.Status}})").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}
	})
}

func ActionRepositoryTags() carapace.Action {
	return carapace.ActionMultiParts(":", func(args []string, parts []string) []string {
		if output, err := exec.Command("docker", "image", "ls", "--format", "{{.Repository}}:{{.Tag}}", "--filter", "dangling=false").Output(); err != nil {
			return []string{}
		} else {
			// TODO add checks to not cause an out of bounds error
			vals := strings.Split(string(output), "\n")
			switch len(parts) {
			case 0:
				reposWithSuffix := make([]string, len(vals))
				for index, val := range vals[:len(vals)-1] {
					if val != " " {
						reposWithSuffix[index] = strings.SplitAfter(val, ":")[0]
					}
				}
				return reposWithSuffix
			case 1:
				tags := make([]string, 0)
				for _, val := range vals[:len(vals)-1] {
					if strings.HasPrefix(val, parts[0]) {
						tag := strings.Split(val, ":")[1]
						tags = append(tags, tag)
					}
				}
				return tags
			default:
				return []string{}
			}
		}
	})
}

// TODO not yet working - also needs multiple characters to split on `:` `/`
func ActionContainerPath() carapace.Action {
	return carapace.ActionMultiParts(":", func(args []string, parts []string) []string {
		switch len(parts) {
		case 0:
			// TODO add description support
			//if output, err := exec.Command("docker", "container", "ls", "--format", "{{.Names}}:\n{{.Image}} ({{.Status}})").Output(); err != nil {
			if output, err := exec.Command("docker", "container", "ls", "--format", "{{.Names}}:").Output(); err != nil {
				return []string{}
			} else {
				vals := strings.Split(string(output), "\n")
				return vals[:len(vals)-1]
			}
		default:
			if output, err := exec.Command("docker", "exec", parts[0], "ls", filepath.Dir(strings.Join(parts[1:], "/"))).Output(); err != nil {
				return []string{}
			} else {
				return strings.Split(string(output), "\n")
			}
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
	// TODO needs better escaping in carapace to work
	// "detach-keys": carapace.ActionValues("{a-z}", `ctrl-\`, "ctrl-@", "ctrl-[", "ctrl-]",  "ctrl-^", "ctrl-_", "ctrl-{a-z}"),
	return carapace.ActionValues()
}

func ActionContexts() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("docker", "context", "ls", "--format", "{{.Name}}\n{{.Description}}").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}
	})
}

func ActionNetworks() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("docker", "network", "ls", "--format", "{{.Name}}\n{{.Driver}}/{{.Scope}}").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}
	})
}

func ActionNodes() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("docker", "node", "ls", "--format", "{{.ID}}\n{{.Hostname}} {{.ManagerStatus}}").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}
	})
}

func ActionPlugins() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("docker", "plugin", "ls", "--format", "{{.Name}}\n{{.Description}}").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}
	})
}

func ActionSecrets() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("docker", "secret", "ls", "--format", "{{.Name}}\nupdated {{.UpdatedAt}}").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}
	})
}

func ActionSerices() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("docker", "service", "ls", "--format", "{{.Name}}\n{{.Image}} {{.Mode}} {{.Replicas}}").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}
	})
}
