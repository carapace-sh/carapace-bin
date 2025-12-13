package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/carapace-sh/carapace"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/spf13/cobra"
)

var macrosCmd = &cobra.Command{
	Use:   "macros",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		macros, err := spec.ScanMacros(args...)
		if err != nil {
			return err
		}

		var s string
		switch {
		case cmd.Flag("code").Changed:
			s, err = macros.Format("actions")
			if err != nil {
				return err
			}
		default:
			m, err := json.Marshal(macros)
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
	carapace.Gen(macrosCmd).Standalone()

	macrosCmd.Flags().Bool("code", false, "output code")
	macrosCmd.Flags().String("output", "", "file to write to")
	rootCmd.AddCommand(macrosCmd)
}
