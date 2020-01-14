FROM golang

RUN dpkg --add-architecture armhf

RUN apt update && \
    apt install -y build-essential crossbuild-essential-arm64 \
    crossbuild-essential-armel \
    crossbuild-essential-armhf \
    linux-libc-dev-arm64-cross \
    linux-libc-dev-armel-cross \
    linux-libc-dev-armhf-cross \
    gcc-arm-linux-gnueabihf g++-arm-linux-gnueabihf \
    gcc-aarch64-linux-gnu g++-aarch64-linux-gnu \
 && \
    rm -rf /var/lib/apt/lists/*