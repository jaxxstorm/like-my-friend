FROM alpine

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

RUN /usr/bin/curl -L https://github.com/jaxxstorm/like-my-friend/releases/download/v0.1.0/like-my-friend_0.1.0_linux_amd64.tar.gz -o /tmp/like-my-friend_0.1.0_linux_amd64.tar.gz

RUN tar zxvf /tmp/like-my-friend_0.1.0_linux_amd64.tar.gz

RUN mv like-my-friend /usr/local/bin/like-my-friend

ENTRYPOINT ["/usr/local/bin/like-my-friend"]
