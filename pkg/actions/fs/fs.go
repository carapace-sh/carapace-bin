// package fs contains filesystem related actions
package fs

import (
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/rsteube/carapace"
)

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

			if len(c.Parts) > 0 {
				operators := carapace.ActionValuesDescribed(
					"+", "adds the specified modes to the specified classes",
					"-", "removes the specified modes from the specified classes",
					"=", "the modes specified are to be made the exact modes for the specified classes",
				).Invoke(c)
				return classes.Merge(operators).Filter(c.Parts).ToA()
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
