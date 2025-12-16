package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/number"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "svgcleaner <in-file> <out-file>",
	Short: "clean up your SVG files from the unnecessary data",
	Long:  "https://github.com/RazrFalcon/svgcleaner",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("allow-bigger-file", false, "Allow bigger file")
	rootCmd.Flags().String("apply-transform-to-gradients", "", "Apply transformations to gradients")
	rootCmd.Flags().String("apply-transform-to-paths", "", "Apply transformations to paths")
	rootCmd.Flags().String("apply-transform-to-shapes", "", "Apply transformations to shapes")
	rootCmd.Flags().String("convert-segments", "", "Convert path segments into shorter ones")
	rootCmd.Flags().String("convert-shapes", "", "Convert basic shapes into paths")
	rootCmd.Flags().String("coordinates-precision", "", "Set coordinates numeric precision")
	rootCmd.Flags().Bool("copy-on-error", false, "Copy an original file to the destination on error")
	rootCmd.Flags().String("group-by-style", "", "Group elements by equal styles")
	rootCmd.Flags().BoolP("help", "h", false, "Print help information")
	rootCmd.Flags().String("indent", "", "Set XML nodes indent")
	rootCmd.Flags().String("join-arcto-flags", "", "Join ArcTo flags")
	rootCmd.Flags().String("join-style-attributes", "", "Join presentational attributes")
	rootCmd.Flags().String("list-separator", "", "Set number list separator")
	rootCmd.Flags().String("merge-gradients", "", "Merge gradients")
	rootCmd.Flags().Bool("multipass", false, "Clean a file multiple times")
	rootCmd.Flags().Bool("no-defaults", false, "Reset default flags to 'no'")
	rootCmd.Flags().String("paths-coordinates-precision", "", "Set path's coordinates numeric precision")
	rootCmd.Flags().String("paths-to-relative", "", "Convert path segments into relative ones")
	rootCmd.Flags().String("properties-precision", "", "Set properties numeric precision")
	rootCmd.Flags().Bool("quiet", false, "Show only warnings and errors")
	rootCmd.Flags().String("regroup-gradient-stops", "", "Regroup gradient 'stop' elements")
	rootCmd.Flags().String("remove-comments", "", "Remove XML comments")
	rootCmd.Flags().String("remove-declarations", "", "Remove XML declarations")
	rootCmd.Flags().String("remove-default-attributes", "", "Remove attributes with default values")
	rootCmd.Flags().String("remove-desc", "", "Remove 'desc' element")
	rootCmd.Flags().String("remove-dupl-cmd-in-paths", "", "Remove subsequent segments command from paths")
	rootCmd.Flags().String("remove-dupl-fegaussianblur", "", "Remove duplicated 'feGaussianBlur' elements")
	rootCmd.Flags().String("remove-dupl-lineargradient", "", "Remove duplicated 'linearGradient' elements")
	rootCmd.Flags().String("remove-dupl-radialgradient", "", "Remove duplicated 'radialGradient' elements")
	rootCmd.Flags().String("remove-gradient-attributes", "", "Remove inheritable gradient attributes")
	rootCmd.Flags().String("remove-invalid-stops", "", "Remove invalid 'stop' elements")
	rootCmd.Flags().String("remove-invisible-elements", "", "Remove invisible elements")
	rootCmd.Flags().String("remove-metadata", "", "Remove 'metadata' element")
	rootCmd.Flags().String("remove-needless-attributes", "", "Remove attributes that doesn't belong to this element")
	rootCmd.Flags().String("remove-nonsvg-attributes", "", "Remove non-SVG attributes")
	rootCmd.Flags().String("remove-nonsvg-elements", "", "Remove non-SVG elements")
	rootCmd.Flags().String("remove-text-attributes", "", "Remove text-related attributes if there is no text")
	rootCmd.Flags().String("remove-title", "", "Remove 'title' element")
	rootCmd.Flags().String("remove-unreferenced-ids", "", "Remove unreferenced 'id' attributes")
	rootCmd.Flags().String("remove-unresolved-classes", "", "Remove unresolved classes from 'class' attributes")
	rootCmd.Flags().String("remove-unused-coordinates", "", "Remove unused coordinate attributes")
	rootCmd.Flags().String("remove-unused-defs", "", "Remove unused referenced elements")
	rootCmd.Flags().String("remove-unused-segments", "", "Remove unused path segments")
	rootCmd.Flags().String("remove-version", "", "Remove 'version' and 'baseProfile' attributes")
	rootCmd.Flags().String("remove-xmlns-xlink-attribute", "", "Remove an unused 'xmlns:xlink' attribute")
	rootCmd.Flags().String("resolve-use", "", "Resolve 'use' elements")
	rootCmd.Flags().String("simplify-transforms", "", "Simplify transform matrices")
	rootCmd.Flags().BoolP("stdout", "c", false, "Print result to the standard output")
	rootCmd.Flags().String("transforms-precision", "", "Set transform values numeric precision")
	rootCmd.Flags().String("trim-colors", "", "Use #RGB notation")
	rootCmd.Flags().String("trim-ids", "", "Trim 'id' attributes")
	rootCmd.Flags().String("trim-paths", "", "Use compact notation for paths")
	rootCmd.Flags().String("ungroup-defs", "", "Ungroup 'defs' element")
	rootCmd.Flags().String("ungroup-groups", "", "Ungroup groups")
	rootCmd.Flags().String("use-implicit-cmds", "", "Use implicit LineTo commands")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{

		"apply-transform-to-gradients": actionOption(),
		"apply-transform-to-paths":     actionOption(),
		"apply-transform-to-shapes":    actionOption(),
		"convert-segments":             actionOption(),
		"convert-shapes":               actionOption(),
		"group-by-style":               actionOption(),
		"indent":                       carapace.ActionValues("none", "0", "1", "2", "3", "4", "tabs").StyleF(style.ForKeyword),
		"join-arcto-flags":             actionOption(),
		"join-style-attributes":        carapace.ActionValues("no", "some", "all").StyleF(style.ForKeyword),
		"list-separator":               carapace.ActionValues("space", "comma", "comma-space"),
		"merge-gradients":              actionOption(),
		"paths-to-relative":            actionOption(),
		"regroup-gradient-stops":       actionOption(),
		"remove-comments":              actionOption(),
		"remove-declarations":          actionOption(),
		"remove-default-attributes":    actionOption(),
		"remove-desc":                  actionOption(),
		"remove-dupl-cmd-in-paths":     actionOption(),
		"remove-dupl-fegaussianblur":   actionOption(),
		"remove-dupl-lineargradient":   actionOption(),
		"remove-dupl-radialgradient":   actionOption(),
		"remove-gradient-attributes":   actionOption(),
		"remove-invalid-stops":         actionOption(),
		"remove-invisible-elements":    actionOption(),
		"remove-metadata":              actionOption(),
		"remove-needless-attributes":   actionOption(),
		"remove-nonsvg-attributes":     actionOption(),
		"remove-nonsvg-elements":       actionOption(),
		"remove-text-attributes":       actionOption(),
		"remove-title":                 actionOption(),
		"remove-unreferenced-ids":      actionOption(),
		"remove-unresolved-classes":    actionOption(),
		"remove-unused-coordinates":    actionOption(),
		"remove-unused-defs":           actionOption(),
		"remove-unused-segments":       actionOption(),
		"remove-version":               actionOption(),
		"remove-xmlns-xlink-attribute": actionOption(),
		"resolve-use":                  actionOption(),
		"simplify-transforms":          actionOption(),
		"transforms-precision":         number.ActionRange(number.RangeOpts{Start: 1, End: 12}),
		"trim-colors":                  actionOption(),
		"trim-ids":                     actionOption(),
		"trim-paths":                   actionOption(),
		"ungroup-defs":                 actionOption(),
		"ungroup-groups":               actionOption(),
		"use-implicit-cmds":            actionOption(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}

func actionOption() carapace.Action {
	return carapace.ActionValues("true", "false", "yes", "no", "y", "n").StyleF(style.ForKeyword)
}
