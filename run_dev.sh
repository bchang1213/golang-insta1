#!/usr/bin/env bash
go run cmd/main.go &
fswatch --batch-marker=EOF -xn . | while read file event; do
    if [[ ${file} =~ \.go$ ]] || [[ ${file} =~ \.gohtml$ ]]; then
        echo File Changed: ${file}
        go run cmd/main.go &
        lsof -i tcp:8080 | awk 'NR!=1 {print $2}' | xargs kill
    fi
done
