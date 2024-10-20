#!/bin/bash

function usage() {
    echo "Usage: $0 [options]"
    echo "Options:"
    echo "  -h, --help     Show this message"
    echo "  --elminster    Use Elminster"
    echo "  --tasha        Use Tasha"
}

function usage_elminster() {
    echo "Usage: $0 --elminster [options]"
    echo "Optiosn:"
    echo "  -h, --help     Show this message"
}

function handle_elminster() {
    case $1 in
    -h|--help)
        usage_elminster
        ;;
    *)
        usage_elminster
        exit 1
        ;;
    esac
}

# function usage_tasha() {

# }

# function usage_owlbear() {

# }

# function usage_bigby() {

# }

function init() {
    if [ $# -eq 0 ]; then
        usage
        exit 1
    fi

    case $1 in
    -h|--help)
        usage
        ;;
    --elminster)
        handle_elminster "$2"
        ;;
    *)
        usage
        exit 1
        ;;
    esac
}

init "$@"
