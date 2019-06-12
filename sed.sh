#!/bin/sh

for i in `find . -name '*.go'`;
do
 	sed -i 's/github.com\/k0kubun\/sqldef/github.com\/kawakami-o3\/sqldef-sandbox/' $i
done

