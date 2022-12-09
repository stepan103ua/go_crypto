# GO Crypto

## Sign Up:
Method: __POST__
Endpoint: __/auth/signUp__

__Body example:__ 
```json
{
  "email": "example@gmail.com",
  "name": "username",
  "password": "password"
}
```

__Response body example:__
```json
{
  "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzA2MzA1MzYsImlhdCI6MTY3MDU4NzMzNiwidXNlcl9pZCI6NH0.CasSTbAPQN9CZGvaYRNOxoweNWsrLcGkbaXL_DEArLs"
}
```


## Sign In:
Method: __POST__
Endpoint: __/auth/signIn__

__Body example:__ 
```json
{
  "email": "example@gmail.com",
  "password": "password"
}
```

__Response body example:__
```json
{
  "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzA2MzA1MzYsImlhdCI6MTY3MDU4NzMzNiwidXNlcl9pZCI6NH0.CasSTbAPQN9CZGvaYRNOxoweNWsrLcGkbaXL_DEArLs"
}
```

## Get All Posts:
Method: __GET__
Endpoint: __/api/posts__

__Response body example:__
```json
{
    "posts": [
        {
            "id": 1,
            "title": "rest",
            "description": "description",
            "image_url": "url",
            "crypto_currency": "USD",
            "created_at": "2022-12-08T17:56:43.333331Z",
            "owner_id": 1,
            "owner_username": "Stepan0000"
        },
        {
            "id": 2,
            "title": "rest",
            "description": "description",
            "image_url": "url",
            "crypto_currency": "USD",
            "created_at": "2022-12-08T20:30:05.164461Z",
            "owner_id": 1,
            "owner_username": "Stepan0000"
        },
        {
            "id": 3,
            "title": "restf",
            "description": "description",
            "image_url": "url",
            "crypto_currency": "USD",
            "created_at": "2022-12-08T21:03:04.645295Z",
            "owner_id": 4,
            "owner_username": "Stepan00013"
        }
    ]
}
```

## Get Post By Id:
Method: __GET__
Endpoint: __/api/posts/1__ _(post id)_

__Response body example:__
```json
{
    "id": 1,
    "title": "rest",
    "description": "description",
    "image_url": "url",
    "crypto_currency": "USD",
    "created_at": "2022-12-08T17:56:43.333331Z",
    "owner_id": 1,
    "owner_username": "Stepan0000"
}
```

## Delete Post:
Method: __DELETE__
Endpoint: __/api/posts/1__ _(post id)_

__Response body example:__
```json
{
    "message": "Successfully deleted"
}
```

## Update Post:
Method: __PUT__
Endpoint: __/api/posts/2__ _(post id)_

__Body example:__ 
```json
{
    "title": "update 1",
    "description": "description",
    "image_url": "url",
    "crypto_currency": "USD"
}
```

__Response body example:__
```json
{
    "message": "Successfully updated"
}
```

## Create Post:
Method: __POST__
Endpoint: __/api/posts/new__

__Body example:__ 
```json
{
    "title": "new post 1",
    "description": "description",
    "image_url": "url",
    "crypto_currency": "USD"
}
```

__Response body example:__
```json
{
    "id": 4
}
```


## Get All Authenticated User Posts:
Method: __GET__
Endpoint: __/api/posts/myAll__

__Response body example:__
```json
{
    "posts": [
        {
            "id": 3,
            "title": "restf",
            "description": "description",
            "image_url": "url",
            "crypto_currency": "USD",
            "created_at": "2022-12-08T21:03:04.645295Z",
            "owner_id": 4,
            "owner_username": "Stepan00013"
        },
        {
            "id": 4,
            "title": "update 1",
            "description": "description",
            "image_url": "url",
            "crypto_currency": "USD",
            "created_at": "2022-12-09T12:43:40.273297Z",
            "owner_id": 4,
            "owner_username": "Stepan00013"
        }
    ]
}
```

## Get All Posts By User Id:
Method: __GET__
Endpoint: __/api/posts/users/1__ _(user id)_

__Response body example:__
```json
{
    "posts": [
        {
            "id": 1,
            "title": "rest",
            "description": "description",
            "image_url": "url",
            "crypto_currency": "USD",
            "created_at": "2022-12-08T17:56:43.333331Z",
            "owner_id": 1,
            "owner_username": "Stepan0000"
        },
        {
            "id": 2,
            "title": "rest",
            "description": "description",
            "image_url": "url",
            "crypto_currency": "USD",
            "created_at": "2022-12-08T20:30:05.164461Z",
            "owner_id": 1,
            "owner_username": "Stepan0000"
        }
    ]
}
```

