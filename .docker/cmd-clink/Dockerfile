FROM archlinux

RUN echo -e "\n[multilib]\nInclude = /etc/pacman.d/mirrorlist\n" >> /etc/pacman.conf

RUN pacman -Sy --noconfirm elvish unzip wine

RUN curl -L "https://github.com/chrisant996/clink/releases/download/v1.7.7/clink.1.7.7.521fa7.zip" > /tmp/clink.zip \
 && unzip -d /clink /tmp/clink.zip \
 && rm /tmp/clink.zip

ENV WINEPATH="Z:\\clink;Z:\\carapace-bin\\cmd\\carapace"

# TODO register clink/carapace and fix tab completion (https://bugs.winehq.org/show_bug.cgi?id=24026)

RUN mkdir -p ~/.config/elvish \
 && echo -e "set paths = [ /carapace-bin/cmd/carapace \$@paths ]\neval (carapace _carapace|slurp)" > ~/.config/elvish/rc.elv
ENV PATH="/carapace-bin/cmd/carapace:$PATH"

CMD ["wine", "cmd"]
