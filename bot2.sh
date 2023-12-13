#!/bin/bash
cd ~
echo 3 > /proc/sys/vm/drop_caches
cd test-bot
./main2
