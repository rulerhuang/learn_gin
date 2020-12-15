FROM golang:alpine AS builder
WORKDIR /home
COPY . .
RUN mkdir -p /root/.ssh
COPY ./ssh  /root/.ssh
ENV GO111MODULE=on GOPRIVATE=code.clouderwork.com GOPROXY=https://goproxy.cn,direct \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN chmod 0600 /root/.ssh/id_rsa \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add git \
    && apk add openssh
RUN go mod download && go build -o learn_gin

FROM alpine AS server
LABEL maintainer huangchi@yunzujia.com
WORKDIR /home/gin
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && apk del tzdata
COPY --from=builder /home/learn_gin .
COPY --from=builder /home/conf ./conf
EXPOSE 8000
ENTRYPOINT [ "./learn_gin" ]