#! /bin/bash

mongoimport  --host mongodb -u attackSeaman -p cuccs1sgreat --db attackSeaman --collection $1 --drop --type json --file $2 --jsonArray
mongo  --host mongodb -u attackSeaman -p cuccs1sgreat --eval "db.$1.find().forEach(function(doc){doc.created = new Date(doc.created);db.$1.save(doc)});"
mongo  --host mongodb -u attackSeaman -p cuccs1sgreat --eval "db.$1.find().forEach(function(doc){doc.modified = new Date(doc.modified);db.$1.save(doc)});"
