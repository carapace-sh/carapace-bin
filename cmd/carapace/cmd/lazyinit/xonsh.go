package lazyinit

import (
	"fmt"
	"strings"
)

func Xonsh(completers []string) string {
	snippet := `from xonsh.completers._aliases import _add_one_completer
from xonsh.completers.tools import *
import os

%v%v

@contextual_completer
def _carapace_lazy(context):
    """carapace lazy"""
    if (context.command and
        context.command.arg_index > 0 and
        context.command.args[0].value in [%v]):
        exec(compile(subprocess.run(['carapace', context.command.args[0].value, 'xonsh'], stdout=subprocess.PIPE).stdout.decode('utf-8'), "", "exec"))
        return XSH.completers[context.command.args[0].value](context)
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`'%v'`, completer)
	}
	snippet += `_add_one_completer('carapace_lazy', _carapace_lazy, 'start')`
	return fmt.Sprintf(snippet, pathSnippet("xonsh"), envSnippet("xonsh"), strings.Join(complete, ", "))
}
