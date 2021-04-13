#!/bin/bash

nohup yarn --cwd /app/nav-app/ start>nav-app.log 2>&1 &
/app/main