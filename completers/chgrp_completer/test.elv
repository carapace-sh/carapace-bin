edit:completion:arg-completer[chgrp] = [@arg]{
  fn _chgrp_callback [uid]{
    # TODO there is no 'eval' in elvish and '-source' needs a file so use a tempary one for callback 
    tmpfile=(mktemp -t carapace_chgrp_callback-XXXXX.elv)
    echo (joins ' ' $arg) | xargs chgrp_completer _carapace elvish $uid > $tmpfile
    -source $tmpfile
    rm $tmpfile
  }

  fn subindex [subcommand]{
    # TODO 'edit:complete-getopt' needs the arguments shortened for subcommmands - pretty optimistic here
    index=1
    for x $arg { if (eq $x $subcommand) { break } else { index = (+ $index 1) } } 
    echo $index
  }
  
  state=(echo (joins ' ' $arg) | xargs chgrp_completer _carapace elvish state)
  if (eq 1 0) {
  }  elif (eq $state '_chgrp') {
    opt-specs = [
        [&long='H' &desc='if a command line argument is a symbolic link to a directory, traverse it' &short='H']
        [&long='L' &desc='traverse every symbolic link to a directory encountered' &short='L']
        [&long='P' &desc='do not traverse any symbolic links (default)' &short='P']
        [&long='changes' &desc='like verbose but report only when a change is made' &short='c']
        [&long='dereference' &desc='affect the referent of each symbolic link (this is the default), rather than the symbolic link itself']
        [&long='help' &desc='display this help and exit']
        [&long='no-dereference' &desc='affect symbolic links instead of any referenced file (useful only on systems that can change the ownership of a symlink)' &short='h']
        [&long='no-preserve-root' &desc='do not treat \"/\" specially (the default)']
        [&long='preserve-root' &desc='fail to operate recursively on \"/\"']
        [&long='quiet' &desc='suppress most error messages']
        [&long='recursive' &desc='operate on files and directories recursively' &short='R']
        [&long='reference' &desc='use RFILE\"s group rather than specifying a GROUP value' &arg-required=$true &completer=[_]{  }]
        [&long='silent' &desc='suppress most error messages' &short='f']
        [&long='verbose' &desc='output a diagnostic for every file processed' &short='v']
        [&long='version' &desc='output version information and exit']
    ]
    arg-handlers = [
      [_]{  }
      [_]{ edit:complete-filename $arg[-1] }
      ...
    ]
    subargs = $arg[(subindex chgrp):] 
    if (> (count $subargs) 0) {
      edit:complete-getopt $subargs $opt-specs $arg-handlers
    }
  }
}

