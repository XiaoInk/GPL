FROM alpine:latest

RUN apk add --no-cache tzdata

ENV TZ=Asia/Shanghai

WORKDIR /opt/work/bin

COPY bin/main .

EXPOSE 8080

ENTRYPOINT ["./main"]