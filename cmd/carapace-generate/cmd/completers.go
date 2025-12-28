package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

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

		// TODO completers should be sorted for target GOOS
		var s string
		switch {
		case cmd.Flag("code").Changed:
			s = completers.Format("completers", cmd.Flag("tag").Value.String()) // TODO pass package by flag

		default:
			completers.SortVariants()
			m, err := json.Marshal(completers)
			if err != nil {
				return err
			}
			s = string(m)

		}

		if f := cmd.Flag("output"); f.Changed {
			if err := os.WriteFile(f.Value.String(), []byte(s), 0644); err != nil {
				return err
			}
			return exec.Command("go", "fmt", f.Value.String()).Run()
		}
		fmt.Println(s)
		return nil
	},
}

func init() {
	carapace.Gen(completersCmd).Standalone()

	completersCmd.Flags().Bool("code", false, "output go code")
	completersCmd.Flags().String("tag", "", "build tag for code output")
	completersCmd.Flags().String("output", "", "file to write to")
	rootCmd.AddCommand(completersCmd)

	carapace.Gen(completersCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionValues("linux", "darwin", "windows", "force_all"), // TODO others, termux as well?
	)
}
