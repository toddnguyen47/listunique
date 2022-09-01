#!/bin/bash

# To be used with Goland's file watcher

PKG_CONFIG_PATH="/opt/homebrew/opt/openssl@3/lib/pkgconfig"

#echo $1
#staticcheck $1
staticcheck -tags dynamic "${1}/..."
