package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_installScriptCmd = &cobra.Command{
	Use:   "install-script",
	Short: "Generate a script that downloads and runs a pinned mise",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_installScriptCmd).Standalone()

	generate_installScriptCmd.Flags().BoolP("help", "h", false, "Print help")
	generate_installScriptCmd.Flags().BoolP("localize", "l", false, "Keep mise data and cache in a project-local directory (`.mise` by default)")
	generate_installScriptCmd.Flags().String("localized-dir", ".mise", "Directory for mise data when using --localize, relative to the project")
	generate_installScriptCmd.Flags().StringP("version", "V", "", "mise version to pin; defaults to the release `mise self-update` would install")
	generate_installScriptCmd.Flags().Bool("windows", false, "Also write a Windows launcher, `<PATH>.cmd`")
	generate_installScriptCmd.Flags().StringP("write", "w", "", "Write the script to PATH (./bin/mise when no path is given) and make it executable, instead of printing it")
	generate_installScriptCmd.Flag("write").NoOptDefVal = " "
	generateCmd.AddCommand(generate_installScriptCmd)

	carapace.Gen(generate_installScriptCmd).FlagCompletion(carapace.ActionMap{
		"localized-dir": carapace.ActionDirectories(),
	})
}
