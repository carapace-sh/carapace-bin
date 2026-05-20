package action

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionDomains completes launchd target domains.
func ActionDomains() carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"system", "system-wide domain",
				"user", "per-user domain",
				"gui", "GUI user domain",
				"session", "audit session domain",
				"pid", "process domain",
			).Suffix("/").StyleF(style.ForKeyword)
		case 1:
			switch c.Parts[0] {
			case "user", "gui":
				return ActionUserIds()
			case "pid":
				return ps.ActionProcessIds()
			default:
				return carapace.ActionValues()
			}
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionServiceTargets completes launchd target specifiers, including services
// under the current domain when launchctl can list them.
func ActionServiceTargets() carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"system", "system-wide domain",
				"user", "per-user domain",
				"gui", "GUI user domain",
				"session", "audit session domain",
				"pid", "process domain",
			).Suffix("/").StyleF(style.ForKeyword)
		case 1:
			switch c.Parts[0] {
			case "system":
				return ActionServices("system")
			case "user", "gui":
				return carapace.Batch(
					ActionUserIds().Invoke(c).Suffix("/").ToA(),
					ActionServices(c.Parts[0]),
				).ToA()
			case "pid":
				return ps.ActionProcessIds().Invoke(c).Suffix("/").ToA()
			default:
				return carapace.ActionValues()
			}
		case 2:
			switch c.Parts[0] {
			case "user", "gui", "pid":
				return ActionServices(strings.Join(c.Parts, "/"))
			default:
				return carapace.ActionValues()
			}
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionServices completes services in the current GUI domain.
func ActionCurrentServices() carapace.Action {
	uid := strconv.Itoa(os.Getuid())
	return carapace.Batch(
		ActionServices("gui/"+uid),
		ActionServices("user/"+uid),
		ActionServices("system"),
	).ToA()
}

// ActionServices completes launchd service labels for a domain.
func ActionServices(domain string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		output, err := exec.Command("launchctl", "print", domain).Output()
		if err != nil {
			return carapace.ActionValues()
		}

		vals := make([]string, 0)
		inServices := false
		for _, line := range strings.Split(string(output), "\n") {
			trimmed := strings.TrimSpace(line)
			switch {
			case trimmed == "services = {":
				inServices = true
				continue
			case inServices && trimmed == "}":
				inServices = false
			case inServices:
				fields := strings.Fields(trimmed)
				if len(fields) == 0 {
					continue
				}
				service := fields[len(fields)-1]
				if strings.Contains(service, ".") {
					vals = append(vals, service)
				}
			}
		}

		return carapace.ActionValues(vals...)
	}).Tag("launchd services").Uid("launchctl", "services", "domain", domain)
}

// ActionUserIds completes local user ids.
func ActionUserIds() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		content, err := os.ReadFile("/etc/passwd")
		if err != nil {
			return carapace.ActionValues(strconv.Itoa(os.Getuid()))
		}

		vals := make([]string, 0)
		for _, entry := range strings.Split(string(content), "\n") {
			fields := strings.Split(entry, ":")
			if len(fields) > 4 {
				vals = append(vals, fields[2], fields[0])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("user ids").Uid("os", "uid")
}
