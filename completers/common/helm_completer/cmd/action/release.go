package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/helm"
	"github.com/spf13/cobra"
)

func ActionReleases(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return helm.ActionReleases(helm.ReleasesOpts{
			Namespace:   cmd.Root().Flag("namespace").Value.String(),
			KubeContext: cmd.Root().Flag("kube-context").Value.String(),
		})
	})
}
