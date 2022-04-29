package router

import (
    "github.com/gorilla/mux"
    "github.com/wael-anwar/chat-system-api-go/internal/handlers"
    "github.com/wael-anwar/chat-system-api-go/configs"
)

func InitRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    
    router.HandleFunc(configs.ChatsRoute, handlers.CreateChat).Methods("POST")
    router.HandleFunc(configs.MessagesRoute, handlers.CreateMessage).Methods("POST")

    return router
}
