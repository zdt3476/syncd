FROM golang:1.18.3-alpine3.16 AS build
COPY . /usr/local/src
WORKDIR /usr/local/src
RUN apk --no-cache add build-base && make

FROM alpine:3.16
WORKDIR /syncd
COPY --from=build /usr/local/src/output /syncd
RUN apk add --no-cache bash
EXPOSE 8878
CMD [ "/syncd/bin/syncd" ]
