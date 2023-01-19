FROM golang:1.18 as builder

COPY . /rosen
WORKDIR /rosen/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzRosen.a && \
    clang -fsanitize=fuzzer fuzzRosen.a -o fuzzRosen.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /rosen/mayhem/fuzzRosen.libfuzzer /