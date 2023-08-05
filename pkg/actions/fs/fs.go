// package fs contains filesystem related actions
package fs

import (
	"os"
	"strings"
	"unicode"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionSubDirectories completes subdirectories of a given path
//
//	subdir/subsubdir
//	subdir/subsubder2
func ActionSubDirectories(path string) carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		if files, err := os.ReadDir(path + "/" + strings.Join(c.Parts, "/") + "/"); err != nil {
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
//
//	a+rw
//	g=rx
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
				return classes.Merge(operators).Filter(c.Parts...).ToA()
			}
			return classes.Filter(c.Parts...).ToA()
		} else {
			return carapace.ActionStyledValuesDescribed(
				"r", "read", style.Green,
				"w", "write", style.Red,
				"x", "execute", style.Yellow,
				"X", "special execute", style.Yellow,
				"s", "setuid/gid", style.Magenta,
				"t", "sticky", style.Magenta,
			).Invoke(c).Filter(c.Parts...).ToA()
		}
	}).Tag("symbolic filemodes")
}

// ActionFileModesNumeric completes numeric file modes
//
//	644
//	755
func ActionFileModesNumeric() carapace.Action {
	return carapace.ActionMultiParts("", func(c carapace.Context) carapace.Action {
		if len(c.Parts) < 3 {
			return carapace.ActionStyledValuesDescribed(
				"7", "read, write and execute", style.Red,
				"6", "read and write", style.Red,
				"5", "read and execute", style.Yellow,
				"4", "read only", style.Green,
				"3", "write and execute", style.Red,
				"2", "write only", style.Red,
				"1", "execute only", style.Yellow,
				"0", "none", style.Default,
			)
		} else {
			return carapace.ActionValues()
		}
	}).Tag("numeric filemodes")
}

// ActionFileModes completes numeric or symbolic file modes
//
//	644
//	a+rw
func ActionFileModes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		a := carapace.ActionValues().Invoke(c)
		if len(c.Value) == 0 || !unicode.IsDigit([]rune(c.Value)[0]) {
			symbolic := carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return ActionFileModesSymbolic().NoSpace()
			}).Invoke(c)
			a = a.Merge(symbolic)
		}
		if len(c.Value) == 0 || unicode.IsDigit([]rune(c.Value)[0]) {
			numeric := ActionFileModesNumeric().Invoke(c)
			a = a.Merge(numeric)
		}
		return a.ToA()
	})
}
