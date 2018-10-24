FROM ubuntu:bionic
ENV PATH $PATH:/root/.cargo/bin
RUN apt-get update && \
    apt-get install -y curl gcc lib32ncurses5 lib32z1 lib32gcc1 libc6-dev-i386 && \
    curl https://sh.rustup.rs -sSf | sh -s -- -y && \
    rustup target add i686-unknown-linux-musl
