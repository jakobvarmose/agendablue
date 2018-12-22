FROM node:latest as frontend
WORKDIR /root/
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

FROM golang:latest as backend
WORKDIR /go/src/github.com/jakobvarmose/agendablue/backend/
RUN go get \
    github.com/go-sql-driver/mysql \
    github.com/jinzhu/gorm \
    github.com/whyrusleeping/cbor/go \
    golang.org/x/crypto/nacl/sign
COPY backend/ ./
RUN CGO_ENABLED=0 go build -o /root/app

FROM alpine:latest
WORKDIR /root/
COPY --from=frontend /root/dist/ dist/
COPY --from=backend /root/app app
CMD ["./app"]
EXPOSE 8041
