FORMAT: 1A

# golang-crud-spa API
[GitHub's golang-crud-spa](https://github.com/albertojnk/golang-crud-spa).

## Media Types
Requests with a message-body are using plain JSON to set or update resource states.

## Error States
The common [HTTP Response Status Codes](https://github.com/for-GET/know-your-http-well/blob/master/status-codes.md) are used.

# golang-crud-spa Root [/]
golang-crud-spa entry point.

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

## User/Repositories [/list?username={_id}]
Multiples github starred repositories objects of an single user. The List resource is the central resource in the golang-crud-spa. It represents one github user and all his starred github public repositories.

The List resource has the following attributes:

+ _id


+ Parameters
    + _id (string) - ID of the github user.

+ Model (application/json)

    + Headers

            Content-Type: application/json

    + Body

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
<!-- 
### Retrieve a Single Gist [GET]
+ Response 200

    [Gist][]

### Edit a Gist [PATCH]
To update a Gist send a JSON with updated value for one or more of the Gist resource attributes. All attributes values (states) from the previous version of this Gist are carried over by default if not included in the hash.

+ Request (application/json)

        {
            "content": "Updated file contents"
        }

+ Response 200

    [Gist][]

### Delete a Gist [DELETE]
+ Response 204

## Gists Collection [/gists{?since}]
Collection of all Gists.

The Gist Collection resource has the following attribute:

+ total

In addition it **embeds** *Gist Resources* in the golang-crud-spa.


+ Model (application/hal+json)

    HAL+JSON representation of Gist Collection Resource. The Gist resources in collections are embedded. Note the embedded Gists resource are incomplete representations of the Gist in question. Use the respective Gist link to retrieve its full representation.

    + Headers

            Link: <http:/api.gistfox.com/gists>;rel="self"

    + Body

            {
                "_links": {
                    "self": { "href": "/gists" }
                },
                "_embedded": {
                    "gists": [
                        {
                            "_links" : {
                                "self": { "href": "/gists/42" }
                            },
                            "id": "42",
                            "created_at": "2014-04-14T02:15:15Z",
                            "description": "Description of Gist"
                        }
                    ]
                },
                "total": 1
            }

### List All Gists [GET]
+ Parameters
    + since (string, optional) - Timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ` Only gists updated at or after this time are returned.

+ Response 200

    [Gists Collection][]

### Create a Gist [POST]
To create a new Gist simply provide a JSON hash of the *description* and *content* attributes for the new Gist.

+ Request (application/json)

        {
            "description": "Description of Gist",
            "content": "String content"
        }

+ Response 201

    [Gist][]

## Star [/gists/{id}/star]
Star resource represents a Gist starred status.

The Star resource has the following attribute:

+ starred


+ Parameters

    + id (string) - ID of the gist in the form of a hash

+ Model (application/hal+json)

    HAL+JSON representation of Star Resource.

    + Headers

            Link: <http:/api.gistfox.com/gists/42/star>;rel="self"

    + Body

            {
                "_links": {
                    "self": { "href": "/gists/42/star" }
                },
                "starred": true
            }

### Star a Gist [PUT]
+ Response 204

### Unstar a Gist [DELETE]
+ Response 204

### Check if a Gist is Starred [GET]
+ Response 200

    [Star][] -->