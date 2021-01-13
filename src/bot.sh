#!/bin/bash

go get -u github.com/go-sql-driver/mysql && \
go get -u gopkg.in/tucnak/telebot.v2 && \
go get -u gorm.io/gorm && \
go get -u gorm.io/driver/mysql && \
go install entity && \
go install repositories && \
go install httpclient && \
go install httpclient && \

go run main.go
