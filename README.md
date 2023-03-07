Blog App in Go 
 - Backend API for a simple blog system

WHAT CAN YOU DO WITH THIS APP?
 - Publish an article with Title and Content by hitting CREATE POST API
 - Read a particular article by hitting the GET POST API
 - Read all the articles published by hitting GET ALL POSTS API

 APIs listening to port 8080


Create an article
Method: POST
Path: /articles
Request Body:
{
    "title": "Hello World",
    "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
    "author": "John",
}
Response Body:
{
    "status": 201,
    "message": "Success",
    "data": {
      "id": <article_id>
    }
}

Get article by id
Method: GET
Path: /articles/<article_id>
Response Header: HTTP 200
Response Body:
{
    "status": 200,
    "message": "Success",
    "data": [
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      }
    ]
}

Get all article
Method: GET
Path: /articles
Response Header: HTTP 200
Response Body:
{
    "status": 200,
    "message": "Success",
    "data": [
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      },
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      }
    ]
}


