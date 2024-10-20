#!/bin/bash

function usage() {
    echo "Uso: $0 [opção]"
    echo "Opções:"
    echo "  -h, --help     Mostra esta ajuda"
    echo "  -g, --greet    Saúda o usuário"
    echo "  -t, --time     Mostra a hora atual"
}

if [ $# -eq 0 ]; then
    usage
    exit 1
fi

case $1 in
    -h|--help)
        usage
        ;;
    -g|--greet)
        echo "Olá, usuário!"
        ;;
    -t|--time)
        echo "A hora atual é: $(date +"%H:%M:%S")"
        ;;
    *)
        echo "Opção inválida."
        usage
        exit 1
        ;;
esac