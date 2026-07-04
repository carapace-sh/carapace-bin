package os

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
)

func currentRelease(c carapace.Context) (string, error) {
	if output, err := c.Command("uname", "-r").Output(); err != nil {
		return "", err
	} else {
		return strings.Split(string(output), "\n")[0], nil
	}
}

// ActionKernelModulesLoaded completes currently loaded kernel modules
//
//	ac97_bus
//	crc32c_intel
func ActionKernelModulesLoaded() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch runtime.GOOS {
		case "darwin":
			return carapace.ActionExecCommand("kextstat", "-l")(func(output []byte) carapace.Action {
				vals := make([]string, 0)
				for line := range strings.SplitSeq(string(output), "\n") {
					fields := strings.Fields(line)
					if len(fields) >= 6 && strings.HasPrefix(fields[5], "(") {
						name := strings.Trim(fields[5], "()")
						if name != "" {
							vals = append(vals, name)
						}
					}
				}
				return carapace.ActionValues(vals...)
			}).Tag("kernel modules")
		default:
			return carapace.ActionExecCommand("lsmod")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				vals := make([]string, 0)
				for _, line := range lines[1 : len(lines)-1] {
					vals = append(vals, strings.Fields(line)[0])
				}
				return carapace.ActionValues(vals...)
			}).Tag("kernel modules")
		}
	})
}

type KernelModulesOpts struct {
	Basedir string
	Release string
}

// ActionKernelModules completes kernel modules
//
//	ac97_bus
//	crc32c_intel
func ActionKernelModules(opts KernelModulesOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch runtime.GOOS {
		case "darwin":
			return actionKernelModulesDarwin()
		default:
			if strings.TrimSpace(opts.Release) == "" {
				var err error
				if opts.Release, err = currentRelease(c); err != nil {
					return carapace.ActionMessage(err.Error())
				}
			}

			vals := make([]string, 0)
			err := filepath.WalkDir(fmt.Sprintf("%v/lib/modules/%v/kernel", opts.Basedir, opts.Release), func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if !d.IsDir() {
					name := filepath.Base(path)
					name = strings.TrimSuffix(name, ".xz")
					name = strings.TrimSuffix(name, ".ko")
					vals = append(vals, name)
				}
				return nil
			})
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionValues(vals...)
		}
	}).Tag("kernel modules")
}

func actionKernelModulesDarwin() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		dirs := []string{"/Library/Extensions", "/System/Library/Extensions"}
		for _, dir := range dirs {
			entries, err := os.ReadDir(dir)
			if err != nil {
				continue
			}
			for _, entry := range entries {
				if entry.IsDir() && strings.HasSuffix(entry.Name(), ".kext") {
					vals = append(vals, strings.TrimSuffix(entry.Name(), ".kext"))
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionKernelReleases completes kernel releases
//
//	5.10.79-1-MANJARO
//	5.4.159-1-MANJARO
func ActionKernelReleases(basedir string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if runtime.GOOS == "darwin" {
			return carapace.ActionValues()
		}

		files, err := os.ReadDir(basedir + "/lib/modules")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, file := range files {
			if file.IsDir() {
				vals = append(vals, file.Name())
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("kernel releases")
}
