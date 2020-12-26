#!/bin/bash

go get -u github.com/go-sql-driver/mysql && \
go install entity && \
go install repositories && \
go install httpclient && \
go install httpclient && \

go run ticker.go
