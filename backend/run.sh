#!/usr/bin/env bash
set -e

killall app.o || true

pushd ./cmd/saver
go build -o app.o .
./app.o serve --gcp-api-key=$GCP_API_KEY --gcp-proj-id=$GCP_PROJ_ID --rest-address 127.0.0.1:8700 > run.log 2>&1 &
echo "Starting on port 8700"
popd
