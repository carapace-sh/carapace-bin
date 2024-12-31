FROM archlinux

RUN pacman -Sy --noconfirm fzf git zsh

RUN git clone --depth 1 https://github.com/Aloxaf/fzf-tab /root/fzf-tab
ADD .zshrc /root/.zshrc
ENV PATH="/carapace-bin/cmd/carapace:$PATH"

CMD ["zsh"]
