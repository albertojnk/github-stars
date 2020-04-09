FORMAT: 1A

# github-stars API
[GitHub's github-stars](https://github.com/albertojnk/github-stars).

NOTE: This document is a **work in progress**.

## Error States
The common [HTTP Response Status Codes](https://github.com/for-GET/know-your-http-well/blob/master/status-codes.md) are used.

# github-stars Root [/]
github-stars entry point.

This resource does not have any attributes.

## Retrieve the Entry Point [GET]

+ Response 200 (application/hal+json)
    + Headers

            Access-Control-Allow-Origin: *
            Content-Type: text/html; charset=utf-8

    + Body

            <!DOCTYPE html>
            <html lang=en>

            <head>
                <meta charset=utf-8>
                <meta http-equiv=X-UA-Compatible content="IE=edge">
                <meta name=viewport content="width=device-width,initial-scale=1">
                <link rel=icon href=./favicon.ico> <title>Crud SPA</title>
                <link href=/css/app.a339d1f8.css rel=preload as=style>
                <link href=/css/chunk-vendors.b89365ac.css rel=preload as=style>
                <link href=/js/app.c62866fe.js rel=preload as=script>
                <link href=/js/chunk-vendors.f4725632.js rel=preload as=script>
                <link href=/css/chunk-vendors.b89365ac.css rel=stylesheet>
                <link href=/css/app.a339d1f8.css rel=stylesheet>
            </head>

            <body>
                <div id=app></div>
                <script src=/js/chunk-vendors.f4725632.js> </script> <script src=/js/app.c62866fe.js> </script> </body> </html>

### Create an User and his repositories [POST][/create]
To create an User send a JSON with his github username.

+ Request (application/json)

        {
            "username": "albertojnk"
        }

+ Response 201
    
    + Headers

            Content-Type: application/json

    + Body 

            {
                "_id": "albertojnk",
                "repositories": [
                    {
                        "id": 203700322,
                        "name": "github-stars",
                        "description": "A simple CRUD in a SPA",
                        "html_url": "https://github.com/albertojnk/github-stars",
                        "language": "Go",
                        "tags": [],
                        "tag_suggester": "Go"
                    }
                ]
            }

### List repositories [GET][/list]
+ Response 200

    + Headers

            Content-Type: application/json

    + Body 

            {
                "_id": "test_user1",
                "repositories": [
                    {
                        "id": 1,
                        "name": "test_name1",
                        "description": "test repository tags deletion 1",
                        "html_url": "http://google.com",
                        "language": "Gotest",
                        "tags": [
                            "Test1",
                            "Test2",
                            "Test3"
                        ],
                        "tag_suggester": "Gotest"
                    }
                ]
            }

# Group Tags
Tags-related resources of **github-stars**

## Repository & TAGS
Update overwrite the tags of a given repository ID. Repository is the central resource of **github-stars**. It represents a starred github repository of a given user.

The Repository resource has the following attributes:

+ id
+ name
+ description
+ html_url
+ language
+ tags
+ tag_suggester

### Update repository tags [PATCH][/update]

+ Request 

    + Headers

            Content-Type: application/json
    
    + Body

            {
                "username": "test_user1",
                "repo_id": 1,
                "tags": ["Test1", "Test2", "Test3", "Teste4", "Teste5"]
            }

+ Response 200

    + Headers

            Content-Type: application/json

    + Body

            [
                {
                    "id": 203700322,
                    "name": "github-stars",
                    "description": "A simple CRUD in a SPA",
                    "html_url": "https://github.com/albertojnk/github-stars",
                    "language": "Go",
                    "tags": [
                        "Test1",
                        "Test2",
                        "Test3",
                        "Teste4",
                        "Teste5"
                    ],
                    "tag_suggester": "Go"
                }
            ]

### Delete tags [DELETE][/delete]
+ Request
    
    + Headers

            Content-Type: application/json

    + Body

            {
                "username": "albertojnk",
                "repo_id": 203700322,
                "tags": ["Test1", "Test3"]
            }

+ Response 

    + Headers

            Content-Type: application/json

    + Body

            {
                "_id": "albertojnk",
                "repositories": [
                    {
                        "id": 203700322,
                        "name": "github-stars",
                        "description": "A simple CRUD in a SPA",
                        "html_url": "https://github.com/albertojnk/github-stars",
                        "language": "Go",
                        "tags": [
                            "Test2",
                            "Teste4",
                            "Teste5"
                        ],
                        "tag_suggester": "Go"
                    }
                ]
            }