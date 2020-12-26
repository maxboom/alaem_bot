FROM golang

RUN go get -u github.com/go-sql-driver/mysql && \
    go get -u gopkg.in/tucnak/telebot.v2 && \
    go install entity && \
    go install repositories && \
    go install httpclient && \
    go install telegramapi
