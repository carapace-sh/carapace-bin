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
    carapace_state = os.path.join(os.environ.get('XDG_RUNTIME_DIR') or __import__('tempfile').gettempdir(), 'carapace', 'builtins', str(os.getuid()), str(os.getpid()))
    os.makedirs(carapace_state, mode=0o700, exist_ok=True)
    with open(os.path.join(carapace_state, 'aliases'), 'w') as f: f.write('\n'.join(__xonsh__.aliases.keys()))
    with open(os.path.join(carapace_state, 'builtins'), 'w') as f: f.write('\n'.join(vars(__xonsh__.builtins).keys()))
    with open(os.path.join(carapace_state, 'functions'), 'w') as f: f.write('\n'.join(k for k, v in __xonsh__.ctx.items() if callable(v)))
    with open(os.path.join(carapace_state, 'variables'), 'w') as f: f.write('\n'.join(__xonsh__.env.keys()))
    carapace_env['CARAPACE_SHELL_STATE'] = carapace_state
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

