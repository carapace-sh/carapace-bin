package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsusb",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("s", "s", "", "Show only devices with specified device and/or bus numbers")
	rootCmd.Flags().StringS("d", "d", "", "Show only devices with the specified vendor and product ID numbers")
	rootCmd.Flags().StringS("D", "D", "", "Selects which device lsusb will examine")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage and help")
	rootCmd.Flags().BoolP("tree", "t", false, "Dump the physical USB device hierarchy as a tree")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase verbosity")
	rootCmd.Flags().BoolP("version", "V", false, "Show version of program")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"s": ActionDevices(),
		"d": ActionProducts(),
		"D": carapace.ActionFiles(),
	})
}

func ActionDevices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if output, err := exec.Command("lsusb").Output(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				r := regexp.MustCompile(`^Bus (?P<bus>\d+) Device (?P<device>\d+): ID (?P<vendor>[^ ]+):(?P<product>[^ ]+) (?P<name>.*)$`)

				vals := make([]string, 0)
				for _, line := range strings.Split(string(output), "\n") {
					if r.MatchString(line) {
						matches := r.FindStringSubmatch(line)
						vals = append(vals, fmt.Sprintf("%v:%v", matches[1], matches[2]), matches[5])
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			}
		}).Invoke(c).ToMultiPartsA(":")
	})
}

func ActionProducts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if output, err := exec.Command("lsusb").Output(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				r := regexp.MustCompile(`^Bus (?P<bus>\d+) Device (?P<device>\d+): ID (?P<vendor>[^ ]+):(?P<product>[^ ]+) (?P<name>.*)$`)

				vals := make([]string, 0)
				for _, line := range strings.Split(string(output), "\n") {
					if r.MatchString(line) {
						matches := r.FindStringSubmatch(line)
						vals = append(vals, fmt.Sprintf("%v:%v", matches[3], matches[4]), matches[5])
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			}
		}).Invoke(c).ToMultiPartsA(":")
	})
}
