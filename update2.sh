#!/bin/bash
cd ~
echo 1 > /proc/sys/vm/drop_caches
echo 2 > /proc/sys/vm/drop_caches
echo 3 > /proc/sys/vm/drop_caches
cp test-bot/data2.json data2.json.backup
cp test-bot/token2.txt token2.txt.backup
rm -r test-bot
cp data2.json.backup test-bot/data2.json
cp token.txt2.backup test-bot/token2.txt
cd test-bot
go build main2.go
cd ~
cp test-bot/update2.sh update2.sh
chmod +x update2.sh
cp test-bot/bot2.sh bot2.sh
chmod +x bot2.sh
