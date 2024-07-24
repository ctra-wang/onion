FROM alpine:3.19

# Define the project name | 定义项目名称
ARG PROJECT=onion
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=onion.yaml
# Define the author | 定义作者
ARG AUTHOR="lei.wang@example.com"

RUN apk update && apk add tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime


COPY ./app /app/
COPY ./etc/onion.yaml /app/etc/config.yaml
workdir /app/
EXPOSE 2165
ENTRYPOINT ["./app"]