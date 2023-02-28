FROM alpine:3.17.0

RUN apk add --no-cache \
	go \
	git \
	make

RUN go install github.com/cweill/gotests/gotests@v1.6.0 \
	&& go install github.com/fatih/gomodifytags@v1.16.0 \
	&& go install github.com/josharian/impl@v1.1.0 \
	&& go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0 \
	&& go install github.com/go-delve/delve/cmd/dlv@latest \
	&& go install honnef.co/go/tools/cmd/staticcheck@latest \
	&& go install golang.org/x/tools/gopls@latest