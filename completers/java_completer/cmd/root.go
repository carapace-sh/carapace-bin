package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/java_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "java",
	Short: "Launches a Java application",
	Long:  "https://en.wikipedia.org/wiki/Java_(programming_language)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	for _, arg := range os.Args {
		// yuck - these are flags that require an argument but don't allow space between name and value
		// which essentially makes them new flags. To prevent unkown flag errors these are added as fake
		// flags just as they are.
		for _, flag := range []string{"-D", "-Xms", "-Xmx", "-Xss"} {
			if strings.HasPrefix(arg, flag) && len(arg) > 2 {
				fakeflag := strings.SplitN(arg[1:], "=", 2)[0]
				if rootCmd.Flag(fakeflag) == nil {
					rootCmd.Flags().String(fakeflag, "", "") // fake flag to prevent errors
					rootCmd.Flag(fakeflag).NoOptDefVal = " "
				}
			}
		}
	}

	carapace.Override(carapace.Opts{
		LongShorthand:   true,
		OptArgDelimiter: ":",
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("?", false, "print help message")
	rootCmd.Flags().StringArray("D", []string{}, "set a system property")
	rootCmd.Flags().Bool("X", false, "print help on non-standard options")
	rootCmd.Flags().String("agentlib", "", "load native agent library")
	rootCmd.Flags().String("agentpath", "", "load native agent library by full pathname")
	rootCmd.Flags().String("classpath", "", "class search path of directories and zip/jar files")
	rootCmd.Flags().String("cp", "", "class search path of directories and zip/jar files")
	rootCmd.Flags().Bool("d32", false, "use a 32-bit data model if available")
	rootCmd.Flags().Bool("d64", false, "use a 64-bit data model if available")
	rootCmd.Flags().String("da", "", "disable assertions with specified granularity")
	rootCmd.Flags().String("disableassertions", "", "disable assertions with specified granularity")
	rootCmd.Flags().Bool("disablesystemassertions", false, "disable system assertions")
	rootCmd.Flags().Bool("dsa", false, "disable system assertions")
	rootCmd.Flags().String("ea", "", "enable assertions with specified granularity")
	rootCmd.Flags().String("enableassertions", "", "enable assertions with specified granularity")
	rootCmd.Flags().Bool("enablesystemassertions", false, "enable system assertions")
	rootCmd.Flags().Bool("esa", false, "enable system assertions")
	rootCmd.Flags().Bool("help", false, "print help message")
	rootCmd.Flags().String("jar", "", "jar file to execute")
	rootCmd.Flags().String("javaagent", "", "load Java programming language agent, see java.lang.instrument")
	rootCmd.Flags().Bool("server", false, "to select the \"server\" VM")
	rootCmd.Flags().Bool("showversion", false, "print product version and continue")
	rootCmd.Flags().String("splash", "", "show splash screen with specified image")
	rootCmd.Flags().String("verbose", "", "enable verbose output")
	rootCmd.Flags().Bool("version", false, "print product version and exit")

	rootCmd.Flags().StringArray("XX", []string{}, "Advanced options")
	rootCmd.Flags().Bool("Xbatch", false, "disable background compilation")
	rootCmd.Flags().String("Xbootclasspath", "", "set search path for bootstrap classes and resources")
	rootCmd.Flags().String("Xbootclasspath/a", "", "append to end of bootstrap class path")
	rootCmd.Flags().String("Xbootclasspath/p", "", "prepend in front of bootstrap class path")
	rootCmd.Flags().String("Xcheck", "", "perform additional checks")
	rootCmd.Flags().Bool("Xdiag", false, "show additional diagnostic messages")
	rootCmd.Flags().Bool("Xfuture", false, "enable strictest checks, anticipating future default")
	rootCmd.Flags().Bool("Xincgc", false, "enable incremental garbage collection")
	rootCmd.Flags().Bool("Xint", false, "interpreted mode execution only")
	rootCmd.Flags().String("Xloggc", "", "log GC status to a file with time stamps")
	rootCmd.Flags().Bool("Xmixed", false, "mixed mode execution (default)")
	rootCmd.Flags().String("Xms", "", "set initial Java heap size")
	rootCmd.Flags().String("Xmx", "", "set maximum Java heap size")
	rootCmd.Flags().Bool("Xnoclassgc", false, "disable class garbage collection")
	rootCmd.Flags().Bool("Xprof", false, "output cpu profiling data")
	rootCmd.Flags().Bool("Xrs", false, "reduce use of OS signals by Java/VM (see documentation)")
	rootCmd.Flags().String("Xshare", "", "set shared data usage")
	rootCmd.Flags().String("XshowSettings", "", "show all settings and continue")
	rootCmd.Flags().String("Xss", "", "set java thread stack size")

	rootCmd.Flag("D").NoOptDefVal = " "
	rootCmd.Flag("agentlib").NoOptDefVal = " "
	rootCmd.Flag("agentpath").NoOptDefVal = " "
	rootCmd.Flag("da").NoOptDefVal = " "
	rootCmd.Flag("disableassertions").NoOptDefVal = " "
	rootCmd.Flag("ea").NoOptDefVal = " "
	rootCmd.Flag("enableassertions").NoOptDefVal = " "
	rootCmd.Flag("javaagent").NoOptDefVal = " "
	rootCmd.Flag("splash").NoOptDefVal = " "
	rootCmd.Flag("verbose").NoOptDefVal = " "
	rootCmd.Flag("Xbootclasspath").NoOptDefVal = " "
	rootCmd.Flag("Xbootclasspath/a").NoOptDefVal = " "
	rootCmd.Flag("Xbootclasspath/p").NoOptDefVal = " "
	rootCmd.Flag("Xloggc").NoOptDefVal = " "
	rootCmd.Flag("Xms").NoOptDefVal = " "
	rootCmd.Flag("Xmx").NoOptDefVal = " "
	rootCmd.Flag("Xss").NoOptDefVal = " "
	rootCmd.Flag("Xcheck").NoOptDefVal = " "
	rootCmd.Flag("Xshare").NoOptDefVal = " "
	rootCmd.Flag("XshowSettings").NoOptDefVal = " "
	rootCmd.Flag("XX").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"XX": action.ActionAdvanced(),
		"Xbootclasspath": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles(".zip", ".jar")
		}),
		"Xbootclasspath/a": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles(".zip", ".jar")
		}),
		"Xbootclasspath/p": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles(".zip", ".jar")
		}),
		"Xcheck": carapace.ActionValuesDescribed(
			"jni", "perform additional checks for JNI functions",
		),
		"Xloggc": carapace.ActionFiles(),
		"Xshare": carapace.ActionValuesDescribed(
			"auto", "use shared class data if possible",
			"jni", "perform additional checks for JNI functions",
			"off", "do not attempt to use shared class data",
			"on", "require using shared class data, otherwise fail",
		),
		"XshowSettings": carapace.ActionValuesDescribed(
			"all", "show all settings and continue",
			"vm", "show all vm related settings and continue",
			"system", "show host system or container configuration and continue",
			"properties", "show all property settings and continue",
			"locale", "show all locale related settings and continue",
		),
		"agentpath": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles(".jar", ".zip")
			default:
				return carapace.ActionValues()
			}
		}),
		"classpath": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"cp": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"da": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return ActionClasspathClasses(rootCmd)
		}),
		"disableassertions": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return ActionClasspathClasses(rootCmd)
		}),
		"ea": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return ActionClasspathClasses(rootCmd)
		}),
		"enableassertions": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return ActionClasspathClasses(rootCmd)
		}),
		"jar": carapace.ActionFiles(".jar", ".zip"),
		"javaagent": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles(".jar", ".zip")
			default:
				return carapace.ActionValues()
			}
		}),
		"splash": carapace.ActionFiles(),
		"verbose": carapace.ActionValuesDescribed(
			"class", "Displays information about each loaded class",
			"gc", "Displays information about each garbage collection (GC) event",
			"jni", "Displays information about the use of native methods and other Java Native Interface (JNI) activity",
		),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !rootCmd.Flag("jar").Changed {
				return ActionClasspathClasses(rootCmd)
			}
			return carapace.ActionValues()
		}),
	)
}

func ActionClasspathClasses(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		paths := make([]string, 0)
		if f := cmd.Flag("cp"); f.Changed {
			paths = append(paths, strings.Split(f.Value.String(), ":")...)
		}
		if f := cmd.Flag("classpath"); f.Changed {
			paths = append(paths, strings.Split(f.Value.String(), ":")...)
		}
		if f := cmd.Flag("jar"); f.Changed {
			paths = append(paths, f.Value.String())
		}

		files := make([]string, 0)
		for _, path := range paths {
			if f, err := os.Stat(path); err == nil {
				if !f.IsDir() {
					files = append(files, path)
				} else {
					if fileInfos, err := ioutil.ReadDir(path); err == nil {
						for _, file := range fileInfos {
							files = append(files, fmt.Sprintf("%v/%v", path, file.Name()))
						}
					}
				}
			}
		}

		actions := make([]carapace.Action, 0)
		for _, file := range files {
			if strings.HasSuffix(file, ".jar") ||
				strings.HasSuffix(file, ".zip") {
				actions = append(actions, fs.ActionJarFileClasses(file))
			}
		}
		return carapace.Batch(actions...).Invoke(c).Merge().ToMultiPartsA(".")
	})
}
