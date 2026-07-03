package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "set",
	Short: "Set or unset values of shell options and positional parameters",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "array assignment")
	rootCmd.Flags().BoolS("e", "e", false, "exit immediately on error")
	rootCmd.Flags().BoolS("f", "f", false, "disable file name generation")
	rootCmd.Flags().BoolS("h", "h", false, "remember command locations")
	rootCmd.Flags().BoolS("k", "k", false, "assign arguments as environment variables")
	rootCmd.Flags().BoolS("m", "m", false, "enable job control")
	rootCmd.Flags().BoolS("n", "n", false, "read but do not execute")
	rootCmd.Flags().StringS("o", "o", "", "set the option corresponding to option-name")
	rootCmd.Flags().BoolS("s", "s", false, "sort the specified arguments before assigning them")
	rootCmd.Flags().BoolS("t", "t", false, "exit after one command")
	rootCmd.Flags().BoolS("u", "u", false, "treat unset variables as error")
	rootCmd.Flags().BoolS("v", "v", false, "print shell input lines")
	rootCmd.Flags().BoolS("x", "x", false, "print commands and arguments")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"A": carapace.ActionValues("name", "array"),
		"o": carapace.ActionValues(
			"aliases", "aliasfuncdef", "allexport", "alwayslastprompt", "alwaystoend",
			"appendcreate", "appendhistory", "autocd", "autocontinue", "autolist",
			"automenu", "autonamedirs", "autoparamkeys", "autoparamslash", "autopushd",
			"autoremoveslash", "autoresume", "badpattern", "banghist", "bareglobqual",
			"bashautolist", "bashrematch", "beep", "bgnice", "braceccl",
			"bsdecho", "caseglob", "casematch", "casepaths", "cbases",
			"cdablevars", "cdsilent", "chasedots", "chaselinks", "checkjobs",
			"checkrunningjobs", "clobber", "clobberempty", "combiningchars", "completealiases",
			"completeinword", "continueonerror", "correct", "correctall", "cprecedences",
			"cshjunkiehistory", "cshjunkieloops", "cshjunkiequotes", "cshnullcmd", "cshnullglob",
			"debugbeforecmd", "dvorak", "emacs", "equals", "errexit",
			"errreturn", "evallineno", "exec", "extendedglob", "extendedhistory",
			"flowcontrol", "forcefloat", "functionargzero", "glob", "globalexport",
			"globalrcs", "globassign", "globcomplete", "globdots", "globstarshort",
			"globsubst", "hashcmds", "hashexecutablesonly", "hashlistall", "histallowclobber",
			"histbeep", "histexpiredupsfirst", "histfcntllock", "histfindnodups", "histignorealldups",
			"histignoredups", "histignorespace", "histlexwords", "histnofunctions", "histnostore",
			"histreduceblanks", "histsavebycopy", "histsavenodups", "histsubstpattern", "histverify",
			"hup", "ignorebraces", "ignoreclosebraces", "ignoreeof", "incappendhistory",
			"incappendhistorytime", "interactive", "interactivecomments", "ksharrays", "kshautoload",
			"kshglob", "kshoptionprint", "kshtypeset", "kshzerosubscript", "listambiguous",
			"listbeep", "listpacked", "listrowsfirst", "listtypes", "localloops",
			"localoptions", "localpatterns", "localtraps", "login", "longlistjobs",
			"magicequalsubst", "mailwarning", "markdirs", "menucomplete", "monitor",
			"multibyte", "multifuncdef", "multios", "nomatch", "notify",
			"nullglob", "numericglobsort", "octalzeroes", "overstrike", "pathdirs",
			"pathscript", "pipefail", "posixaliases", "posixargzero", "posixbuiltins",
			"posixcd", "posixidentifiers", "posixjobs", "posixstrings", "posixtraps",
			"printeightbit", "printexitvalue", "privileged", "promptbang", "promptcr",
			"promptpercent", "promptsp", "promptsubst", "pushdignoredups", "pushdminus",
			"pushdsilent", "pushdtohome", "rcexpandparam", "rcquotes", "rcs",
			"recexact", "rematchpcre", "restricted", "rmstarsilent", "rmstarwait",
			"sharehistory", "shfileexpansion", "shglob", "shinstdin", "shnullcmd",
			"shoptionletters", "shortloops", "shortrepeat", "shwordsplit", "singlecommand",
			"singlelinezle", "sourcetrace", "sunkeyboardhack", "transientrprompt", "trapsasync",
			"typesetsilent", "typesettounset", "unset", "verbose", "vi",
			"warncreateglobal", "warnnestedvar", "xtrace", "zle",
		),
	})
}
