#!/bin/bash

nohup npm start --prefix /app/nav-app/ >nav-app.log 2>&1 &
/app/main