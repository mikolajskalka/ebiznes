#!/bin/bash

GENERATED_SECRET=$(head -c 32 /dev/urandom | base64 | tr -d '\n')

exec /app/bin/e-shop-api -Dplay.http.secret.key="$GENERATED_SECRET" "$@"