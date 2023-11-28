#!/bin/bash
echo 3 > /proc/sys/vm/drop_caches
cp test-bot/data.json data.json.backup
rm -r test-bot
git clone https://github.com/tawhan695/test-bot.git
cp data.json.backup test-bot/data.json
go build test-bot/main.go
./main token