package action

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action/utils"
)

func ActionActivePackages() carapace.Action {
	return carapace.ActionExecCommand("pub", "global", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, " ", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

type searchResult struct {
	Packages []struct {
		Package string
	}
}

func ActionPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO verify on windows - should have a curl alias
		return carapace.ActionExecCommand("curl", fmt.Sprintf("https://pub.dev/api/search?q=package:%v", url.QueryEscape(c.CallbackValue)))(func(output []byte) carapace.Action {
			var result searchResult
			err := json.Unmarshal(output, &result)
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, pkg := range result.Packages {
				vals = append(vals, pkg.Package)
			}
			return carapace.ActionValues(vals...)
		})
	})
}

type dartPackage struct {
	Versions []struct {
		Version   string
		Published time.Time
	}
}

func ActionPackageVersions(pkg string) carapace.Action {
	// TODO verify on windows - should have a curl alias
	return carapace.ActionExecCommand("curl", fmt.Sprintf("https://pub.dev/api/packages/%v", url.QueryEscape(pkg)))(func(output []byte) carapace.Action {
		var result dartPackage
		err := json.Unmarshal(output, &result)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, v := range result.Versions {
			vals = append(vals, v.Version, utils.FuzzyAgo(time.Since(v.Published)))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
