package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shopt",
	Short: "Set and unset shell options",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-shopt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("o", "o", false, "restrict optnames to those defined for use with set -o")
	rootCmd.Flags().BoolS("p", "p", false, "print each shell option with an indication of its status")
	rootCmd.Flags().BoolS("q", "q", false, "suppress output")
	rootCmd.Flags().BoolS("s", "s", false, "enable (set) each optname")
	rootCmd.Flags().BoolS("u", "u", false, "disable (unset) each optname")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flags().Lookup("o").Changed {
				return carapace.ActionValuesDescribed(
					"allexport", "same as -a",
					"braceexpand", "same as -B",
					"emacs", "use an emacs-style line editing interface",
					"errexit", "same as -e",
					"errtrace", "same as -E",
					"functrace", "same as -T",
					"hashall", "same as -h",
					"histexpand", "same as -H",
					"history", "enable command history",
					"ignoreeof", "shell does not exit on EOF",
					"keyword", "same as -k",
					"monitor", "same as -m",
					"noclobber", "same as -C",
					"noexec", "same as -n",
					"noglob", "same as -f",
					"nolog", "do not save function definitions in history",
					"notify", "same as -b",
					"nounset", "same as -u",
					"onecmd", "same as -t",
					"physical", "same as -P",
					"pipefail", "return value of pipeline is the rightmost command to exit with non-zero status",
					"posix", "change the behavior of bash to match the POSIX standard",
					"privileged", "same as -p",
					"verbose", "same as -v",
					"vi", "use a vi-style line editing interface",
					"xtrace", "same as -x",
				).Filter(c.Args...)
			}
			return carapace.ActionValuesDescribed(
				"assoc_expand_once", "expand associative array subscripts only once",
				"autocd", "a command name is treated as a candidate for cd",
				"cdable_vars", "a word that is not a directory is assumed to be a variable name",
				"cdspell", "minor errors in the spelling of a directory are corrected",
				"checkhash", "check whether a command found in the hash table exists",
				"checkjobs", "check for running jobs before exiting an interactive shell",
				"checkwinsize", "check window size after each command",
				"cmdhist", "save all lines of a multi-line command in the same history entry",
				"compat31", "compatibility mode for bash 3.1",
				"compat32", "compatibility mode for bash 3.2",
				"compat40", "compatibility mode for bash 4.0",
				"compat41", "compatibility mode for bash 4.1",
				"compat42", "compatibility mode for bash 4.2",
				"compat43", "compatibility mode for bash 4.3",
				"compat44", "compatibility mode for bash 4.4",
				"complete_fullquote", "quote all readline completions",
				"direxpand", "expand directory names during word completion",
				"dirspell", "attempt spelling correction on directory names during word completion",
				"dotglob", "include filenames beginning with a '.' in pathname expansion",
				"execfail", "non-interactive shell does not exit on exec failure",
				"expand_aliases", "expand aliases",
				"extdebug", "enable extended debugging",
				"extglob", "enable extended pattern matching",
				"extquote", "allow $'...' and $\"...\" quoting",
				"failglob", "pattern expansion that matches no files produces an error",
				"force_fignore", "suffixes ignored by FIGNORE also affect word completion",
				"globasciiranges", "use ASCII collation order for range expressions in pattern matching",
				"globskipdots", "skip . and .. in pathname expansion",
				"globstar", "** matches all files and directories including subdirectories",
				"gnu_errfmt", "use GNU error message format",
				"histappend", "append to history file rather than overwrite",
				"histreedit", "allow user to re-edit a failed history substitution",
				"histverify", "allow user to edit a history substitution result",
				"hostcomplete", "attempt hostname completion",
				"huponexit", "send SIGHUP to all jobs on exit",
				"inherit_errexit", "command substitution inherits the errexit setting",
				"interactive_comments", "allow comments in interactive shells",
				"lastpipe", "run the last command of a pipeline in the current shell",
				"lithist", "save multi-line commands with embedded newlines",
				"localvar_inherit", "local variables inherit the value and attributes of a variable of the same name",
				"localvar_unset", "local variables inherit the value and attributes of a variable of the same name",
				"login_shell", "shell is a login shell",
				"mailwarn", "warn if mail has been read and new mail has arrived",
				"nocaseglob", "match filenames in a case-insensitive fashion",
				"nocasematch", "match patterns in a case-insensitive fashion",
				"no_empty_cmd_completion", "do not attempt to search PATH for completions on an empty line",
				"noexpand_translation", "do not expand translated strings",
				"nullglob", "allow patterns which match no files to expand to a null string",
				"patsub_replacement", "perform pattern substitution on the right-hand side of variable expansion",
				"progcomp", "enable programmable completion",
				"progcomp_alias", "allow programmable completion to use aliases",
				"promptvars", "prompt strings undergo variable expansion",
				"restricted_shell", "shell is running in restricted mode",
				"shift_verbose", "shift prints a message when the shift count exceeds the number of positional parameters",
				"sourcepath", "source uses PATH to find files",
				"varredir_close", "close file descriptors created by variable redirections",
				"xpg_echo", "echo expands escape sequences by default",
			).Filter(c.Args...)
		}),
	)
}
