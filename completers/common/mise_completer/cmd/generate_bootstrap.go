package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_bootstrapCmd = &cobra.Command{
	Use:    "bootstrap",
	Short:  "Deprecated; use `mise generate install-script`",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_bootstrapCmd).Standalone()

	generate_bootstrapCmd.Flags().BoolP("help", "h", false, "Print help")
	generate_bootstrapCmd.Flags().BoolP("localize", "l", false, "Keep mise data and cache in a project-local directory (`.mise` by default)")
	generate_bootstrapCmd.Flags().String("localized-dir", ".mise", "Directory for mise data when using --localize, relative to the project")
	generate_bootstrapCmd.Flags().StringP("version", "V", "", "mise version to pin; defaults to the release `mise self-update` would install")
	generate_bootstrapCmd.Flags().Bool("windows", false, "Also write a Windows launcher, `<PATH>.cmd`")
	generate_bootstrapCmd.Flags().StringP("write", "w", "", "Write the script to PATH (./bin/mise when no path is given) and make it executable, instead of printing it")
	generate_bootstrapCmd.Flag("write").NoOptDefVal = " "
	generateCmd.AddCommand(generate_bootstrapCmd)

	carapace.Gen(generate_bootstrapCmd).FlagCompletion(carapace.ActionMap{
		"localized-dir": carapace.ActionDirectories(),
	})
}
