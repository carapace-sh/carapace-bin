package cmd

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"

	os "github.com/rsteube/carapace-bin/pkg/actions/os"
)

var rootCmd = &cobra.Command{
	Use:                "carapace",
	Short:              "multi-shell multi-command argument completer",
	Long:               "https://github.com/rsteube/carapace-bin",
	DisableFlagParsing: true,
	Run:                func(cmd *cobra.Command, args []string) {},
}

func flagCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "carapace",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().String("bridge", "", "generic completion bridge")
	cmd.Flags().BoolP("help", "h", false, "help for carapace")
	cmd.Flags().Bool("list", false, "list completers")
	cmd.Flags().BoolP("version", "v", false, "version for carapace")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"bridge": carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return os.ActionPathExecutables().Invoke(c).Suffix("/").ToA()
			case 1:
				return carapace.ActionValuesDescribed(
					"argcomplete", "kislyuk/argcomplete based completion like gcloud",
					"cobra", "spf13/cobra based completions like docker and kubectl",
					"posener", "posener/complete based completion like hashicorp tools",
				)
			default:
				return carapace.ActionValues()
			}
		}),
	})
	return cmd
}

func posCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "carapace",
		DisableFlagParsing: true,
		Run:                func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).PositionalCompletion(
		ActionCompleters(),
		carapace.ActionValues("bash", "elvish", "export", "fish", "ion", "nushell", "oil", "powershell", "tcsh", "xonsh", "zsh"),
		carapace.ActionValues("_"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(c.Args[0])
		}),
	)
	return cmd
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				// TODO Batch can't be used yet as ActionInvoke needs locking mutex for os.Stdout
				f := carapace.ActionInvoke(flagCmd().Execute)
				p := carapace.ActionInvoke(posCmd().Execute)
				return f.Invoke(c).Merge(p.Invoke(c)).ToA()
			}

			if strings.HasPrefix(c.Args[0], "-") {
				return carapace.ActionInvoke(flagCmd().Execute)
			}
			return carapace.ActionInvoke(posCmd().Execute)
		}),
	)
}

func ActionCompleters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("carapace", "--list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			re := regexp.MustCompile("^(?P<name>[^ ]+) +(?P<description>.*)$")

			vals := make([]string, 0)
			for _, line := range lines {
				if re.MatchString(line) {
					matched := re.FindStringSubmatch(line)
					vals = append(vals, matched[1])
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
