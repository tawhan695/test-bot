#!/bin/bash
cd ~
echo 3 > /proc/sys/vm/drop_caches
cp test-bot/data.json data.json.backup
cp test-bot/token.txt token.txt.backup
rm -r test-bot
git clone https://github.com/tawhan695/test-bot.git
cp data.json.backup test-bot/data.json
cp token.txt.backup test-bot/token.txt
cd test-bot
go build main.go
./main token
