#! /bin/bash

dir="${HOME}/.GoChat/.key"

mkdir -p ${dir}

openssl rand -writerand "${HOME}/.rnd"

openssl genrsa -out "${dir}/server.key" 2048

printf "CN\nZhejiang\nHangzhou\n.\n.\n.\n.\n" | xargs `openssl req -new -x509 -sha256 -key "${dir}/server.key" -out "${dir}/server.crt" -days 3650`
