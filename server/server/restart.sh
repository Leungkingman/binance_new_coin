#!/bin/bash

server_ips=(
  "8.218.81.24"
  "8.210.77.136"
  "8.210.127.104"
  "8.210.3.245"
  "8.210.92.126"
  "47.242.77.62"
  "8.210.217.245"
  "47.242.207.148"
  "47.243.59.19"
  "8.210.174.216"
)

for ip in ${server_ips[*]}
do
  ssh root@$ip "
    killall -9 bnl;
    cd /www;
    nohup ./bnl > nohup.out 2>&1 &
    exit;
  "
done
