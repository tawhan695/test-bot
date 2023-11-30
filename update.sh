#!/bin/bash
cd ~
echo 3 > /proc/sys/vm/drop_caches
cp test-bot/data.json data.json.backup
# cp test-bot/token.txt token.txt.backup
rm -r test-bot
git clone https://github.com/tawhan695/test-bot.git
cp data.json.backup test-bot/data.json
# cp token.txt.backup test-bot/token.txt
cd test-bot
go build main.go
cd ~
cp test-bot/update.sh update.sh
chmod +x update.sh
cp test-bot/bot.sh bot.sh
# cp botService.service /etc/systemd/system/botService.service
chmod +x bot.sh
# systemctl enamble botService.service
# systemctl disable botService.service
# systemctl start botService.service
# systemctl status botService.service
reboot
