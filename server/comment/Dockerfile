
# 基础镜像
FROM golang:1.19
#FROM golang:1.19-alpine

# 维护者
MAINTAINER CSN
# docker build 时执行命令 - 创建目录
RUN mkdir -p "/data/app" \
&& ln -sf /dev/stdout /data/app/commentservice.log
# 工作目录
WORKDIR "/data/app"
# 拷贝
COPY comment  /data/app/comment
RUN chmod a+x comment
# docker run 时执行命令
ENTRYPOINT ["./comment"]

