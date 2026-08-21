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
	return carapace.ActionExecCommand("eopkg", "list-installed")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 {
				vals = append(vals, fields[0], strings.Join(fields[1:], " "))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionPackageSearch completes available packages
//
//	zlib (3.19-1)
//	glibc (2.31-1)
func ActionPackageSearch() carapace.Action {
	return carapace.ActionExecCommand("eopkg", "list-available")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 {
				vals = append(vals, fields[0], strings.Join(fields[1:], " "))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionRepositories completes configured repositories
//
//	Solus ([enabled])
//	Solus-Unstable ([disabled])
func ActionRepositories() carapace.Action {
	return carapace.ActionExecCommand("eopkg", "list-repo")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 {
				vals = append(vals, fields[0], strings.Join(fields[1:], " "))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
