FROM node:12.14-alpine as uibuilder
COPY . .
RUN cd client && npm install
RUN cd client && node_modules/@angular/cli/bin/ng build --prod

FROM golang:1.13.5 as serverbuilder
COPY server .
RUN ls
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch

COPY --from=uibuilder /client  /server/client/dist
COPY --from=serverbuilder main  /server

WORKDIR  /server

EXPOSE 8091

CMD ["./main"]
