#!/bin/bash

HOST="mongodb"
DB="attackSeaman"
COLLECTION="mitre_attack"
mkdir /app/attack-navigator

mongoexport --host "$HOST" --db "$DB" -u attackSeaman -p cuccs1sgreat --collection "$COLLECTION" --jsonArray --pretty > '/app/attack-navigator/new.json'
cd /app/attack-navigator/

sed -i '1s/^/{"type": "bundle","id": "bundle--ad5f3bce-004b-417e-899d-392f8591ab55","spec_version": "2.0","objects":/' new.json
# 在尾部添加 }
echo '}'>>new.json

# git add .
# git commit -m "update src"
# git push -f fork


