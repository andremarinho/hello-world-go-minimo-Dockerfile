FROM golang:alpine AS builder

WORKDIR /andre/src

COPY . .

RUN go build -ldflags '-s -w' hello.go

FROM scratch

WORKDIR /

COPY --from=builder /andre/src / 

CMD ["./hello"]

