FROM golang

ADD . /go/src/github.com/tasks/TaskDillyDilly

RUN go install github.com/tasks/TaskDillyDilly

ENTRYPOINT [ "/go/bin/TaskDillyDilly" ]

EXPOSE 8080