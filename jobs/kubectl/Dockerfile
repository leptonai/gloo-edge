FROM bitnami/kubectl:1.24.7 as kubectl

FROM alpine:3.17.3

RUN apk upgrade --update-cache

COPY --from=kubectl /opt/bitnami/kubectl/bin/kubectl /usr/local/bin/
