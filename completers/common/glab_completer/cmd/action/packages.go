package action

import (
	"fmt"
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

type pkg struct {
	Id          int
	Name        string
	Version     string
	PackageType string `json:"package_type"`
}

func actionPackages(cmd *cobra.Command, transform func(pkg) (string, string)) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []pkg
		return actionApi(cmd, "projects/:fullpath/packages?per_page=100", &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, package_ := range queryResult {
				value, description := transform(package_)
				vals = append(vals, value, description)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionPackages(cmd *cobra.Command) carapace.Action {
	return actionPackages(cmd, func(package_ pkg) (string, string) {
		return strconv.Itoa(package_.Id), fmt.Sprintf("%v %v (%v)", package_.Name, package_.Version, package_.PackageType)
	})
}

func ActionPackageNames(cmd *cobra.Command) carapace.Action {
	return actionPackages(cmd, func(package_ pkg) (string, string) {
		return package_.Name, fmt.Sprintf("%v (%v)", package_.Version, package_.PackageType)
	})
}
