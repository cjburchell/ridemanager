FROM node:12.14-alpine as uibuilder
WORKDIR /client
COPY client .
RUN npm install
RUN node_modules/@angular/cli/bin/ng build --prod

FROM golang:1.14 as serverbuilder
WORKDIR /servers
COPY servers .
WORKDIR /servers/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch

COPY --from=uibuilder /client  /server/client/dist
COPY --from=serverbuilder /servers/api/main  /server/main

WORKDIR  /server

CMD ["./main"]
