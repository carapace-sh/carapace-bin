package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

func ActionPackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return npm.ActionPackageSearch(cmd.Flag("registry").Value.String())
	})
}

func ActionPackageNames(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return npm.ActionPackageNames(cmd.Flag("registry").Value.String())
	})
}

func ActionPackageVersions(cmd *cobra.Command, pkg string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return npm.ActionPackageVersions(npm.PackageOpts{Registry: cmd.Flag("registry").Value.String(), Package: pkg})
	})
}

func ActionPackageTags(cmd *cobra.Command, pkg string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return npm.ActionPackageTags(npm.PackageOpts{Registry: cmd.Flag("registry").Value.String(), Package: pkg})
	})
}

func ActionConfigKeys(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if flag := cmd.Flag("global"); flag != nil && flag.Changed {
			return npm.ActionGlobalConfigKeys()
		}
		return npm.ActionLocalConfigKeys()
	})
}
