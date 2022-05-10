package cmd

//go:generate go run ../../generate/gen.go

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getSpecs() (specs []string, dir string) {
	r := regexp.MustCompile(`^[0-9a-zA-Z_\-.]+\.yaml$`) // sanity check
	specs = make([]string, 0)
	if configDir, err := os.UserConfigDir(); err == nil {
		dir = configDir + "/carapace/specs"
		if entries, err := os.ReadDir(dir); err == nil {
			for _, entry := range entries {
				if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".yaml") && r.MatchString(entry.Name()) {
					specs = append(specs, strings.TrimSuffix(entry.Name(), ".yaml"))
				}
			}
		} else if os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}
	return
}

func bash_lazy(completers []string) string {
	snippet := `_carapace_lazy() {
  source <(carapace $1 bash)
   $"_$1_completion"
}
complete -F _carapace_lazy %v
`

	if specs, dir := getSpecs(); len(specs) > 0 {
		snippet += fmt.Sprintf(`
_carapace_lazy_spec() {
  source <(carapace --spec "%v/$1.yaml" bash)
   $"_$1_completion"
}
complete -F _carapace_lazy_spec %v
`, dir, strings.Join(specs, " "))
	}

	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}

func bash_ble_lazy(completers []string) string {
	snippet := `_carapace_lazy() {
  source <(carapace $1 bash-ble)
   $"_$1_completion"
}
complete -F _carapace_lazy %v
`

	if specs, dir := getSpecs(); len(specs) > 0 {
		snippet += fmt.Sprintf(`
_carapace_lazy_spec() {
  source <(carapace --spec %v/$1.yaml bash-ble)
   $"_$1_completion"
}
complete -F _carapace_lazy_spec %v
`, dir, strings.Join(specs, " "))
	}

	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}

func elvish_lazy(completers []string) string {
	snippet := `put %v | each {|c|
    set edit:completion:arg-completer[$c] = {|@arg|
        set edit:completion:arg-completer[$c] = {|@arg| }
        eval (carapace $c elvish | slurp)
        $edit:completion:arg-completer[$c] $@arg
    }
}
`
	if specs, dir := getSpecs(); len(specs) > 0 {
		snippet += fmt.Sprintf(`
put %v | each {|c|
    set edit:completion:arg-completer[$c] = {|@arg|
        set edit:completion:arg-completer[$c] = {|@arg| }
        eval (carapace --spec "%v/"$c".yaml" elvish | slurp)
        $edit:completion:arg-completer[$c] $@arg
    }
  
}
`, strings.Join(specs, " "), dir)

	}

	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}

