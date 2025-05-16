FROM homebrew/brew

RUN brew install elvish

RUN mkdir -p ~/.config/elvish \
 && echo "set paths = [ /carapace-bin/cmd/carapace \$@paths ]\neval (carapace _carapace|slurp)" > ~/.config/elvish/rc.elv
ENV PATH="/carapace-bin/cmd/carapace:$PATH"

ENTRYPOINT [ "elvish" ]
