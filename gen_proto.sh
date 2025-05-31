#!/bin/bash
# File: test_client.sh
# Description: Script for testing the gacha client
# Author: [Your Name]
# Date: [Date]
# Usage: ./test_client.sh
# Notes: Replace [Your Name] and [Date] with appropriate values.

if [ $# -eq 0 ]
    then
        echo "No arguments supplied"
fi

# IFS="." read $FILENAME $EXTENSION <<< $1
PROTO_FILE="$1"

FILENAME="$(basename "$PROTO_FILE" .proto)"

CURRENT_DIR="."
PROTO_PATH="${CURRENT_DIR}/proto"

if [ ! -d "$PROTO_PATH/${FILENAME}" ]; then
    mkdir $PROTO_PATH/${FILENAME}
fi

OUTPUT_PATH="${PROTO_PATH}/${FILENAME}/."

PROTO_OUT="${PROTO_PATH}/${FILENAME}/github.com/inonsdn/gacha-system/proto/${FILENAME}"

protoc --go_out=$OUTPUT_PATH --go-grpc_out=$OUTPUT_PATH $1

mv $PROTO_OUT/* $OUTPUT_PATH
rm -rf $PROTO_PATH/${FILENAME}/github.com

echo "Generate go file at ${OUTPUT_PATH}"