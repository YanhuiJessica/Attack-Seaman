#!/bin/bash

HOST="mongodb"
DB="attackSeaman"
COLLECTION="mitre_attack"
USER="attackSeaman"
PSW="cuccs1sgreat"

cd /app/nav-app/src/assets
mongoexport --host "$HOST" -u "$USER" -p "$PSW" --db "$DB"  --collection "$COLLECTION" --jsonArray --pretty > 'new.json'

sed -i '1s/^/{"type": "bundle","id": "bundle--ad5f3bce-004b-417e-899d-392f8591ab55","spec_version": "2.0","objects":/' new.json
# 在尾部添加 }
echo '}'>>new.json


