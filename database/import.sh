#! /bin/bash

mongoimport --host mongodb -u attackSeaman -p cuccs1sgreat --db attackSeaman --collection enterprise --drop --type json --file /database/enterprise-attack.json --jsonArray

# 将 mongodb 中string字段转化为 Date，方便排序
mongo --host mongodb -u attackSeaman -p cuccs1sgreat --eval "db.enterprise.find().forEach(function(doc){doc.created = new Date(doc.created);db.mitre_attack.save(doc)});"
mongo --host mongodb -u attackSeaman -p cuccs1sgreat --eval "db.enterprise.find().forEach(function(doc){doc.modified = new Date(doc.modified);db.mitre_attack.save(doc)});"
