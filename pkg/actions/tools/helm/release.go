package helm

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type release struct {
	Name       string
	Namespace  string
	Revision   string
	Updated    string
	Status     string
	Chart      string
	AppVersion string `json:"app_version"`
}

type ReleasesOpts struct {
	KubeContext string
	Namespace   string
}

// ActionReleases completes releases
func ActionReleases(opts ReleasesOpts) carapace.Action {
	return carapace.ActionCallback(func(_ carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("helm", "list", "--kube-context", opts.KubeContext, "--namespace", opts.Namespace, "--output", "json")(func(output []byte) carapace.Action {
			var releases []release
			if err := json.Unmarshal(output, &releases); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0, len(releases)*2)
			for _, release := range releases {
				vals = append(vals, release.Name, release.Chart)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
