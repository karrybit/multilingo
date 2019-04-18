#!/bin/bash

export DEBUG=true

case $1 in
    "swift" )
        ONLY_EXECUTE=$(<./debug/swift/input_only_execute.txt)
        ONLY_EXECUTE_DOUBLE_SLASH=$(<./debug/swift/input_only_execute_double_slash.txt)
        WITH_ARGS=$(<./debug/swift/input_with_args.txt)
    ;;
    "python" )
        ONLY_EXECUTE=$(<./debug/python/input_only_execute.txt)
        WITH_ARGS=$(<./debug/python/input_with_args.txt) ;;
    * ) exit 1 ;;
esac

./app $1 "${ONLY_EXECUTE}"
./app $1 "${ONLY_EXECUTE_DOUBLE_SLASH}"
./app $1 "${WITH_ARGS}"
