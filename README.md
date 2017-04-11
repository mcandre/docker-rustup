# mcandre/docker-rustup: Docker images for (Linux) rustup build bots

# EXAMPLE

```console
$ docker pull mcandre/docker-rustup:x86_64-gnu

$ git clone https://github.com/mcandre/ios7crypt-rs.git
$ cd ios7crypt-rs

$ docker run -v "$(pwd):/src" mcandre/docker-rustup:x86_64-gnu sh -c "cd /src && cargo build --release"

$ file target/x86_64-unknown-linux-gnu/release/ios7crypt
target/x86_64-unknown-linux-gnu/release/ios7crypt: ELF 64-bit LSB shared object, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, for GNU/Linux 2.6.32, BuildID[sha1]=d1cb7423e44172aeb754c827084ffde9edcafe91, not stripped
```

# REQUIREMENTS

* [Docker](https://www.docker.com)

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)

# BUILD DOCKER IMAGES

```console
$ make

$ docker images | grep mcandre/docker-rustup
mcandre/docker-rustup            x86_64-gnu          47f8035c6594        11 minutes ago      543 MB
```

# PUBLISH DOCKER IMAGES

```console
$ make publish
```
