package fs

import (
	"os"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionMounts completes file system mounts
//
//	/boot/efi (/dev/sda1)
//	/dev (dev)
func ActionMounts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch {
		case runtime.GOOS == "darwin":
			return actionMountsDarwin()
		case runtime.GOOS == "windows":
			return actionMountsWindows()
		default:
			return actionMountsProc()
		}
	}).Tag("mounts")
}

func actionMountsProc() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		content, err := os.ReadFile("/proc/mounts")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		lines := strings.Split(string(content), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 {
				vals = append(vals, strings.ReplaceAll(fields[1], `\040`, " "), fields[0])
			}
		}
		return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPath)
	})
}

func actionMountsDarwin() carapace.Action {
	return carapace.ActionExecCommand("mount")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if idx := strings.Index(line, " on "); idx != -1 {
				rest := line[idx+4:]
				if pIdx := strings.Index(rest, " ("); pIdx != -1 {
					mountPoint := rest[:pIdx]
					device := line[:idx]
					vals = append(vals, mountPoint, device)
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPath)
	})
}

func actionMountsWindows() carapace.Action {
	return carapace.ActionExecCommand("wmic", "logicaldisk", "get", "Caption,ProviderName", "/format:csv")(func(output []byte) carapace.Action {
		lines := strings.Split(strings.ReplaceAll(string(output), "\r", ""), "\n")
		if len(lines) < 2 {
			return carapace.ActionValues()
		}

		header := strings.Split(lines[0], ",")
		indexes := make(map[string]int)
		for i, col := range header {
			indexes[strings.TrimSpace(col)] = i
		}

		vals := make([]string, 0)
		for _, line := range lines[1:] {
			fields := strings.Split(line, ",")
			if len(fields) < len(header) {
				continue
			}

			caption := strings.TrimSpace(fields[indexes["Caption"]])
			if caption == "" {
				continue
			}
			provider := strings.TrimSpace(fields[indexes["ProviderName"]])
			vals = append(vals, caption, provider)
		}
		return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPath)
	})
}
