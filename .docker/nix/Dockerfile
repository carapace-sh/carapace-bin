FROM archlinux

RUN pacman -Sy --noconfirm nix
RUN pacman -Sy --noconfirm elvish

RUN mkdir -p ~/.config/elvish \
 && echo -e "set paths = [ /carapace-bin/cmd/carapace \$@paths ]\neval (carapace _carapace|slurp)" > ~/.config/elvish/rc.elv
ENV PATH="/carapace-bin/cmd/carapace:$PATH"

RUN mkdir -p ~/.config/nix \
&&  echo -e "experimental-features = nix-command flakes" > ~/.config/nix/nix.conf

CMD ["elvish"]
