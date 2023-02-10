# syntax=docker/dockerfile:1.3

ARG ROCKY9_VERSION=9.1.20221221

FROM docker.io/rockylinux/rockylinux:${ROCKY9_VERSION} as builder

FROM docker.io/rockylinux/rockylinux:${ROCKY9_VERSION}

# update - pacakges
RUN --mount=type=cache,target=/var/cache,id=build-cache \
    --mount=type=cache,target=/tmp,id=build-tmp \
    \
    dnf clean expire-cache \
    && dnf makecache --refresh \
    && dnf install -y epel-release* \
    && dnf makecache --refresh \
    && dnf update --skip-broken -y

# install - releated tools
RUN --mount=type=cache,target=/var/cache,id=build-cache \
    --mount=type=cache,target=/tmp,id=build-tmp \
    \
    dnf install -y git
