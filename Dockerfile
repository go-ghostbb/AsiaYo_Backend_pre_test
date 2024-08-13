FROM    golang:alpine AS stage1
ENV     RUN_PATH=/AsiaYo
RUN     mkdir -p $RUN_PATH
WORKDIR $RUN_PATH
ENV     GO111MODULE=on
# 複製所有檔案到/AsiaYo
COPY .. .
# 下載make
RUN apk update && apk add make
# 編譯
RUN make build

FROM alpine
USER root
# 主程式操作
RUN mkdir /AsiaYo
WORKDIR /AsiaYo
# 從stage1編譯好的二進制檔案複製到/AsiaYo裡
COPY   --from=stage1 /AsiaYo/bin .
# expose port
EXPOSE 8080
ENTRYPOINT ["./test"]