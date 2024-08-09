#!/bin/bash

# Aguarda o banco de dados estar disponível
./wait-for-it.sh ${WSRS_DATABASE_HOST}:${WSRS_DATABASE_PORT} --timeout=60 --strict -- \
&& go generate ./... \
&& ./main
