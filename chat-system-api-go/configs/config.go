package configs

const AppEndpoint = "http://chat-system-api-rails:3000/"
const ChatsRoute = "/api/v1/applications/{access_token}/chats"
const MessagesRoute = "/api/v1/applications/{access_token}/chats/{chat_number}/messages"

const RedisChatQueue = "chat"
const ChatWorkerClass = "ChatWorker"
const RedisChatKeyPrefix = "CHAT_"

const MsgRedisQueue = "message"
const MsgWorkerClass = "MessageWorker"
const MsgRedisKeyPrefix = "MSG_"

const RedisAddress = "redis:6379"
