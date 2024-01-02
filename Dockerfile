FROM golang:1.19

ARG app_name

RUN apt-get update && apt-get install -y wget

RUN	wget https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-arm64 \
    && chmod +x tailwindcss-linux-arm64 \
    && mv tailwindcss-linux-arm64 /usr/local/bin/tailwindcss

RUN mkdir -p /go/src/github.com/Blackmocca/wedding-ui
WORKDIR /go/src/github.com/Blackmocca/wedding-ui

ENV GO111MODULE=on
ENV ADDR=0.0.0.0
ENV TZ=Asia/Bangkok

# Copy app service 
RUN mkdir -p build/web/resources
COPY resources build/web

COPY go.mod .
COPY . .

RUN go mod tidy     

RUN cd tailwind && tailwindcss -i ./tailwind-min.css -o ../resources/styles/tailwind/tailwind-min.css --minify
RUN GOARCH=wasm GOOS=js go build -o ./build/web/app.wasm main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o  ./build/app main.go

FROM alpine:latest 
RUN apk --no-cache add ca-certificates
WORKDIR /usr/app

ARG app_name
ARG build_number
ARG version
ENV env_build_number=${build_number}
ENV env_version=${version}
ENV ADDR=0.0.0.0
ENV TZ=Asia/Bangkok

COPY --from=0 /go/src/github.com/Blackmocca/wedding-ui/build .

EXPOSE 3000
EXPOSE 3100

CMD ["./app"]  

