#!/bin/bash

function usage() {
    echo "Usage: $0 [opção]"
    echo "Options:"
    echo "  -h, --help     Show this message"
    echo "  --elminster    Use Elminster"
    echo "  --tasha        Use Tasha"
}

if [ $# -eq 0 ]; then
    usage
    exit 1
fi

case $1 in
    -h|--help)
        usage
        ;;
    *)
        usage
        exit 1
        ;;
esac