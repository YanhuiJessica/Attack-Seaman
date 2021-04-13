#! /bin/bash

mongoimport --host mongodb -u attackSeaman -p cuccs1sgreat --db attackSeaman --collection $1 --drop --type json --file $2 --jsonArray
# alpine 需要安装 mongodb 才能使用 mongo 这个命令，这样的话这个 container 太大了
# 影响不是特别大，但没法按time排序。想办法看是否可以远程执行
# mongo --host mongodb -u attackSeaman -p cuccs1sgreat --eval "db.$1.find().forEach(function(doc){doc.created = new Date(doc.created);db.$1.save(doc)});"
# mongo --host mongodb -u attackSeaman -p cuccs1sgreat --eval "db.$1.find().forEach(function(doc){doc.modified = new Date(doc.modified);db.$1.save(doc)});"
