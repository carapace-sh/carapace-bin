package vagrant

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionBoxes completes boxes
//
//	ubuntu/jammy64
//	debian/bookworm64
func ActionBoxes() carapace.Action {
	return carapace.ActionExecCommand("vagrant", "box", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		r := regexp.MustCompile(`^(?P<name>[^ ]+) \((?P<provider>[^,]+), (?P<version>.*)\)$`)
		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
}

// BoxProvidersOpts provides target filters for box provider completion
type BoxProvidersOpts struct {
	Box     string // box name
	Version string // box version
}

// ActionBoxProviders completes box providers
//
//	virtualbox
//	vmware
func ActionBoxProviders(opts BoxProvidersOpts) carapace.Action {
	return carapace.ActionExecCommand("vagrant", "box", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		r := regexp.MustCompile(`^(?P<name>[^ ]+) \((?P<provider>[^,]+), (?P<version>.*)\)$`)
		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				if opts.Box == matches[1] &&
					(opts.Version == "" || opts.Version == matches[3]) {
					vals = append(vals, matches[2])
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}

// BoxVersionsOpts provides target filters for box version completion
type BoxVersionsOpts struct {
	Box      string // box name
	Provider string // box provider
}

// ActionBoxVersions completes box versions
//
//	2024.01.01
//	2024.02.01
func ActionBoxVersions(opts BoxVersionsOpts) carapace.Action {
	return carapace.ActionExecCommand("vagrant", "box", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		r := regexp.MustCompile(`^(?P<name>[^ ]+) \((?P<provider>[^,]+), (?P<version>.*)\)$`)
		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				if opts.Box == matches[1] &&
					(opts.Provider == "" || opts.Provider == matches[2]) {
					vals = append(vals, matches[3])
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}
