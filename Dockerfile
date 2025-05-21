FROM	golang:1.22-alpine AS builder
WORKDIR	/src
COPY	go.mod go.sum ./
RUN	go mod tidy
COPY	. .

RUN	ls -lah /src && \
	cd /src/apps/server && \
	env GOOS=linux GOARCH=amd64 go build -o /src/queued

FROM	scratch
WORKDIR	/app
COPY	--from=builder /src/queued /app/queued
EXPOSE	3200
CMD	["./queued", "server"]
