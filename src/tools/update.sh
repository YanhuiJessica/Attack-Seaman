#!/bin/bash

HOST="mongodb"
DB="attackSeaman"
COLLECTION=$1
USER="attackSeaman"
PSW="cuccs1sgreat"

dir=`pwd`
target_path='/nav-app/src/assets/new.json'
filepath=$dir$target_path
echo $filepath
# cd /app/nav-app/src/assets
mongoexport --host "$HOST" -u "$USER" -p "$PSW" --db "$DB"  --collection "$1" --jsonArray --pretty > $filepath

sed -i '1s/^/{"type": "bundle","id": "bundle--ad5f3bce-004b-417e-899d-392f8591ab55","spec_version": "2.0","objects":/' $filepath
# 在尾部添加 }
echo '}'>> $filepath


