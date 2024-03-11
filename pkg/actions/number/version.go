package number

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"golang.org/x/mod/semver"
)

// ActionSemanticVersions completes the next semantic version based on the given existing versions.
//
//	v0.20.4 (next patch version)
//	v0.21.0 (next minor version)
func ActionSemanticVersions(versions ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(versions) == 0 {
			versions = append(versions, "v0.0.0")
		}

		withoutPrefix := make(map[string]bool, 0)
		for index, version := range versions {
			switch {
			case semver.IsValid(version):
				// valid
			case semver.IsValid("v" + version):
				// support versions without 'v' prefix as well
				withoutPrefix[version] = true
				versions[index] = "v" + version
			default:
				// skip invalid
				versions[index] = ""
			}
		}

		sort.Sort(semver.ByVersion(versions))
		canonical := semver.Canonical(versions[len(versions)-1])
		if canonical == "" {
			return carapace.ActionMessage("invalid version %#v", versions[0])
		}

		prefix := "v"
		if _, ok := withoutPrefix[strings.TrimPrefix(canonical, "v")]; ok {
			prefix = ""
		}

		r := regexp.MustCompile(`^v(?P<major>[0-9]+)\.(?P<minor>[0-9]+)\.(?P<patch>[0-9]+)`)
		vals := make([]string, 0)
		if matches := r.FindStringSubmatch(canonical); matches != nil {
			// valid format ensured by semver.Canonical
			major, _ := strconv.Atoi(matches[1])
			minor, _ := strconv.Atoi(matches[2])
			patch, _ := strconv.Atoi(matches[3])

			if semver.Prerelease(canonical) != "" {
				vals = append(vals,
					fmt.Sprintf("%v%v.%v.%v", prefix, major, minor, patch), "next version",
				)
			} else {
				vals = append(vals,
					fmt.Sprintf("%v%v.%v.%v", prefix, major+1, 0, 0), "next major version",
					fmt.Sprintf("%v%v.%v.%v", prefix, major, minor+1, 0), "next minor version",
					fmt.Sprintf("%v%v.%v.%v", prefix, major, minor, patch+1), "next patch version",
				)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
