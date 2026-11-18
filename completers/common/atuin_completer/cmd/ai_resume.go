package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_resumeCmd = &cobra.Command{
	Use:    "resume",
	Short:  "`atuin threads resume`, as it was called before",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_resumeCmd).Standalone()

	ai_resumeCmd.Flags().Bool("as-is", false, "Resume this machine's copy of the thread QUERY names as it is, without catching it up with sync")
	ai_resumeCmd.Flags().String("branch", "", "The branch of the thread QUERY names to catch up to, to switch to with `--switch`, to fork from with `--fork`, or to continue with `--in`: `this`, `@<host id>`, or the start of its head's id. The newest by default; `--in`, and `--switch` with several branches to switch to, need it named")
	ai_resumeCmd.Flags().String("filter-mode", "", "The filter the picker opens in (default: global)")
	ai_resumeCmd.Flags().Bool("fork", false, "Fork the thread QUERY names (by id, or a unique id prefix): it is written out as a new session of the same agent, with the same history, and resumed. The original is left as it is")
	ai_resumeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ai_resumeCmd.Flags().String("in", "", "Continue the thread QUERY names (by id, or a unique id prefix) in another agent: it is written out as a new session there, tool calls carried over as its own where it has a like tool (notes where they weren't captured whole), and resumed. One that went on separately on several machines needs `--branch`")
	ai_resumeCmd.Flags().String("inline-height", "", "Height of the inline picker; 0 for fullscreen")
	ai_resumeCmd.Flags().String("keymap-mode", "", "Keymap mode, as passed by the shell widget")
	ai_resumeCmd.Flags().Bool("print", false, "Print the resume command instead of running it")
	ai_resumeCmd.Flags().Bool("shell-widget", false, "Report the result on stderr for the shell widget")
	ai_resumeCmd.Flags().Bool("switch", false, "Switch this machine's copy of the thread QUERY names to another branch (the one `--branch` names, else the only one it can be switched to), in place, and resume it: the history it shares with that branch is kept as it is, and the branch's messages since are appended from sync. The branch it was on stays in atuin, for `--fork` to bring back, and the copy as it was is kept as a backup. Only when sync holds all of it and no agent here has it open")
	ai_resumeCmd.Flag("keymap-mode").Hidden = true
	ai_resumeCmd.Flag("shell-widget").Hidden = true
	aiCmd.AddCommand(ai_resumeCmd)

	carapace.Gen(ai_resumeCmd).FlagCompletion(carapace.ActionMap{
		"filter-mode": carapace.ActionValues("branch", "directory", "global", "host", "workspace"),
		"in":          carapace.ActionValues("amp", "claude-code", "codex", "copilot", "grok", "hermes", "kiro", "omp", "opencode", "pi"),
		"keymap-mode": carapace.ActionValues("auto", "emacs", "vim-insert", "vim-normal"),
	})
}
