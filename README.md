# Instabug Chat System Backend Challenge
Instabug Backend Challenge (Ruby on Rails) and Golang APIs

Table of Contents
=================
  * [Setup](#setup)
  * [Usage](#usage)

## Setup
Fistly, clone this repository:
```
git clone https://github.com/wael-anwar/Instabug.git
```

### Run Docker Container for the first time
 ```bash
  - cd Instabug
  - docker-compose up
  ```

## Usage
  
### Ruby on Rails Part
 - Chat System API for the Application and Searching http://localhost:3000/api/v1/
    ```
    Type   Endpoint
    ----  -----------

    GET   /applications/
    GET   /applications/{access_token}
    GET   /applications/{access_token}/chats
    GET   /applications/{access_token}/chats/{chat_number}
    GET   /applications/{access_token}/chats/{chat_number}/messages
    GET   /applications/{access_token}/chats/{chat_number}/messages/{message_number}
    GET   /applications/{access_token}/chats/{chat_number}/messages/search?keyword={keyword}

    POST  /applications?name={name}

    PUT   /applications/{access_token}?name={name}
    PUT   /applications/{access_token}/chats/{chat_number}/messages/{message_number}?body={message_body}
    ```

**Examples**:

- Ex : Post Create Application Tested in Postman http://localhost:3000/api/v1/applications?name=InstaChat
  ```
  Post: in Headers
  {
    Content-Type: application/json
  }
  Post: in Query Param 
  {
    name = InstaChat
  }
  Receive:
  {
          "name": "InstaChat",
          "access_token": "qmhVkNi1NJx2ppXzvKM5appy",
          "chat_count": 0,
          "created_at": "2022-04-29T12:08:40.353Z",
          "updated_at": "2022-04-29T12:08:40.353Z"
  }
  ```

- Ex : Get All Applications Tested in Postman http://localhost:3000/api/v1/applications
  ```
  Receive:
  {
          "name": "InstaChat",
          "access_token": "qmhVkNi1NJx2ppXzvKM5appy",
          "chat_count": 0,
          "created_at": "2022-04-29T12:08:40.353Z",
          "updated_at": "2022-04-29T12:08:40.353Z"
  }
  ```

### Golang Part
 - Chat System API for the Chat and Messages http://localhost:5000/api/v1/
    ```
    Type   Endpoint
    ----  -----------

    POST  /applications/{access_token}/chats/
    POST  /applications/{access_token}/chats/{chat_number}/messages?body={message_body}
 
    ```

**Examples**:

- Ex : Post Create Chat Tested in Postman http://localhost:5000/api/v1/applications/qmhVkNi1NJx2ppXzvKM5appy/chats
  ```
  Post: in Headers
  {
    Content-Type: application/json
  }
  Receive:
  {
          "number": "1",
          "access_token": "qmhVkNi1NJx2ppXzvKM5appy"
  }
  ```

- Ex : Post Create Message Tested in Postman http://localhost:5000/api/v1/applications/qmhVkNi1NJx2ppXzvKM5appy/chats/1/messages
  ```
  Post: in Headers
  {
    Content-Type: application/json
  }
  Post: in body
  {
    body: "Test first Message from chat 1"
  }
  Receive:
  {
          "number": "1",
          "chat_number": "1",
          "access_token": "qmhVkNi1NJx2ppXzvKM5appy"
  }
  ```
