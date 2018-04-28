#!/bin/bash

cat input | while read line
do
    echo "${line}" | ./a.out
done
