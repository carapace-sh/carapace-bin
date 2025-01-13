package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp [OPTIONS] SERVICE:SRC_PATH DEST_PATH|-",
	Short: "Copy files/folders between a service container and the local filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpCmd).Standalone()

	cpCmd.Flags().Bool("all", false, "Include containers created by the run command")
	cpCmd.Flags().BoolP("archive", "a", false, "Archive mode (copy all uid/gid information)")
	cpCmd.Flags().BoolP("follow-link", "L", false, "Always follow symbol link in SRC_PATH")
	cpCmd.Flags().String("index", "", "Index of the container if service has multiple replicas")
	rootCmd.AddCommand(cpCmd)

	// TODO index flag

	carapace.Gen(cpCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionServices(cpCmd).Suffix(":")
				default:
					if index, err := cpCmd.Flags().GetInt("index"); err != nil {
						return carapace.ActionMessage(err.Error())
					} else {
						return action.ActionFiles(cpCmd, c.Parts[0], index)
					}
				}
			}).UnlessF(condition.CompletingPath),
		).ToA(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !condition.File(c.Args[0])(c) {
				return carapace.ActionFiles()
			}

			return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionServices(cpCmd).Suffix(":").UnlessF(condition.CompletingPath)
				case 1:
					if index, err := cpCmd.Flags().GetInt("index"); err != nil {
						return carapace.ActionMessage(err.Error())
					} else {
						return action.ActionFiles(cpCmd, c.Parts[0], index)
					}
				default:
					return carapace.ActionValues()
				}
			})
		}),
	)
}
