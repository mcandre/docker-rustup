# mcandre/docker-rustup: Docker images for (Linux) rustup build bots

# ABOUT

mcandre/docker-rustup Docker images help you compile Rust applications for different systems, without requiring Rust on your host system.

# EXAMPLE

```console
$ docker run -v "$(pwd):/src" mcandre/docker-rustup:x86_64-gnu sh -c "cd /src && cargo build --release"

$ file target/x86_64-unknown-linux-gnu/release/hello
...
```

# DOCKERHUB

https://hub.docker.com/r/mcandre/docker-rustup/

# REQUIREMENTS

* [Docker](https://www.docker.com)

## Optional

* UNIX [file](https://linux.die.net/man/1/file) utility
* [Mage](https://magefile.org)
* [Node.js](https://nodejs.org/en/) (for dockerfile_lint, dockerfilelint, dockerlint, dockerfile-utils)
* [hadolint](https://github.com/hadolint/hadolint)

# BUILD AND TEST IMAGES

```console
$ mage

$ docker images | grep mcandre/docker-rustup
mcandre/docker-rustup            i686-musl           db7e17ec244e        4 minutes ago       637 MB
mcandre/docker-rustup            x86_64-musl         d762f3a1c85b        38 minutes ago      586 MB
mcandre/docker-rustup            i686-gnu            650b68943a9a        About an hour ago   599 MB
mcandre/docker-rustup            x86_64-gnu          47f8035c6594        2 hours ago         543 MB
```

# PUBLISH IMAGES

```console
$ mage publish
```
