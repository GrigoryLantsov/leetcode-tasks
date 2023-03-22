#!/bin/bash
if [ $(cat file.txt | wc -l) -ge 10 ]; then
    sed -n '10p' file.txt
fi
