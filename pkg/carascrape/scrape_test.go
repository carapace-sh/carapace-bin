package carascrape

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestScrape(t *testing.T) {
	cmd := &cobra.Command{
		Use:   "main",
		Short: "short description",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	cmd.PersistentFlags().BoolP("boolflag", "b", false, "bool flag example")
	cmd.Flags().BoolSlice("boolsliceflag", []bool{true, false}, "boolslice flag example")
	cmd.Flags().String("stringflag", "defaultValue", "string flag example")
	cmd.Flags().StringSlice("stringsliceflag", []string{"defaultValue", "defaultValue2"}, "stringslice flag example")
	cmd.Flags().StringArray("stringarrayflag", []string{"defaultValue", "defaultValue2"}, "stringarray flag example")
	cmd.Flags().Int("intflag", 0, "string flag example")
	cmd.Flags().IntSlice("intsliceflag", []int{0, 1}, "intlice flag example")

	subCmd := &cobra.Command{
		Use:   "sub",
		Short: "short description",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	subCmd.Flags().String("substringflag", "defaultValue", "string flag example")
	subCmd.Flag("substringflag").NoOptDefVal = "nooptval"
	cmd.AddCommand(subCmd)

	Scrape(cmd)
}
