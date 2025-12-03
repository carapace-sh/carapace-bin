package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace-generate/pkg/completer"
	"github.com/spf13/cobra"
)

var completersCmd = &cobra.Command{
	Use:   "completers DIR [GOOS]",
	Short: "",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		completers, err := completer.ReadCompleters(args[0], args[1])
		if err != nil {
			return err
		}

		if cmd.Flag("code").Changed {
			fmt.Println(completers.Format("completers", cmd.Flag("tag").Value.String())) // TODO pass package by flag
			return nil
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
	carapace.Gen(completersCmd).Standalone()

	completersCmd.Flags().Bool("code", false, "output go code")
	completersCmd.Flags().String("tag", "", "build tag for code output")
	rootCmd.AddCommand(completersCmd)

	carapace.Gen(completersCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionValues("linux", "darwin", "windows"), // TODO others, termux as well?
	)
}
