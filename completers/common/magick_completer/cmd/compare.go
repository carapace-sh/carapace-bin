package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "mathematically and visually annotate the difference between an image and its reconstruction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compareCmd).Standalone()

	compareCmd.Flags().CountS("adjoin", "adjoin", "join images into a single multi-image file")
	compareCmd.Flags().CountS("alpha", "alpha", "on, activate, off, deactivate, set, opaque, copy\", transparent, extract, background, or shape the alpha channel")
	compareCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	compareCmd.Flags().CountS("auto-orient", "auto-orient", "automagically orient image")
	compareCmd.Flags().StringSliceS("background", "background", nil, "background color")
	compareCmd.Flags().StringSliceS("brightness-contrast", "brightness-contrast", nil, "improve brightness / contrast of the image")
	compareCmd.Flags().StringSliceS("channel", "channel", nil, "apply option to select image channels")
	compareCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "set image colorspace")
	compareCmd.Flags().StringSliceS("compose", "compose", nil, "set image composite operator")
	compareCmd.Flags().StringSliceS("crop", "crop", nil, "crop the image")
	compareCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	compareCmd.Flags().StringSliceS("decipher", "decipher", nil, "convert cipher pixels to plain")
	compareCmd.Flags().StringSliceS("define", "define", nil, "define one or more image format options")
	compareCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	compareCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	compareCmd.Flags().StringSliceS("dissimilarity-threshold", "dissimilarity-threshold", nil, "maximum distortion for (sub)image match (default 0.2)")
	compareCmd.Flags().StringSliceS("distort", "distort", nil, "distort image")
	compareCmd.Flags().StringSliceS("encipher", "encipher", nil, "convert plain pixels to cipher pixels")
	compareCmd.Flags().StringSliceS("extract", "extract", nil, "extract area from image")
	compareCmd.Flags().StringSliceS("fuzz", "fuzz", nil, "colors within this distance are considered equal")
	compareCmd.Flags().StringSliceS("gravity", "gravity", nil, "horizontal and vertical text placement")
	compareCmd.Flags().BoolS("help", "help", false, "print program options")
	compareCmd.Flags().StringSliceS("highlight-color", "highlight-color", nil, "emphasize pixel differences with this color")
	compareCmd.Flags().CountS("identify", "identify", "identify the format and characteristics of the image")
	compareCmd.Flags().StringSliceS("interlace", "interlace", nil, "type of image interlacing scheme")
	compareCmd.Flags().StringSliceS("level", "level", nil, "adjust the level of image contrast")
	compareCmd.Flags().StringSliceS("limit", "limit", nil, "pixel cache resource limit")
	compareCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	compareCmd.Flags().StringSliceS("lowlight-color", "lowlight-color", nil, "de-emphasize pixel differences with this color")
	compareCmd.Flags().StringSliceS("metric", "metric", nil, "measure differences between images with this metric.  The default metric is RMSE.")
	compareCmd.Flags().CountS("negate", "negate", "replace each pixel with its complementary color")
	compareCmd.Flags().StringSliceS("profile", "profile", nil, "add, delete, or apply an image profile")
	compareCmd.Flags().StringSliceS("quality", "quality", nil, "JPEG/MIFF/PNG compression level")
	compareCmd.Flags().StringSliceS("quantize", "quantize", nil, "reduce image colors in this colorspace")
	compareCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	compareCmd.Flags().StringSliceS("read-mask", "read-mask", nil, "associate a read mask with the image")
	compareCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	compareCmd.Flags().StringSliceS("repage", "repage", nil, "size and location of an image canvas")
	compareCmd.Flags().StringSliceS("resize", "resize", nil, "resize the image")
	compareCmd.Flags().CountS("respect-parentheses", "respect-parentheses", "settings remain in effect until parenthesis boundary.")
	compareCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	compareCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factor")
	compareCmd.Flags().StringSliceS("seed", "seed", nil, "seed a new sequence of pseudo-random numbers")
	compareCmd.Flags().CountS("separate", "separate", "separate an image channel into a grayscale image")
	compareCmd.Flags().StringSliceS("set", "set", nil, "set an image attribute")
	compareCmd.Flags().StringSliceS("sigmoidal-contrast", "sigmoidal-contrast", nil, "increase the contrast without saturating highlights or shadows")
	compareCmd.Flags().StringSliceS("similarity-threshold", "similarity-threshold", nil, "minimum distortion for (sub)image match (default 0.0)")
	compareCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	compareCmd.Flags().CountS("subimage-search", "subimage-search", "search for subimage")
	compareCmd.Flags().CountS("synchronize", "synchronize", "synchronize image to storage device")
	compareCmd.Flags().CountS("taint", "taint", "mark the image as modified")
	compareCmd.Flags().StringSliceS("transparent-color", "transparent-color", nil, "transparent color")
	compareCmd.Flags().CountS("trim", "trim", "trim image edges")
	compareCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	compareCmd.Flags().BoolS("version", "version", false, "print version information")
	compareCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "access method for pixels outside the boundaries of the image")
	compareCmd.Flags().StringSliceS("write-mask ", "write-mask ", nil, "associate a write mask with the image")
	rootCmd.AddCommand(compareCmd)

	carapace.Gen(compareCmd).FlagCompletion(carapace.ActionMap{
		"background":        action.ActionColors(),
		"channel":           action.ActionCompleteOption("Channel").UniqueList(","),
		"compose":           action.ActionCompleteOption("Compose"),
		"debug":             action.ActionCompleteOption("Debug").UniqueList(","),
		"decipher":          carapace.ActionFiles(),
		"distort":           action.ActionCompleteOption("Distort"),
		"encipher":          carapace.ActionFiles(),
		"gravity":           action.ActionCompleteOption("Gravity"),
		"highlight-color":   action.ActionColors(),
		"interlace":         action.ActionCompleteOption("Interlace"),
		"lowlight-color":    action.ActionColors(),
		"metric":            action.ActionCompleteOption("Metric"),
		"profile":           carapace.ActionFiles(),
		"quantize":          action.ActionCompleteOption("Colorspace"),
		"read-mask":         carapace.ActionFiles(),
		"transparent-color": action.ActionColors(),
		"virtual-pixel":     action.ActionCompleteOption("VirtualPixel"),
	})

	carapace.Gen(compareCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
