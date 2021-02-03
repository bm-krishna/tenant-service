FROM golang:alpine
COPY . $GOPATH/src/github.com/bm-krishna/tenant
WORKDIR $GOPATH/src/github.com/bm-krishna/tenant
RUN go get -d -v ./...
RUN go install  -v ./...
RUN go build -o ./cmd/server ./cmd/server/main.go 
EXPOSE 7171
CMD [ "cmd/server/main" ]