#!/bin/bash

rm -f mythictools
GOOS=linux GOARCH=amd64 go build -o mythictools

if [ -f "mythictools" ]; then
    rm -f deploymt.zip
    zip -r deploymt.zip mythictools assets templates data .env

    if [ -f "deploymt.zip" ]; then
        scp -i ~/culpanvtt_org deploymt.zip mythictools@culpanvtt.org:~
    fi
fi

