package lazyinit

import (
	"fmt"
	"runtime"
	"strings"
)

func Powershell(completers []string) string {
	snippet := `using namespace System.Management.Automation
using namespace System.Management.Automation.Language

%v%v

$_carapace_completer = {
    [System.Diagnostics.CodeAnalysis.SuppressMessageAttribute("PSAvoidUsingInvokeExpression", "", Scope="Function", Target="*")]
    param($wordToComplete, $commandAst, $cursorPosition)
    $commandElements = $commandAst.CommandElements

    # double quoted value works but seems single quoted needs some fixing (e.g. "example 'acti" -> "example acti")
    $elems = @()
    foreach ($_ in $commandElements) {
      if ($_.Extent.StartOffset -gt $cursorPosition) {
          break
      }
      $t = $_.Extent.Text
      if ($_.Extent.EndOffset -gt $cursorPosition) {
          $t = $t.Substring(0, $_.Extent.Text.get_Length() - ($_.Extent.EndOffset - $cursorPosition))
      }

      if ($t.Substring(0,1) -eq "'"){
        $t = $t.Substring(1)
      }
      if ($t.get_Length() -gt 0 -and $t.Substring($t.get_Length()-1) -eq "'"){
        $t = $t.Substring(0,$t.get_Length()-1)
      }
      if ($t.get_Length() -eq 0){
        $t = '""'
      }
      $elems += $t.replace('` + "`" + `,', ',') # quick fix
    }

    $completions = @(
      if (!$wordToComplete) {
        carapace $elems[0] powershell $($elems| ForEach-Object {$_}) '' | ConvertFrom-Json | ForEach-Object { [CompletionResult]::new($_.CompletionText, $_.ListItemText.replace('` + "`" + `e[', "` + "`" + `e["), [CompletionResultType]::ParameterValue, $_.ToolTip.replace('` + "`" + `e[', "` + "`" + `e[")) }
      } else {
        carapace $elems[0] powershell $($elems| ForEach-Object {$_}) | ConvertFrom-Json | ForEach-Object { [CompletionResult]::new($_.CompletionText, $_.ListItemText.replace('` + "`" + `e[', "` + "`" + `e["), [CompletionResultType]::ParameterValue, $_.ToolTip.replace('` + "`" + `e[', "` + "`" + `e[")) }
      }
    )

    if ($completions.count -eq 0) {
      return "" # prevent default file completion
    }

    $completions
}

%v
`

	prefix := " # "
	if runtime.GOOS == "windows" {
		prefix = ""
	}

	complete := make([]string, 0, len(completers)*2)
	for _, completer := range completers {
		complete = append(complete, fmt.Sprintf(`Register-ArgumentCompleter -Native -ScriptBlock $_carapace_completer -CommandName '%v'%v,'%v.exe'`, completer, prefix, completer))
	}
	return fmt.Sprintf(snippet, pathSnippet("powershell"), envSnippet("powershell"), strings.Join(complete, "\n"))
}
