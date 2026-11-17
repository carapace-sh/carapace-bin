from xonsh.completers.completer import add_one_completer
from xonsh.completers.tools import contextual_command_completer
import os

# __xonsh__.env['PATH'].insert(0, '/tmp/carapace-snippet-fixtures/carapace/bin')

@contextual_command_completer
def _carapace_completer(context):
    """carapace completer"""
    from json import loads
    from xonsh.completers.tools import sub_proc_get_output, RichCompletion

    if context.command not in ['example']:
        return

    def fix_prefix(s):
        """quick fix for partially quoted prefix completion ('prefix',<TAB>)"""
        return s.translate(str.maketrans('', '', '\'"'))

    carapace_env = {'CARAPACE_SHELL': 'xonsh'}
    carapace_env.update({
        'CARAPACE_SHELL_ALIASES': '\n'.join(__xonsh__.aliases.keys()),
        'CARAPACE_SHELL_BUILTINS': '\n'.join(vars(__xonsh__.builtins).keys()),
        'CARAPACE_SHELL_FUNCTIONS': '\n'.join(k for k, v in __xonsh__.ctx.items() if callable(v)),
        'CARAPACE_SHELL_VARIABLES': '\n'.join(__xonsh__.env.keys()),
    })
    output, _ = sub_proc_get_output(
        'carapace', context.command, 'xonsh', *[a.value for a in context.args], fix_prefix(context.prefix),
        **carapace_env,
    )

    try:
        result = {RichCompletion(c["Value"], display=c["Display"], description=c["Description"], prefix_len=len(context.raw_prefix), append_closing_quote=False, style=c["Style"]) for c in loads(output)}
    except:
        result = {}
    if len(result) == 0:
        result = {RichCompletion(context.prefix, display=context.prefix, description='', prefix_len=len(context.raw_prefix), append_closing_quote=False)}
    return result

add_one_completer('carapace', _carapace_completer, 'start')

