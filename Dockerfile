FROM node:12.14-alpine as uibuilder
WORKDIR /client
COPY client .
RUN npm install
RUN node_modules/@angular/cli/bin/ng build --prod

FROM golang:1.14 as serverbuilder
WORKDIR /servers
COPY servers .
WORKDIR /servers/server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:3.13 as certs
RUN apk --no-cache add ca-certificates=20191127-r4

FROM scratch

COPY --from=uibuilder /client/dist  /server/client/dist
COPY --from=serverbuilder /servers/server/main  /server/main
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR  /server

CMD ["./main"]
