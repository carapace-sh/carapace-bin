// package fs contains filesystem related actions
package fs

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionMounts completes file system mounts
//   /boot/efi (/dev/sda1)
//   /dev (dev)
func ActionMounts() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("mount").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			re := regexp.MustCompile(`^(?P<target>\S+) on (?P<mountt>\S+) type (?P<type>\S+) (?P<mode>.+)`)
			mounts := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if re.MatchString(line) {
					matches := re.FindStringSubmatch(line)
					mounts = append(mounts, matches[2], matches[1])
				}
			}
			return carapace.ActionValuesDescribed(mounts...)
		}
	})
}

// ActionSubDirectories completes subdirectories of a given path
//   subdir/subsubdir
//   subdir/subsubder2
func ActionSubDirectories(path string) carapace.Action {
	return carapace.ActionMultiParts("/", func(args, parts []string) carapace.Action {
		if files, err := ioutil.ReadDir(path + "/" + strings.Join(parts, "/") + "/"); err != nil {
			return carapace.ActionValues()
		} else {
			dirs := make([]string, 0)
			for _, file := range files {
				if file.IsDir() && !strings.HasPrefix(file.Name(), ".") {
					dirs = append(dirs, file.Name()+"/")
				}
			}
			return carapace.ActionValues(dirs...)
		}
	})
}
