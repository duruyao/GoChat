#!/bin/bash

key_path="${HOME}/.GoChat/.key/.server.key"
crt_path="${HOME}/.GoChat/.key/.server.crt"
input="CN\nZhejiang\nHangzhou\n.\n.\n.\n.\n"

openssl rand -writerand "${HOME}/.rnd" && \
openssl genrsa -out "${key_path}" 2048 && \
printf "${input}" | xargs `openssl req -new -x509 -sha256 -key "${key_path}" -out "${crt_path}" -days 3650`
