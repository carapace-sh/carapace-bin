// package eopkg contains Solus eopkg package manager related actions
package eopkg

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPackages completes installed packages
//
//	zlib (3.19-1)
//	glibc (2.31-1)
func ActionPackages() carapace.Action {
	return carapace.ActionExecCommand("eopkg", "--no-color", "list-installed")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if name, desc, found := strings.Cut(line, " - "); found {
				vals = append(vals, strings.TrimSpace(name), strings.TrimSpace(desc))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("installed packages")
}

// ActionPackageSearch completes available packages
//
//	zlib (3.19-1)
//	glibc (2.31-1)
func ActionPackageSearch() carapace.Action {
	return carapace.ActionExecCommand("eopkg", "--no-color", "list-available")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			// skip "Repository : <name>" headers and the color legend line
			if name, desc, found := strings.Cut(line, " - "); found {
				vals = append(vals, strings.TrimSpace(name), strings.TrimSpace(desc))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("available packages")
}

// ActionRepositories completes configured repositories
//
//	Solus ([active])
//	Solus-Unstable ([inactive])
func ActionRepositories() carapace.Action {
	return carapace.ActionExecCommand("eopkg", "--no-color", "list-repo")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 {
				vals = append(vals, fields[0], strings.Join(fields[1:], " "))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("repositories")
}

// ActionComponents completes components
//
//	desktop.kde (KDE desktop environment)
//	server.python (Python server stack)
func ActionComponents() carapace.Action {
	return carapace.ActionExecCommand("eopkg", "--no-color", "list-components")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if name, desc, found := strings.Cut(line, " - "); found {
				vals = append(vals, strings.TrimSpace(name), strings.TrimSpace(desc))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("components")
}
