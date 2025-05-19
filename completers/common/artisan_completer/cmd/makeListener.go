package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeListenerCmd = &cobra.Command{
	Use:   "make:listener [-e|--event [EVENT]] [-f|--force] [--queued] [--test] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new event listener class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeListenerCmd).Standalone()

	makeListenerCmd.Flags().String("event", "", "The event class being listened for")
	makeListenerCmd.Flags().Bool("force", false, "Create the class even if the listener already exists")
	makeListenerCmd.Flags().Bool("pest", false, "Generate an accompanying Pest test for the Listener")
	makeListenerCmd.Flags().Bool("phpunit", false, "Generate an accompanying PHPUnit test for the Listener")
	makeListenerCmd.Flags().Bool("queued", false, "Indicates the event listener should be queued")
	makeListenerCmd.Flags().Bool("test", false, "Generate an accompanying Test test for the Listener")
	rootCmd.AddCommand(makeListenerCmd)
}
