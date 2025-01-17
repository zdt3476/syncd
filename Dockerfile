FROM golang:1.18.3-alpine3.16 AS build
COPY . /usr/local/src
WORKDIR /usr/local/src
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && apk --no-cache add build-base && make

FROM alpine:3.16
WORKDIR /syncd
COPY --from=build /usr/local/src/output /syncd
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && apk add --no-cache git && apk add --no-cache bash && apk add --no-cache openssh
EXPOSE 8878
CMD [ "/syncd/bin/syncd" ]
