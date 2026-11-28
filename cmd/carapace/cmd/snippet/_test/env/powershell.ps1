using namespace System.Management.Automation
using namespace System.Management.Automation.Language

# [Environment]::SetEnvironmentVariable("PATH", '/tmp/carapace-snippet-fixtures/carapace/bin' + [IO.Path]::PathSeparator + [Environment]::GetEnvironmentVariable("PATH"))

Function get-env([string]$name) { Get-Item "env:$name" }
Function set-env([string]$name, [string]$value) { Set-Item "env:$name" "$value" }
Function unset-env([string]$name) { Remove-Item "env:$name" }

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
      $elems += $t.replace('`,', ',') # quick fix
    }

    $completions = @(
      $old = @{
        CARAPACE_SHELL = $env:CARAPACE_SHELL
        CARAPACE_SHELL_STATE = $env:CARAPACE_SHELL_STATE
      }
      try {
        $env:CARAPACE_SHELL = "powershell"

        if (!$wordToComplete) {
          carapace ($elems[0] -replace ('\.exe$', '')) powershell $($elems| ForEach-Object {$_}) '' | ConvertFrom-Json | ForEach-Object { [CompletionResult]::new($_.CompletionText, $_.ListItemText.replace('`e[', "`e["), [CompletionResultType]::ParameterValue, $_.ToolTip.replace('`e[', "`e[")) }
        } else {
          carapace ($elems[0] -replace ('\.exe$', '')) powershell $($elems| ForEach-Object {$_}) | ConvertFrom-Json | ForEach-Object { [CompletionResult]::new($_.CompletionText, $_.ListItemText.replace('`e[', "`e["), [CompletionResultType]::ParameterValue, $_.ToolTip.replace('`e[', "`e[")) }
        }
      } finally {
        foreach ($kv in $old.GetEnumerator()) {
          if ($null -eq $kv.Value) {
            Remove-Item "env:$($kv.Key)" -ErrorAction SilentlyContinue
          } else {
            Set-Item "env:$($kv.Key)" $kv.Value
          }
        }
      }
    )

    if ($completions.count -eq 0) {
      return "" # prevent default file completion
    }

    $completions
}

Register-ArgumentCompleter -Native -ScriptBlock $_carapace_completer -CommandName example # ,example.exe

