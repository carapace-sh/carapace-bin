package fury

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionPackages completes packages
//
//	deb:carapace-bin
//	rpm:carapace-bin
func ActionPackages() carapace.Action {
	return carapace.ActionExecCommand("fury", "packages")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[4:] {
			if fields := strings.Fields(line); len(fields) == 4 {
				switch fields[3] {
				case "public":
					vals = append(vals, fmt.Sprintf("%v:%v", fields[1], fields[0]), style.Green)
				default:
					vals = append(vals, fmt.Sprintf("%v:%v", fields[1], fields[0]), style.Red)
				}
			}
		}
		return carapace.ActionStyledValues(vals...)
	}).Tag("packages")
}

// ActionPackageVersions completes package versions
//
//	0.13.4 (2022-07-29 14:11)
//	0.13.5 (2022-08-17 17:08)
func ActionPackageVersions(pkg string) carapace.Action {
	return carapace.ActionExecCommand("fury", "versions", pkg)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		unique := make(map[string]string)
		for _, line := range lines[4:] {
			if fields := strings.Fields(line); len(fields) == 5 {
				unique[fields[0]] = fields[2] + " " + fields[3]
			}
		}

		vals := make([]string, 0)
		for k, v := range unique {
			vals = append(vals, k, v)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("package versions")
}
