package lazyinit

import (
	"fmt"
	"strings"
)

// Xonsh creates a generic completer (not lazy).
// Name overlaps with specific `carapace` completer but that shouldn't be an issue.
func Xonsh(completers []string) string {
	snippet := `from xonsh.completers.completer import add_one_completer
from xonsh.completers.tools import contextual_command_completer
import os

%v%v

@contextual_command_completer
def _carapace_completer(context):
    """carapace completer"""
    from json import loads
    from xonsh.completers.tools import sub_proc_get_output, RichCompletion

    if context.command not in [%v]:
        return

    def fix_prefix(s):
        """quick fix for partially quoted prefix completion ('prefix',<TAB>)"""
        return s.translate(str.maketrans('', '', '\'"'))

    output, _ = sub_proc_get_output(
        'carapace', context.command, 'xonsh', *[a.value for a in context.args], fix_prefix(context.prefix)
    )

    try:
        result = {RichCompletion(c["Value"], display=c["Display"], description=c["Description"], prefix_len=len(context.raw_prefix), append_closing_quote=False, style=c["Style"]) for c in loads(output)}
    except:
        result = {}
    if len(result) == 0:
        result = {RichCompletion(context.prefix, display=context.prefix, description='', prefix_len=len(context.raw_prefix), append_closing_quote=False)}
    return result

add_one_completer('carapace', _carapace_completer, 'start')
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`'%v'`, completer)
	}
	return fmt.Sprintf(snippet, pathSnippet("xonsh"), envSnippet("xonsh"), strings.Join(complete, ", "))
}
