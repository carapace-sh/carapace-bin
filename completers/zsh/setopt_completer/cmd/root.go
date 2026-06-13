package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "setopt",
	Short: "Set shell options",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-setopt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("e", "e", false, "exit on error")
	rootCmd.Flags().BoolS("k", "k", false, "all arguments as keyword")
	rootCmd.Flags().BoolS("m", "m", false, "job control")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValues(
			"allexport", "alwayslastprompt", "alwaystoend", "autocd", "autocontinue",
			"autolist", "automenu", "autoparamkeys", "autoparamslash", "autopushd",
			"autoresume", "badpattern", "banghist", "beep", "bgnice", "braceccl",
			"cdablevars", "cdsilent", "chaselinks", "checkjobs", "checkrunningjobs",
			"clobber", "completeinword", "correct", "correctall", "dirsspell",
			"dotglob", "dvorak", "emacs", "errexit", "errreturn", "exec",
			"extendedglob", "extendedhistory", "flowcontrol", "forcefloat", "glob",
			"globassign", "globcomplete", "globdots", "globstarshort", "globsubst",
			"hashcmds", "hashdirs", "hashexecutablesonly", "hashlistall", "histallowclobber",
			"histbeep", "histexpiredupsfirst", "histfcntllock", "histfindnodups",
			"histignorealldups", "histignoredups", "histignorespace", "histlexwords",
			"histnofunctions", "histnostore", "histreduceblanks", "histsavenodups",
			"histsubstpattern", "history", "historyappend", "historysubsel",
			"hookstyle", "hostcomplete", "huponexit", "ignorebraces", "ignoreclosebraces",
			"incappendhistory", "incappendhistorytime", "interactivecomments", "interactiveimmediate",
			"interactivestart", "keeplogs", "listambiguous", "listbeep", "listpacked",
			"listrowsfirst", "localloops", "localoptions", "localpatterns", "login",
			"longlistjobs", "magicequalsubst", "mailwarning", "markdirs", "menucomplete",
			"monitor", "multios", "noclobber", "nocomments", "nodie", "noerrreturn",
			"noflowcontrol", "noglobalexport", "nohup", "notify", "nullglob",
			"numericglobsort", "octalzeroes", "overstrike", "pathdirs", "pipefail",
			"posixbuiltins", "posixident", "printeightbit", "prn", "promptbang",
			"promptcr", "promptpercent", "promptsp", "promptsubst", "quietzle",
			"rcquotes", "receivesignals", "restricted", "rmstarwait", "schedule",
			"shallowersetopt", "shglob", "shortloops", "showcomplete", "showtype",
			"silent", "singlecommands", "singlelinezle", "smartcomplete", "sunkeyboardhack",
			"transposechars", "trapsasync", "unset", "verbose", "vi", "warncreateglobal",
			"warnnestedvar", "xtrace", "zle",
		),
	)
}
