FROM golang:1.2.2
MAINTAINER Jujhar Singh <jujhar@jujhar.com>

#run our script
ENTRYPOINT ["go","run","prune-dead-servers.go"]

#DEFAULT params if --api-key not sent in
CMD ["--help"]
