#! /bin/bash
mongo --host mongodb -u attackSeaman -p cuccs1sgreat --eval "db.mitre_attack.remove({});"
mongoimport --host mongodb -u attackSeaman -p cuccs1sgreat --db attackSeaman --collection mitre_attack --type json --file /database/enterprise-attack.json --jsonArray
mongo --host mongodb -u attackSeaman -p cuccs1sgreat --eval "db.mitre_attack.find().forEach(function(doc){doc.created = new Date(doc.created);db.mitre_attack.save(doc)});"
mongo --host mongodb -u attackSeaman -p cuccs1sgreat --eval "db.mitre_attack.find().forEach(function(doc){doc.modified = new Date(doc.modified);db.mitre_attack.save(doc)});"
