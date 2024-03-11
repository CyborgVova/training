#!/bin/bash

i=$1

while [[ $i -gt 0 ]]; do
echo $i
sleep 1
i=$(($i-1))
done
echo "Boom ..."
sleep 1

exec 6<&0
exec 0< myfile

while read line; do
echo $line
done

exec 0<&6

read -p "enter(y/n): " v
case $v in
y) echo "Y"
n) echo "N"
esac
