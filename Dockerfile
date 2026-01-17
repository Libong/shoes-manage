#第一级构建如果需要进行go build的话 就需要使用自带go环境的alpine源镜像 不过镜像大小很大 100M多
#FROM golang:1.20-alpine3.16 as builder
#作者
#MAINTAINER libong

#RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
#RUN set -ex \
#    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
#    && apk --update add tzdata \
#    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#    && apk --no-cache add ca-certificates

# 创建了一个app-runner的用户, -D表示无密码
#RUN adduser -u 10001 -D app-runner
# 设置环境变量
#ENV GOPROXY https://goproxy.cn
#ENV GOPRIVATE github.com/Libong
#ENV CGO_ENABLED 1
#ENV GO111MODULE on
#RUN git config --global url."https://libong:${{secrets.GO_MOD}}@github.com".insteadOf "https://github.com"
#RUN apk add --no-cache libc6-compat
# 把当前目录的文件拷过去，编译代码
#COPY ./main /app/
#WORKDIR /app
#RUN ls -l
#RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o main ./app/interface/oss/cmd
#RUN ls -l
#RUN go mod tidy
#RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a  -ldflags '-w -s' -o main .

#FROM alpine:3.16 AS final
# 下载时区包
#COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# 设置当前时区
#COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# https ssl证书
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 使用app-runner启动
#USER app-runner
#ENTRYPOINT ["/app/blockchain-middleware"]

# 定义容器运行时的命令
#CMD ["/app/main"]

#第二级构建（不需要进行编译）可以使用原始的alpine镜像 大小会小一点
FROM alpine AS runner
#作者
MAINTAINER libong
#修改Alpine Linux的APK源，将其从默认的dl-cdn.alpinelinux.org更换为阿里云的镜像源mirrors.aliyun.com 加速后续下载其他软件包的速度
RUN set -ex \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
#将编译好的执行文件拷贝到docker容器指定位置
COPY ./main /app/
#设置工作目录
WORKDIR /app
#设置环境变量
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore
# 暴露服务端口
#EXPOSE 8080
# 定义容器运行时的命令
ENTRYPOINT ["./main"]
CMD ["-configPath","/etc/myconfig","-env","prod","-log","./log"]