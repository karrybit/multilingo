#!/bin/bash

export DEBUG=true

case $1 in
    "swift" ) PROGRAM=$(<./debug/print.swift) ;;
    "python" ) PROGRAM=$(<./debug/print.py) ;;
    * ) exit 1 ;;
esac

./app "$PROGRAM" $1
