package fs

import (
	"os"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type mountEntry struct {
	source string
	target string
}

// ActionMounts completes file system mounts
//
//	/boot/efi (/dev/sda1)
//	/dev (dev)
func ActionMounts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch runtime.GOOS {
		case "darwin", "freebsd", "netbsd", "openbsd":
			return actionMountsFromMountCommand()
		default:
			return actionMountsFromProc()
		}
	}).Tag("mounts")
}

func actionMountsFromMountCommand() carapace.Action {
	return carapace.ActionExecCommand("mount")(func(output []byte) carapace.Action {
		return actionMountEntries(parseMountOutput(string(output)))
	})
}

func actionMountsFromProc() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		content, err := os.ReadFile("/proc/mounts")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return actionMountEntries(parseProcMounts(string(content)))
	})
}

func parseProcMounts(content string) []mountEntry {
	lines := strings.Split(content, "\n")
	entries := make([]mountEntry, 0)
	for _, line := range lines {
		if fields := strings.Fields(line); len(fields) > 1 {
			entries = append(entries, mountEntry{
				source: fields[0],
				target: strings.Replace(fields[1], `\040`, " ", -1),
			})
		}
	}
	return entries
}

func parseMountOutput(output string) []mountEntry {
	lines := strings.Split(output, "\n")
	entries := make([]mountEntry, 0)
	for _, line := range lines {
		source, target, ok := strings.Cut(line, " on ")
		if !ok {
			continue
		}
		target, _, ok = strings.Cut(target, " (")
		if !ok {
			continue
		}
		entries = append(entries, mountEntry{source: source, target: target})
	}
	return entries
}

func actionMountEntries(entries []mountEntry) carapace.Action {
	vals := make([]string, 0)
	for _, entry := range entries {
		vals = append(vals, entry.target, entry.source)
		vals = append(vals, entry.source, entry.target)
	}
	return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPath)
}
