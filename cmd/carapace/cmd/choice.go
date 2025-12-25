package cmd

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	carapacebin "github.com/carapace-sh/carapace-bin/pkg/actions/tools/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/choice"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var choiceCmd = &cobra.Command{
	Use:   "--choice [-d] [variant]...",
	Short: "list or edit choices",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			choices, err := choice.List(true)
			if err != nil {
				return err
			}
			for _, choice := range choices {
				fmt.Println(choice.Format())
			}
			return nil
		}

		switch cmd.Flag("delete").Changed {
		case true:
			for _, arg := range args {
				if err := choice.Unset(arg); err != nil {
					return err
				}
			}
		default:
			for _, arg := range args {
				if err := choice.Set(choice.Parse(arg)); err != nil {
					return err
				}
			}
		}
		return nil
	},
}

func init() {
	carapace.Gen(choiceCmd).Standalone()
	choiceCmd.Flags().SetInterspersed(false)

	choiceCmd.Flags().BoolP("delete", "d", false, "delete given choice(s)")

	carapace.Gen(choiceCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if choiceCmd.Flag("delete").Changed {
				choices, err := choice.List(true)
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				vals := make([]string, 0)
				for _, choice := range choices {
					vals = append(vals, choice.Name, choice.Format())
				}
				return carapace.ActionValuesDescribed(vals...)
			}

			return carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionMultiPartsN("/", 2, func(c carapace.Context) carapace.Action {
						switch len(c.Parts) {
						case 0:
							return carapace.Batch(
								carapace.ActionExecutables().Style(style.Dim),
								carapacebin.ActionNames(),
							).ToA().NoSpace()
						default:
							// TODO highlight by group
							return carapace.Batch(
								// TODO provide action in carapace-bridge
								carapace.ActionValuesDescribed(
									"argcomplete", "",
									// "argcomplete@v1", "",
									"aws", "",
									"bash", "",
									"carapace", "",
									"carapace-bin", "",
									"clap", "",
									"click", "",
									"cobra", "",
									"complete", "",
									"fish", "",
									"gcloud", "",
									"inshellisense", "",
									"kingpin", "",
									"kitten", "",
									"powershell", "",
									"urfavecli", "",
									// "urfavecli@v1", "",
									"yargs", "",
									"zsh", "",
								).Style(style.Magenta),
								carapacebin.ActionVariants(c.Parts[0]),
							).ToA().NoSpace()
						}
					})
				default:
					// TODO needs to support unknown bridges
					return carapacebin.ActionGroups(c.Parts[0])
				}
			})
		}),
	)
}
