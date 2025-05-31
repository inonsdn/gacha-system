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

curl  http://localhost:8080/$1
