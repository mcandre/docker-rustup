.PHONY: mcandre/docker-rustup-x86_64-gnu mcandre/docker-rustup-i686-gnu mcandre/docker-rustup-x86_64-musl

all: mcandre/docker-rustup-x86_64-gnu mcandre/docker-rustup-i686-gnu mcandre/docker-rustup-x86_64-musl

mcandre/docker-rustup-x86_64-gnu: x86_64-gnu/Dockerfile
	sh -c "cd x86_64-gnu && docker build -t mcandre/docker-rustup:x86_64-gnu ."

mcandre/docker-rustup-i686-gnu: i686-gnu/Dockerfile
	sh -c "cd i686-gnu && docker build -t mcandre/docker-rustup:i686-gnu ."

mcandre/docker-rustup-x86_64-musl: x86_64-musl/Dockerfile
	sh -c "cd x86_64-musl && docker build -t mcandre/docker-rustup:x86_64-musl ."

publish-x86_64-gnu: mcandre/docker-rustup-x86_64-gnu
	docker push mcandre/docker-rustup:x86_64-gnu

publish-i686-gnu: mcandre/docker-rustup-i686-gnu
	docker push mcandre/docker-rustup:i686-gnu

publish-x86_64-musl: mcandre/docker-rustup-x86_64-musl
	docker push mcandre/docker-rustup:x86_64-musl

publish: publish-x86_64-gnu publish-i686-gnu public-x86_64-musl
