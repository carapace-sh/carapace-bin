// package fs contains filesystem related actions
package fs

import (
	"archive/zip"
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
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
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
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		if files, err := ioutil.ReadDir(path + "/" + strings.Join(c.Parts, "/") + "/"); err != nil {
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

// ActionZipFileContents completes contents of given zip file
//   fileA
//   dirA/fileB
func ActionZipFileContents(file string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if reader, err := zip.OpenReader(file); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			defer reader.Close()
			vals := make([]string, len(reader.File))
			for index, f := range reader.File {
				vals[index] = f.Name
			}
			return carapace.ActionValues(vals...).Invoke(c).ToMultiPartsA("/")
		}
	})
}

// ActionFileModesSymbolic completes symbolic file modes
//   a+rw
//   g=rx
func ActionFileModesSymbolic() carapace.Action {
	return carapace.ActionMultiParts("", func(c carapace.Context) carapace.Action {
		if !strings.ContainsAny(c.CallbackValue, "+-=") {
			classes := carapace.ActionValuesDescribed(
				"u", "user",
				"g", "group",
				"o", "other",
				"a", "all",
			).Invoke(c)

			if c.CallbackValue != "" {
				operators := carapace.ActionValuesDescribed(
					"u", "user",
					"+", "adds the specified modes to the specified classes",
					"-", "removes the specified modes from the specified classes",
					"=", "the modes specified are to be made the exact modes for the specified classes",
				).Invoke(c)
				return classes.Merge(operators).Filter(c.Parts).Filter([]string{c.CallbackValue[len(c.CallbackValue)-1:]}).ToA()
			}
			return classes.Filter(c.Parts).ToA()
		} else {
			return carapace.ActionValuesDescribed(
				"r", "read",
				"w", "write",
				"x", "execute",
				"X", "special execute",
				"s", "setuid/gid",
				"t", "sticky",
			).Invoke(c).Filter(c.Parts).Filter([]string{c.CallbackValue[len(c.CallbackValue)-1:]}).ToA()
		}
	})
}
