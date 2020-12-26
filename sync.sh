#!/bin/bash


rm -rf /home/kyrylokramarenko/go/src/entity
rm -rf /home/kyrylokramarenko/go/src/repositories
rm -rf /home/kyrylokramarenko/go/src/httpclient
rm -rf /home/kyrylokramarenko/go/src/callmebotapi
rm -rf /home/kyrylokramarenko/go/src/telegramapi


/bin/cp -rf /home/kyrylokramarenko/Lessons/Go/AlarmBot/src/entity /home/kyrylokramarenko/go/src/entity
/bin/cp -rf /home/kyrylokramarenko/Lessons/Go/AlarmBot/src/repositories /home/kyrylokramarenko/go/src/repositories
/bin/cp -rf /home/kyrylokramarenko/Lessons/Go/AlarmBot/src/httpclient /home/kyrylokramarenko/go/src/httpclient
/bin/cp -rf /home/kyrylokramarenko/Lessons/Go/AlarmBot/src/callmebotapi /home/kyrylokramarenko/go/src/callmebotapi
/bin/cp -rf /home/kyrylokramarenko/Lessons/Go/AlarmBot/src/telegramapi /home/kyrylokramarenko/go/src/telegramapi

go get -u github.com/go-sql-driver/mysql
go install entity
go install repositories
go install httpclient
go install callmebotapi
go install telegramapi