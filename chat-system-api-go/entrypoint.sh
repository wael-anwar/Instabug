#!/bin/bash

# /usr/bin/wait-for-it.sh redis:6379 -t 0
cd /chat-system-api-go/cmd/chat-system-api-go

go build

/usr/bin/wait-for-it.sh chat-system-api-rails:3000 -t 0
# Then exec the container's main process (what's set as CMD in the Dockerfile).
exec "$@"
