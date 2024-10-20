#!/bin/bash

function usage() {
    echo "Usage: $0 [options]"
    echo "Options:"
    echo "  -h, --help     Show this message"
    echo "  elminster    Use Elminster"
    echo "  tasha        Use Tasha"
    echo "  docker       Use Docker"
}

function usage_elminster() {
    echo "Usage: $0 elminster [options]"
    echo "Options:"
    echo "  -h, --help     Show this message"
}

function usage_docker() {
    echo "Usage: $0 docker [options]"
    echo "Options:"
    echo "  -h, --help     Show this message"
    echo "  down           down the containers"
    echo "  up             up the containers"
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

function handle_docker() {
    case $1 in
    down)
        docker compose -f docker-compose.databases.yaml -f docker-compose.queueservice.yaml down -d
        ;;
    up)
        docker compose -f docker-compose.databases.yaml -f docker-compose.queueservice.yaml up -d
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
    elminster)
        handle_elminster "$2"
        ;;
    docker)
        handle_docker "$2"
        ;;
    *)
        usage
        exit 1
        ;;
    esac
}

init "$@"
