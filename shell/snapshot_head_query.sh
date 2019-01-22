#!/bin/bash
set -e

curl -X POST \
  http://127.0.0.1:48132/ \
  -H 'content-type: application/json' \
  -d '{
	"jsonrpc": "2.0",
	"id": 1,
	"method":"ledger_getSnapshotChainHeight",
	"params":null
}'
