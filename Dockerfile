FROM golang:latest as build

WORKDIR /go/src/parser

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o parserServer main.go


FROM alpine

EXPOSE 8085
EXPOSE 8084

ENV TZ Europe/Moscow

RUN apk update && apk add --no-cache tzdata \
&& cp /usr/share/zoneinfo/$TZ /etc/localtime \
&& echo $TZ > /etc/timezone \
&& apk del tzdata \
&& mkdir -p /usr/share/zoneinfo/Europe \
&& cp /etc/localtime /usr/share/zoneinfo/$TZ

WORKDIR /parser

COPY --from=build /go/src/parser/parserServer ./parserServer
COPY --from=build /go/src/parser/docs/* ./docs/

CMD ["./parserServer"]

