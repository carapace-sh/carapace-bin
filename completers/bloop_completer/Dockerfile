FROM alpine

RUN apk add --no-cache curl gcompat
RUN apk add --no-cache \
            --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community \
            elvish
RUN curl -fL "https://github.com/coursier/launchers/raw/master/cs-x86_64-pc-linux.gz" \
  | gzip -d > /usr/bin/cs \
 && chmod +x /usr/bin/cs
RUN cs install bloop --only-prebuilt=true

ENV PATH="$PATH:/root/.local/share/coursier/bin"
RUN bloop
