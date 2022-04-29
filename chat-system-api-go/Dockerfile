from golang:1.18.1
RUN mkdir /chat-system-api-go
WORKDIR /chat-system-api-rails
COPY . /chat-system-api-go
COPY entrypoint.sh /usr/bin/entrypoint-go.sh
RUN chmod +x /usr/bin/entrypoint-go.sh
COPY wait-for-it.sh /usr/bin
RUN chmod +x /usr/bin/wait-for-it.sh
ENTRYPOINT ["entrypoint-go.sh"]
EXPOSE 5000

CMD ["/chat-system-api-go/cmd/chat-system-api-go"]
