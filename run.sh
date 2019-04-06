#!/bin/bash

export DEBUG=true

case $1 in
    "swift" )
        ONLY_PROGRAM=$(<./debug/OnlyExecute.swift)
        INPUT=$(<./debug/Input.txt)
        PROGRAM=$(<./debug/Program.swift)
    ;;
    "python" ) PROGRAM=$(<./debug/print.py) ;;
    * ) exit 1 ;;
esac

./app $1 "<@$1>\\n\`\`\`${ONLY_PROGRAM}\`\`\`"
./app $1 "<@$1>\\n\`\`\`${INPUT}\`\`\`\\n\`\`\`${PROGRAM}\`\`\`"
