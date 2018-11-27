FROM ubuntu:bionic
LABEL Name=docker-rustup:x86_64-musl \
    Version=0.0.1
ENV PATH=$PATH:/root/.cargo/bin
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        gcc \
        libc-dev \
        curl \
        ca-certificates && \
    curl https://sh.rustup.rs -sSf | \
        sh -s -- -y && \
    rustup target add x86_64-unknown-linux-musl && \
    apt-get remove -y curl ca-certificates && \
    apt-get autoremove -y && \
    apt-get clean -y && \
    rm -rf /var/lib/apt/lists/*
