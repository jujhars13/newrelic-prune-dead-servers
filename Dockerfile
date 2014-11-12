FROM golang:1.2.2
MAINTAINER Jujhar Singh <jujhar@jujhar.com>

#run our script
CMD ["go","run","prune-dead-servers.go","--api-key="]