## Like/Dislike Post:
Method: __POST__
Endpoint: __/api/posts/2__ _(post id)_ __/toggleLike__

__Response body example:__
```json
{
    "message": "ok"
}
```

## Get Post Likes:
Method: __GET__
Endpoint: __/api/posts/2__ _(post id)_ __/likes__

__Response body example:__
```json
{
    "isLiked": true,
    "likes": 1
}
```

## Create Comment:
Method: __POST__
Endpoint: __/api/comments/new__

__Body example:__ 
```json
{
    "post_id": 2,
    "comment": "Comment 1"
}
```

__Response body example:__
```json
{
    "id": 1
}
```

## Delete Comment:
Method: __DELETE__
Endpoint: __/api/comments/1__

__Response body example:__
```json
{
    "message": "Successfully deleted"
}
```

## Update Comment:
Method: __PUT__
Endpoint: __/api/comments/1__

__Body example:__ 
```json
{
    "comment": "Comment 2"
}
```

__Response body example:__
```json
{
    "message": "Successfully updated"
}
```

## Get Comments By Post Id:
Method: __GET__
Endpoint: __/api/comments/posts/1__

__Response body example:__
```json
{
    "comments": [
        {
            "id": 2,
            "post_id": 2,
            "comment": "Comment 2",
            "owner_id": 4,
            "owner_username": "Stepan00013"
        }
    ]
}
```

## Get Comments Count By Post Id:
Method: __GET__
Endpoint: __/api/comments/posts/1/count__

__Response body example:__
```json
{
    "count": 1
}
```

## Create reply:
Method: __POST__
Endpoint: __/api/comments/1/replies/new__

__Body example:__ 
```json
{
    "reply": "reply 1"
}
```

__Response body example:__
```json
{
    "id": 2
}
```

## Get All Replies By Comment Id:
Method: __GET__
Endpoint: __/api/comments/1/replies__

__Response body example:__
```json
{
    "replies": [
        {
            "id": 2,
            "comment_id": 2,
            "owner_id": 4,
            "reply": "reply 1",
            "owner_username": "Stepan00013"
        }
    ]
}
```

## Get User By Id:
Method: __GET__
Endpoint: __/api/users/2__

__Response body example:__
```json
{
    "id": 2,
    "name": "Stepan0001",
    "email": "semaila@gmail.com",
    "about": ""
}
```

## Update User:
Method: __PUT__
Endpoint: __/api/users/update__

__Body example:__ 
```json
{
    "name": "Stepan0002",
    "about": "New About",
    "password": "123456"
}
```

__Response body example:__
```json
{
    "message": "ok"
}
```

## Get Authenticated User:
Method: __GET__
Endpoint: __/api/users/me__

__Response body example:__
```json
{
    "id": 4,
    "name": "Stepan00013",
    "email": "seees@gmail.com",
    "about": ""
}
```

## Follow User:
Method: __POST__
Endpoint: __/api/users/1/toggleFollow__

__Response body example:__
```json
{
    "isFollowing": true
}
```

## Get User Followers Count:
Method: __GET__
Endpoint: __/api/users/1/followersCount__

__Response body example:__
```json
{
    "count": 1
}
```

## Get User Followings Count:
Method: __GET__
Endpoint: __/api/users/1/followingCount__

__Response body example:__
```json
{
    "count": 1
}
```

## Is Following User:
Method: __GET__
Endpoint: __/api/users/1/isFollowing__

__Response body example:__
```json
{
    "isFollowing": true
}
```

## Create Watchlist:
Method: __POST__
Endpoint: __/api/watchlists/new__

__Body example:__ 
```json
{
    "watchlist": "watchlist 1",
    "crypto_currency": "BTC"
}
```

__Response body example:__
```json
{
    "message": "Successfully created"
}
```

## Get All User watchlists:
Method: __GET__
Endpoint: __/api/watchlists__

__Response body example:__
```json
[
    {
        "id": 1,
        "watchlist": "watchlist 1",
        "crypto_currency": "crypto_currency 1"
    },
    {
        "id": 2,
        "watchlist": "watchlist 2",
        "crypto_currency": "crypto_currency 2"
    },
    {
        "id": 3,
        "watchlist": "watchlist 1",
        "crypto_currency": "BTC"
    }
]
```
