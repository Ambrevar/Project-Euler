#!/bin/sh

SCRIPTPATH="$(readlink -f "$(dirname "$0")")"
find "$SCRIPTPATH" -mindepth 2 -type f -executable -print -delete
find "$SCRIPTPATH" -mindepth 2 -type f -name "*.o" -print -delete
