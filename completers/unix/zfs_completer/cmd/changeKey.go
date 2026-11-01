package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var changeKeyCmd = &cobra.Command{
	Use:     "change-key [-l] [-o property=value]... filesystem | change-key -i [-l] filesystem",
	Short:   "change encryption key for a dataset",
	GroupID: "encryption",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changeKeyCmd).Standalone()

	changeKeyCmd.Flags().BoolS("i", "i", false, "inherit key from parent")
	changeKeyCmd.Flags().BoolS("l", "l", false, "load key before changing it")
	changeKeyCmd.Flags().StringArrayS("o", "o", nil, "set encryption property")

	rootCmd.AddCommand(changeKeyCmd)

	carapace.Gen(changeKeyCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("keyformat", "keylocation", "pbkdf2iters").Suffix("=")
			default:
				return zfs.ActionPropertyValues(c.Parts[0])
			}
		}),
	})

	carapace.Gen(changeKeyCmd).PositionalCompletion(
		zfs.ActionFilesystems(),
	)
}
