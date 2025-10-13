FROM	golang:1.22.3-alpine AS builder

# Init required packages
WORKDIR	/app
COPY	go.mod go.mod
RUN	go mod tidy
COPY	. .

# Build Queued Server
WORKDIR	/app/apps/server
RUN	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags "-s -w" -o queued
# Build Qeueued Control
WORKDIR	/app/apps/control
RUN	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags "-s -w" -o qctl

# Final image
FROM	alpine:latest
COPY	--from=builder /app/apps/server/queued /app/apps/control/qctl /bin/
RUN	mkdir -p /var/log/queued && \
	mkdir -p /etc/queued/conf.d 

# Setup local user
# Create a new group
RUN	addgroup -S queued
RUN	adduser -S queued -G queued

RUN	mkdir -p /var/log/queued && \
	mkdir -p /etc/queued/conf.d && \
	chown queued:queued /var/log/queued && \
	chown queued:queued /etc/queued

USER	queued

# Setup default env
ENV	\
	QUEUED_API_HTTPLISTEN=127.0.0.1:3200 \
	QUEUED_API_CORS=* \
	QUEUED_API_AUTH_ENABLED=false \
	QUEUED_LOG_LEVEL=info \
	QUEUED_LOG_LOCATION=/var/log/queued \
	QUEUED_INCLUDE=/etc/queued/conf.d
	
EXPOSE	3200

CMD	["/bin/queued", "server"]
