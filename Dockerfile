FROM node:10.16-alpine as uibuilder
COPY . .
RUN cd ridemanager-client && npm install
RUN cd ridemanager-client && node_modules/@angular/cli/bin/ng build --prod

FROM golang:1.12-alpine as serverbuilder
WORKDIR /ridemanager
COPY . .
WORKDIR /ridemanager
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch

COPY --from=uibuilder /ridemanager-client/dist  /server/ridemanager-client/dist
COPY --from=serverbuilder /ridemanager/main  /server

WORKDIR  /server

EXPOSE 8091

CMD ["./main"]
