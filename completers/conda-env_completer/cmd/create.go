package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an environment based on an environment definition file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().BoolP("dry-run", "d", false, "Only display what can be done with the current command, arguments, and other flags")
	createCmd.Flags().StringP("file", "f", "", "Environment definition file")
	createCmd.Flags().Bool("force", false, "Force creation of environment")
	createCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	createCmd.Flags().BoolP("insecure", "k", false, "Allow conda to perform \"insecure\" SSL connections and transfers")
	createCmd.Flags().Bool("json", false, "Report all output as json")
	createCmd.Flags().StringP("name", "n", "", "Name of environment")
	createCmd.Flags().Bool("no-default-packages", false, "Ignore create_default_packages in the .condarc file")
	createCmd.Flags().Bool("offline", false, "Offline mode")
	createCmd.Flags().StringP("prefix", "p", "", "Full path to environment location")
	createCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar")
	createCmd.Flags().String("solver", "", "Choose which solver backend to use")
	createCmd.Flags().BoolP("use-index-cache", "C", false, "Use cache of channel index files, even if it has expired")
	createCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"file":   carapace.ActionFiles(),
		"solver": carapace.ActionValues("classic"),
	})
}