func fish_lazy(completers []string) string {
	snippet := `function _carapace_lazy
   complete -c $argv[1] -e
   carapace $argv[1] fish | source
   complete --do-complete=(commandline -cp)
end
%v
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`complete -c '%v' -f -a '(_carapace_lazy %v)'`, completer, completer)
	}

	if specs, dir := getSpecs(); len(specs) > 0 {
		snippet += fmt.Sprintf(`
function _carapace_lazy_spec
   complete -c $argv[1] -e
   carapace --spec "%v/$argv[1].yaml" fish | source
   complete --do-complete=(commandline -cp)
end

`, dir)
		for _, spec := range specs {
			snippet += fmt.Sprintf(`complete -c '%v' -f -a '(_carapace_lazy_spec %v)'
`, spec, spec)
		}
	}

	return fmt.Sprintf(snippet, strings.Join(complete, "\n"))
}

func nushell_lazy(completers []string) string {
	exports := make([]string, len(completers))
	for index, completer := range completers {
		exports[index] = fmt.Sprintf(`  export extern "%v" [
    ...args: string@"nu-complete carapace"
  ]`, completer)
	}
	snippet := fmt.Sprintf(`module carapace {
  def "nu-complete carapace" [line: string, pos: int] {
    let words = ($line | str substring [0 $pos] | split row " ")
    if ($line | str substring [0 $pos] | str ends-with " ") {
      carapace $words.0 nushell ($words | append "") | from json
    } else {
      carapace $words.0 nushell $words | from json
    }
  }

%v
`, strings.Join(exports, "\n"))

	//	if specs, dir := getSpecs(); len(specs) > 0 {
	//		snippet += fmt.Sprintf(`
	//  def "nu-complete carapace-spec" [line: string, pos: int] {
	//    let words = ($line | str substring [0 $pos] | split row " ")
	//    if ($line | str substring [0 $pos] | str ends-with " ") {
	//      carapace --spec '%v'/$words.0'.yaml' nushell ($words | append "") | from json
	//    } else {
	//      carapace --spec '%v'/$words.0'.yaml' nushell $words | from json
	//    }
	//  }
	//`, dir, dir)
	//
	//		for _, spec := range specs {
	//			snippet += fmt.Sprintf(`  export extern "%v" [
	//    ...args: string@"nu-complete carapace-spec"
	//  ]`, spec)
	//		}
	//	}

	snippet += `
}
use carapace *
`
	return snippet
}

func oil_lazy(completers []string) string {
	snippet := `_carapace_lazy() {
  source <(carapace $1 oil)
   $"_$1_completion"
}
complete -F _carapace_lazy %v
`

	if specs, dir := getSpecs(); len(specs) > 0 {
		snippet += fmt.Sprintf(`
_carapace_lazy_spec() {
  source <(carapace --spec "%v/$1.yaml" oil)
   $"_$1_completion"
}
complete -F _carapace_lazy_spec %v
`, dir, strings.Join(specs, " "))
	}

	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}

func powershell_lazy(completers []string) string {
	snippet := `$_carapace_lazy = {
    param($wordToComplete, $commandAst, $cursorPosition)
    $completer = $commandAst.CommandElements[0].Value
    carapace $completer powershell | Out-String | Invoke-Expression
    & (Get-Item "Function:_${completer}_completer") $wordToComplete $commandAst $cursorPosition
}
%v
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`Register-ArgumentCompleter -Native -CommandName '%v' -ScriptBlock $_carapace_lazy`, completer)
	}

	if specs, dir := getSpecs(); len(specs) > 0 {
		snippet += fmt.Sprintf(`$_carapace_lazy_spec = {
    param($wordToComplete, $commandAst, $cursorPosition)
    $completer = $commandAst.CommandElements[0].Value
    carapace --spec "%v/$completer.yaml" powershell | Out-String | Invoke-Expression
    & (Get-Item "Function:_${completer}_completer") $wordToComplete $commandAst $cursorPosition
}
`, dir)
		for _, spec := range specs {
			snippet += fmt.Sprintf(`Register-ArgumentCompleter -Native -CommandName '%v' -ScriptBlock $_carapace_lazy_spec
`, spec)
		}

	}

	return fmt.Sprintf(snippet, strings.Join(complete, "\n"))
}

func tcsh_lazy(completers []string) string {
	// TODO hardcoded for now
	snippet := make([]string, len(completers))
	for index, c := range completers {
		snippet[index] = fmt.Sprintf("complete \"%v\" 'p@*@`echo \"$COMMAND_LINE'\"''\"'\" | xargs carapace %v tcsh `@@' ;", c, c)
	}

	if specs, dir := getSpecs(); len(specs) > 0 {
		for _, spec := range specs {
			snippet = append(snippet, fmt.Sprintf("complete \"%v\" 'p@*@`echo \"$COMMAND_LINE'\"''\"'\" | xargs carapace --spec \"%v/%v.yaml\" tcsh `@@' ;", spec, dir, spec))
		}
	}
	return strings.Join(snippet, "\n")
}

func xonsh_lazy(completers []string) string {
	snippet := `from xonsh.completers._aliases import _add_one_completer
from xonsh.completers.tools import *

@contextual_completer
def _carapace_lazy(context):
    """carapace lazy"""
    if (context.command and
        context.command.arg_index > 0 and
        context.command.args[0].value in [%v]):
        builtins.__xonsh__.completers = builtins.__xonsh__.completers.copy()
        exec(compile(subprocess.run(['carapace', context.command.args[0].value, 'xonsh'], stdout=subprocess.PIPE).stdout.decode('utf-8'), "", "exec"))
        return builtins.__xonsh__.completers[context.command.args[0].value](context)
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`'%v'`, completer)
	}

	if specs, dir := getSpecs(); len(specs) > 0 {
		quotedSpecs := make([]string, 0)
		for _, spec := range specs {
			quotedSpecs = append(quotedSpecs, fmt.Sprintf(`'%v'`, spec))
		}
		snippet += fmt.Sprintf(`
    if (context.command and
        context.command.arg_index > 0 and
        context.command.args[0].value in [%v]):
        builtins.__xonsh__.completers = builtins.__xonsh__.completers.copy()
        exec(compile(subprocess.run(['carapace', '--spec', '%v/'+context.command.args[0].value+'.yaml', 'xonsh'], stdout=subprocess.PIPE).stdout.decode('utf-8'), "", "exec"))
        return builtins.__xonsh__.completers[context.command.args[0].value](context)
`, strings.Join(quotedSpecs, ", "), dir)

	}
	snippet += `_add_one_completer('carapace_lazy', _carapace_lazy, 'start')`
	return fmt.Sprintf(snippet, strings.Join(complete, ", "))
}

func zsh_lazy(completers []string) string {
	snippet := `function _carapace_lazy {
    source <(carapace $words[1] zsh)
}
compdef _carapace_lazy %v
`

	if specs, dir := getSpecs(); len(specs) > 0 {
		snippet += fmt.Sprintf(`function _carapace_lazy_spec {
    source <(carapace --spec "%v/$words[1].yaml" zsh)
}
compdef _carapace_lazy_spec %v
`, dir, strings.Join(specs, " "))

	}

	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}
