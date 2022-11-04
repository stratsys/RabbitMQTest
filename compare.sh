#!/bin/bash

file1=/home/farsta/first.json
file2=/home/farsta/second.json
file3=/home/farsta/third.json

while read -r line 
do  
grep "$line" $file2
#grep -c "$line" $file3
done < $file1