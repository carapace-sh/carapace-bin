package lazyinit

import (
	"fmt"
	"strings"
)

func Fish(completers []string) string {
	snippet := `%v%v

function _carapace_completer
  set --local --export CARAPACE_SHELL fish
  set --local --export CARAPACE_SHELL_BUILTINS (builtin --names)
  set --local --export CARAPACE_SHELL_FUNCTIONS (functions --all)
  set --local data
  IFS='' set data (echo (commandline -cp)'' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
  if [ $status -eq 1 ]
    IFS='' set data (echo (commandline -cp)"'" | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
    if [ $status -eq 1 ]
      IFS='' set data (echo (commandline -cp)'"' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
    end
  end
  echo $data
end

%v
`
	complete := make([]string, 0, len(completers)*2)
	for _, completer := range completers {
		complete = append(complete,
			fmt.Sprintf(`complete -e %q`, completer),
			fmt.Sprintf(`complete -c %q -f -a '(_carapace_completer %q)'`, completer, completer),
		)
	}
	return fmt.Sprintf(snippet, pathSnippet("fish"), envSnippet("fish"), strings.Join(complete, "\n"))
}
