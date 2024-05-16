#!/bin/bash

GO_VERSION="1.21.9"
RUST_VERSION="1.72.0"
SKOPEO_VERSION="1.5.0"
YQ_VERSION="v4.35.1"

# Install dependencies
yum install -y curl protobuf-compiler libseccomp-devel openssl openssl-devel perl skopeo-${SKOPEO_VERSION}

# Install Golang
curl https://dl.google.com/go/go${GO_VERSION}.linux-ppc64le.tar.gz -o go${GO_VERSION}.linux-ppc64le.tar.gz && \
rm -rf /usr/local/go && tar -C /usr/local -xzf go${GO_VERSION}.linux-ppc64le.tar.gz && \
rm -f go${GO_VERSION}.linux-ppc64le.tar.gz

# Install Rust
curl https://sh.rustup.rs -sSf | sh -s -- -y --default-toolchain ${RUST_VERSION}
rustup target add powerpc64le-unknown-linux-gnu

curl https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_ppc64le -o /usr/local/bin/yq
chmod a+x /usr/local/bin/yq