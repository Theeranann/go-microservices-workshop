# Games Service Documentation

## Introduction

Welcome to the Games Service. In this service, You can manage requests with the function userRead and this service can produce and consume about the topic of Users Service.

## API Data
An API Games data.

``https://www.freetogame.com/api/games`` -->

## Produce Topic / Endpoint

### GET /v1/games

In this endpoint, you can get all the Games data


**Response**:

```json
{
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
        },
		{...}
    ]
}
```

### GET /v1/games/{game_id}

In this endpoint, you can get the Games data by game_id

**Request Body:**

```json
{
    "id": 523
}
```

**Response**:

```json
{
    "result": [
        {
            "GameID": "523",
            "Title": "Fall Guys",
            "Genre": "Battle Royale",
            "Platform": "PC (Windows)",
            "Publisher": "Mediatonic",
            "Developer": "Mediatonic"
        }
    ]
}
```


### POST /v1/games

In this endpoint, it will check your request that you have an account or not. If you have an account, you can get the Games data with id of games and this function will produce this data about topic : ReadedEvent

**Request Body:**

```json
{
    "UserID": "26",
    "GameID": "523"
}
```

**Response**:

```json
{
	"GameID": "523",
	"Title": "Fall Guys",
	"Genre": "Battle Royale",
	"Platform": "PC (Windows)",
	"Publisher": "Mediatonic",
	"Developer": "Mediatonic"
}
```
**produce**:
```json
Topic : user.read
```

## Consume Topic

### userCreated

when userService was produce data with this topic, it will create a user account to database  

**Request Massage**:

```json
{
    "UserID": "26"
}
```

### userDeleted

when userService was produce data with this topic, it will Delete a user account in database  

**Request Massage**:

```json
{
    "UserID": "26"
}
```