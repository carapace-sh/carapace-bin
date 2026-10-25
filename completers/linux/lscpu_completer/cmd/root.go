package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lscpu",
	Short: "display information about the CPU architecture",
	Long:  "https://man7.org/linux/man-pages/man1/lscpu.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "print both online and offline CPUs (default for -e)")
	rootCmd.Flags().String("annotate", "", "annotate columns with a tooltip (always|never|auto)")
	rootCmd.Flags().String("arm-id", "", "print the known ARM implementers and their names")
	rootCmd.Flags().String("arm-model", "", "print the name of the given ARM implementer and model")
	rootCmd.Flags().BoolP("bytes", "B", false, "print sizes in bytes rather than in human readable format")
	rootCmd.Flags().StringP("caches", "C", "", "info about caches in extended readable format")
	rootCmd.Flags().StringP("extended", "e", "", "print out an extended readable format")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("hex", "x", false, "print hexadecimal masks rather than lists of CPUs")
	rootCmd.Flags().String("hierarchic", "", "use subsections in summary")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON for default or extended format")
	rootCmd.Flags().BoolP("list-columns", "H", false, "list the available columns")
	rootCmd.Flags().BoolP("offline", "c", false, "print offline CPUs only")
	rootCmd.Flags().BoolP("online", "b", false, "print online CPUs only (default for -p)")
	rootCmd.Flags().Bool("output-all", false, "print all available columns for -e, -p or -C")
	rootCmd.Flags().StringP("parse", "p", "", "print out a parsable format")
	rootCmd.Flags().BoolP("physical", "y", false, "print physical instead of logical IDs")
	rootCmd.Flags().BoolP("raw", "r", false, "use raw output format (for -e, -p and -C)")
	rootCmd.Flags().StringP("sysroot", "s", "", "use specified directory as system root")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	rootCmd.Flag("annotate").NoOptDefVal = " "
	rootCmd.Flag("arm-id").NoOptDefVal = " "
	rootCmd.Flag("caches").NoOptDefVal = " "
	rootCmd.Flag("extended").NoOptDefVal = " "
	rootCmd.Flag("hierarchic").NoOptDefVal = " "
	rootCmd.Flag("parse").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"annotate":   carapace.ActionValues("always", "auto", "never").StyleF(style.ForKeyword),
		"caches":     ActionCacheColumns().UniqueList(","),
		"extended":   ActionFormatColumns().UniqueList(","),
		"hierarchic": carapace.ActionValues("always", "auto", "never").StyleF(style.ForKeyword),
		"parse":      ActionFormatColumns().UniqueList(","),
		"sysroot":    carapace.ActionDirectories(),
	})
}

func ActionFormatColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ADDRESS", "physical address of a CPU",
		"BOGOMIPS", "crude measurement of CPU speed",
		"BOOK", "logical book number",
		"CACHE", "shows how caches are shared between CPUs",
		"CLUSTER", "logical cluster number",
		"CONFIGURED", "shows if the hypervisor has allocated the CPU",
		"CORE", "logical core number",
		"CPU", "logical CPU number",
		"DRAWER", "logical drawer number",
		"MAXMHZ", "shows the maximum MHz of the CPU",
		"MHZ", "shows the current MHz of the CPU",
		"MINMHZ", "shows the minimum MHz of the CPU",
		"MODELNAME", "shows CPU model name",
		"NODE", "logical NUMA node number",
		"ONLINE", "shows if Linux currently makes use of the CPU",
		"POLARIZATION", "CPU dispatching mode on virtual hardware",
		"SCALMHZ%", "shows scaling percentage of the CPU frequency",
		"SOCKET", "logical socket number",
	)
}

func ActionCacheColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ALL-SIZE", "size of all system caches",
		"ALLOC-POLICY", "allocation policy",
		"COHERENCY-SIZE", "minimum amount of data in bytes transferred from memory to cache",
		"LEVEL", "cache level",
		"NAME", "cache name",
		"ONE-SIZE", "size of one cache",
		"PHY-LINE", "number of physical cache line per cache t",
		"SETS", "number of sets in the cache; set lines has the same cache index",
		"TYPE", "cache type",
		"WAYS", "ways of associativity",
		"WRITE-POLICY", "write policy",
	)
}
