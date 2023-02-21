#!/bin/sh

architecture=$(arch)

if [ "${architecture}" = 'arm64' ]; then
    ./tt-arm64 "$@"
else
    ./tt-amd64 "$@"
fi
