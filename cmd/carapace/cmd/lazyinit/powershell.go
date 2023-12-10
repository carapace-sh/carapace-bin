package lazyinit

import (
	"fmt"
	"strings"
)

func Powershell(completers []string) string {
	snippet := `%v%v

$_carapace_lazy = {
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
	return fmt.Sprintf(snippet, pathSnippet("powershell"), envSnippet("powershell"), strings.Join(complete, "\n"))
}
