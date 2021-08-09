package pub

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action/utils"
)

type searchResult struct {
	Packages []struct {
		Package string
	}
}

// ActionPackageSearch completes packages from pub.dev
//   animated_text_kit
//   dotted_border
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

// ActionPackageVersions completes versions of a package
//   1.3.1 (about 2 years ago)
//   4.1.1 (about 4 months ago)
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
