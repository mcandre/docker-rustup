FROM ubuntu:bionic
LABEL Name=docker-rustup:i686-gnu \
    Version=0.0.1
ENV PATH=$PATH:/root/.cargo/bin
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        gcc-multilib \
        lib32ncurses5 \
        lib32z1 \
        lib32gcc1 \
        libc6-dev-i386 \
        curl \
        ca-certificates && \
    curl https://sh.rustup.rs -sSf | \
        sh -s -- -y --default-host i686-unknown-linux-gnu && \
    apt-get remove -y curl ca-certificates && \
    apt-get autoremove -y && \
    apt-get clean -y && \
    rm -rf /var/lib/apt/lists/*
