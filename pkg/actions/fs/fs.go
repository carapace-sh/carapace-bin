// package fs contains filesystem related actions
package fs

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
	"unicode"

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
		if !strings.ContainsAny(strings.Join(c.Parts, ""), "+-=") {
			classes := carapace.ActionValuesDescribed(
				"u", "user",
				"g", "group",
				"o", "other",
				"a", "all",
			).Invoke(c)

			operators := carapace.ActionValuesDescribed(
				"+", "adds the specified modes to the specified classes",
				"-", "removes the specified modes from the specified classes",
				"=", "the modes specified are to be made the exact modes for the specified classes",
			).Invoke(c)
			return classes.Merge(operators).Filter(c.Parts).ToA()
		} else {
			return carapace.ActionValuesDescribed(
				"r", "read",
				"w", "write",
				"x", "execute",
				"X", "special execute",
				"s", "setuid/gid",
				"t", "sticky",
			).Invoke(c).Filter(c.Parts).ToA()
		}
	})
}

// ActionFileModesNumeric completes numeric file modes
//   644
//   755
func ActionFileModesNumeric() carapace.Action {
	return carapace.ActionMultiParts("", func(c carapace.Context) carapace.Action {
		if len(c.Parts) < 3 {
			return carapace.ActionValuesDescribed(
				"7", "read, write and execute",
				"6", "read and write",
				"5", "read and execute",
				"4", "read only",
				"3", "write and execute",
				"2", "write only",
				"1", "execute only",
				"0", "none",
			)
		} else {
			return carapace.ActionValues()
		}
	})
}

// ActionFileModes completes numeric or symbolic file modes
//   644
//   a+rw
func ActionFileModes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		a := carapace.ActionValues().Invoke(c)
		if len(c.CallbackValue) == 0 || !unicode.IsDigit([]rune(c.CallbackValue)[0]) {
			symbolic := carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return ActionFileModesSymbolic()
			}).Invoke(c)
			a = a.Merge(symbolic)
		}
		if len(c.CallbackValue) == 0 || unicode.IsDigit([]rune(c.CallbackValue)[0]) {
			numeric := ActionFileModesNumeric().Invoke(c)
			a = a.Merge(numeric)
		}
		return a.ToA()
	})
}

// ActionBlockDevices completes block devices
//   /dev/sda (10G)
//   /dev/sda1 (2G Linux swap)
func ActionBlockDevices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if output, err := exec.Command("lsblk", "-o", "PATH,SIZE,PARTTYPENAME").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[1 : len(lines)-1] {
				splitted := strings.SplitN(line, " ", 3)
				vals = append(vals, splitted[0], strings.TrimSpace(fmt.Sprintf("%v %v", strings.TrimSpace(splitted[1]), strings.TrimSpace(splitted[2]))))
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
