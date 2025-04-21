# User Service Documentation

## Introduction

Welcome to the User Service in this service, You can manage requests with functions userCreated, userUpdated, userReaded and userDeleted.  The last in this service can produce and consume about the topic of Games Service.

## Produce Topic / Endpoints

### POST /v1/users

Create a user account with topic "user.created".

**Request Body:**

```json
{
  "email": "theeranan.h@kkumail.com",
  "username": "theeranan.h",
  "password": "password"
}
```
**Response**:

```json
{
    "message": "Register success",
    "result": {
        "id": 26,
        "username": "theeranan.h"
    },
    "status": "OK",
    "status_code": 200
}
```
**produce**:
```json
Topic : user.created
```

### PUT /v1/users/{user_id}

Update a user account with topic "userUpdated".

**Request Body:**

```json
{
    "username": "theeranan.h_update",
    "email": "theeranan.h@kkumail.com",
    "password": "newpassword"
}
```

**Response**:

```json
{
    "message": "User updated successfully",
    "result": {
        "id": 26,
    "username": "theeranan.h_update",
    "email": "theeranan.h@kkumail.com",
    },
    "status": "OK",
    "status_code": 200
}
```

### GET /v1/users/{user_id}/read

In this endpoint, you can get all the Games data that you have ever read

**Request Body:**

```json
{
    "id": 26
}
```

**Response**:

```json
{
    "message": "UserReaded retrieved successfully",
    "result": [
        {
            "GameID": "523",
            "Title": "Fall Guys",
            "Genre": "Battle Royale",
            "Platform": "PC (Windows)",
            "Publisher": "Mediatonic",
            "Developer": "Mediatonic"
        },
        {
            "GameID": "516",
            "Title": "PUBG: BATTLEGROUNDS",
            "Genre": "Shooter",
            "Platform": "PC (Windows)",
            "Publisher": "KRAFTON, Inc.",
            "Developer": "KRAFTON, Inc."
        }
    ],
    "status": "OK",
    "status_code": 200
}
```

### DELETE /v1/users/{user_id}

Delete a user account with topic "user.deleted".

**Request Body:**

```json
{
    "id": 26
}
```


**Response**:

```json
{
    "id": 26,
    "message": "User deleted successfully",
    "status": "OK",
    "status_code": 200
}
```

**produce**:
```json
Topic : user.deleted
```

## Consume Topic

### user.read

When Games Service was produce data with this topic, it will keep the data into the database.

**Request Massage**:

```json
{
    "UserID": "26",
    "GameID": "516",
    "Title": "PUBG: BATTLEGROUNDS",
    "Genre": "Shooter",
    "Platform": "PC (Windows)",
    "Publisher": "KRAFTON, Inc.",
    "Developer": "KRAFTON, Inc."
}
```