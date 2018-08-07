FROM alpine:3.4
ADD ./promcli /promcli
ENTRYPOINT /promcli
