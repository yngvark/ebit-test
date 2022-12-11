FROM golang:1.19 as builder

RUN mkdir /build
COPY main.go go.mod internal /build/
#COPY internal /build

RUN cd /build && \
    go mod tidy && \
    env GOOS=js GOARCH=wasm go build -o /build/app.wasm ./main.go

FROM nginx

COPY web/mime.types etc/nginx/mime.types
COPY web/index.html web/wasm_exec.js    /usr/share/nginx/html/

COPY --from=builder /build/app.wasm     /usr/share/nginx/html/app.wasm
