package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apk"
	"github.com/spf13/cobra"
)

func ActionPackageSearch(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return apk.ActionPackageSearch(apk.PackageSearchOpts{
			Arch:             cmd.Flag("arch").Value.String(),
			KeysDir:          cmd.Flag("keys-dir").Value.String(),
			RepositoriesFile: cmd.Flag("repositories-file").Value.String(),
			Repository:       cmd.Flag("repository").Value.String(),
		})
	})
}

func ActionPackages(cmd *cobra.Command, installedOnly bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return apk.ActionPackages(apk.PackageOpts{
			InstalledOnly:    installedOnly,
			Arch:             cmd.Flag("arch").Value.String(),
			KeysDir:          cmd.Flag("keys-dir").Value.String(),
			RepositoriesFile: cmd.Flag("repositories-file").Value.String(),
			Repository:       cmd.Flag("repository").Value.String(),
		})
	})
}
