package lazyinit

import (
	"fmt"
	"strings"
)

func Fish(completers []string) string {
	snippet := `%v%v

function _carapace_quote_suffix
  if not commandline -cp | xargs echo 2>/dev/null >/dev/null
    if commandline -cp | sed 's/$/"/'| xargs echo 2>/dev/null >/dev/null
      echo '"'
    else if commandline -cp | sed "s/\$/'/"| xargs echo 2>/dev/null >/dev/null
      echo "'"
    end
  else
    echo ""
  end
end

function _carapace_completer
  commandline -cp | sed "s/\$/"(_carapace_quote_suffix)"/" | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish
end

%v
`
	complete := make([]string, 0, len(completers)*2)
	for _, completer := range completers {
		complete = append(complete,
			fmt.Sprintf(`complete -e '%v'`, completer),
			fmt.Sprintf(`complete -c '%v' -f -a '(_carapace_completer %v)'`, completer, completer),
		)
	}
	return fmt.Sprintf(snippet, pathSnippet("fish"), envSnippet("fish"), strings.Join(complete, "\n"))
}
