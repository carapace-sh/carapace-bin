package cmd

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/completers"
	carapacebin "github.com/carapace-sh/carapace-bin/pkg/actions/tools/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-bridge/pkg/choices"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var choiceCmd = &cobra.Command{
	Use:   "--choice [-d] [variant]...",
	Short: "list or edit choices",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			choices, err := choices.List(true)
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
				if err := choices.Unset(arg); err != nil {
					return err
				}
			}
		default:
			for _, arg := range args {
				if err := choices.Set(choices.Parse(arg)); err != nil {
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
				list, err := choices.List(true)
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				vals := make([]string, 0)
				for _, choice := range list {
					vals = append(vals, choice.Name, choice.Format())
				}
				return carapace.ActionValuesDescribed(vals...).FilterArgs()
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
								carapace.ActionCallback(func(c carapace.Context) carapace.Action {
									filter := []string{"macro", "carapace-bin"}
									if m, err := completers.Completers(choices.Parse(c.Parts[0]), true); err == nil {
										if variants, ok := m[c.Parts[0]]; ok {
											for _, v := range variants {
												filter = append(filter, v.Variant)
											}
										}
									}
									return bridge.ActionBridges(c.Parts[0]).Filter(filter...).Suffix("@bridge")
								}),
								carapacebin.ActionVariants(c.Parts[0]).Style(style.Carapace.KeywordPositive),
							).ToA().NoSpace()
						}
					})
				default:
					// TODO needs to support unknown bridges (add bridge unless it's a variant that doesn't fit)
					return carapacebin.ActionGroups(c.Parts[0])
				}
			})
		}),
	)
}
