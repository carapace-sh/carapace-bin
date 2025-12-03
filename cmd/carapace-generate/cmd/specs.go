package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
	"github.com/spf13/cobra"
)

var specsCmd = &cobra.Command{
	Use:   "specs DIR",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		completers, err := completer.ReadSpecs(args[0], cmd.Flag("group").Value.String(), false)
		if err != nil {
			return err
		}

		completers.SortVariants()
		m, err := json.Marshal(completers)
		if err != nil {
			return err
		}

		fmt.Println(string(m))
		return nil
	},
}

func init() {
	carapace.Gen(specsCmd).Standalone()

	specsCmd.Flags().String("group", "", "spec group")
	rootCmd.AddCommand(specsCmd)

	carapace.Gen(specsCmd).FlagCompletion(carapace.ActionMap{
		"group": carapace.ActionValues("user", "system"),
	})

	carapace.Gen(specsCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionValues("linux", "darwin", "windows"), // TODO others, termux as well?
	)
}
