package action

import (
	"github.com/rsteube/carapace"
	exec "golang.org/x/sys/execabs"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func currentRelease() (string, error) {
	if output, err := exec.Command("uname", "-r").Output(); err != nil {
		return "", err
	} else {
		return strings.Split(string(output), "\n")[0], nil
	}
}

func ActionLoadedKernelModules() carapace.Action {
	return carapace.ActionExecCommand("lsmod")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[1 : len(lines)-1] {
			vals = append(vals, strings.Fields(line)[0])
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionKernelModules(release string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.TrimSpace(release) == "" {
			var err error
			if release, err = currentRelease(); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		}

		vals := make([]string, 0)
		err := filepath.WalkDir("/lib/modules/"+release+"/kernel", func(path string, d fs.DirEntry, err error) error {
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
	})
}

func ActionKernelReleases() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		files, err := ioutil.ReadDir("/lib/modules")
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
	})
}
