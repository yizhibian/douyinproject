
# 基础镜像
FROM golang:1.19
#FROM golang:1.19-alpine

# 维护者
MAINTAINER CSN
# docker build 时执行命令 - 创建目录
RUN mkdir -p "/data/api" \
&& ln -sf /dev/stdout /data/commentapp/apiservice.log
# 工作目录
WORKDIR "/data/api"
# 拷贝
COPY api  /data/commentapp/api
RUN chmod a+x api
# docker run 时执行命令
ENTRYPOINT ["./api"]

