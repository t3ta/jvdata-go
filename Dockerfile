FROM golang:1.19.0
RUN apt update && apt install git
RUN curl -LO https://github.com/kaitai-io/kaitai_struct_compiler/releases/download/0.10/kaitai-struct-compiler_0.10_all.deb
RUN apt install -y ./kaitai-struct-compiler_0.10_all.deb
WORKDIR /go/src/app
ADD . /go/src/app