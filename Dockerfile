FROM golang:1.12.0 AS builder

WORKDIR /build
COPY . ./
#COPY ./jocker /jocker
RUN CGO_ENABLED=0 go build -o /sample ./cmd/

FROM scratch
COPY --from=builder /sample /bin/sample
ENTRYPOINT ["/bin/sample"]
