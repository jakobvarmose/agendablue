FROM node:latest as frontend
WORKDIR /root/
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

FROM golang:latest as backend
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/jakobvarmose/agendablue/backend/
COPY backend/Gopkg.* ./
RUN dep ensure --vendor-only
COPY backend/ ./
RUN CGO_ENABLED=0 go build -o /root/app

FROM alpine:latest
WORKDIR /root/
COPY --from=frontend /root/dist/ dist/
COPY --from=backend /root/app app
CMD ["./app"]
EXPOSE 8041
