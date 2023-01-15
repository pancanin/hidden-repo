# Questions API

Hi! :)
Here is what I did for the time being. :man-mechanic:

## Tested on

- Linux and Go 1.19
- MacOS and Go 1.17, so you can downgrade go in go.mod to 1.17 and it should build.

## Basic requirements

Your solution should meet all these requirements.

- [x] Endpoint that returns a list of all questions
- [x] Endpoint that allows to add a new question
- [x] Endpoint that allows to update an existing question
- [x] Question data is stored in a SQLite database with a **normalised** schema
- [x] The order of questions and options is stable, not random
- [x] The `PORT` environment variable is used as the port number for the server, defaulting to 3000

## Bonus requirements

These requirements are not required, but feel free to complete some of them if they seem interesting, or to come up with your own :)

- [x] Endpoint that allows to delete existing questions
- [x] Pagination for the list endpoint

  This can be in the form of basic offset pagination, or seek pagination. The difference is explained in [this post](https://web.archive.org/web/20210205081113/https://taylorbrazelton.com/posts/2019/03/offset-vs-seek-pagination/).

- [x] JWT authentication mechanism
  
  Clients are required to send a JSON Web Token that identifies the user in some way. The API returns only questions that belong to the authenticated user. Endpoint for generating tokens is not needed, we can generate them through [jwt.io](https://jwt.io/).

- [ ] Use GraphQL instead of REST to implement the API

  Define a schema for the API that covers the basic requirements and implement all queries and resolvers. You do not need to implement the REST API if you choose to do this.

## How to run it

All dependencies are in go.mod. Tested on Linux.

```
cd proj_folder
go build
./questions
```

## The API

I added a postman collection, but not everyone uses that, so here it is...

With auth enabled, set `Authorization` header with value `Bearer token_here`.

- POST /api/v1/question - Create question

Payload:

```
{
    "body": "Is Kotlin pass by reference or pass by value?",
    "options": [
        {
            "body": "By reference",
            "correct": true
        },
        {
            "body": "By value",
            "correct": false
        }
    ]
}
```

- GET /api/v1/questions?page=N&page_size=M - Get questions with pagination

- PUT /api/v1/question/{id} - Update question by id.

Payload:

```
{
    "body": "How do you define a variable that does not hoist in JS? Huh, whut?",
    "options": [
        {
            "body": "With var keyword ofc",
            "correct": false
        },
        {
            "body": "With let keyword",
            "correct": true
        },
        {
            "body": "With const keyword",
            "correct": true
        }
    ]
}
```

- DELETE /api/v1/question/{id} - Delete question by id


## A few notes about the implementation

- By default, the API starts with authentication disabled. This means you can CRUD questions without any setup!

- Authentication can be enabled with 2 environment variables. Read on about that below.

- JWT auth. works with any generated token, but there is a field in the payload called `id` which is important.

- The generated JWT token should be put as a `Authorization` header with a value of `Bearer insert_token_here`

- A user authenticated with JWT will live in their own world in Question API - they can only work with their own data.

- Because of the ordering requirements and difficulty to track the options - whether they are deleted, updated, created, I decided to delete them when updating a question and recreate them from the user request.

## What could be done better, IMO

- More logging. Not really sure the best approach for logging in Go projects. You have fmt, log.Fatal and probably 10 more options :)) I feel like I'm in JS land again :O

- Tests! Most of the stuff is trivial to test manually, but I feel this project can become complex and it is good to be reassured in the basic functionality before adding more features or, heaven forbid, modify existing features!

- More granular DAL methods. Currently there is some business logic there, but as it is part of a transaction, I felt it belongs there.

- Full-fledged authentication - login and register. I did not have the muscle this time to pull it of in 48 hours, but I will continue tinkering with it in the future! :)

## Env. variables

- `API_PORT` - the port on which the api will run

- `AUTH_ENABLED` - set it to anything to enable auth. To disable it again in the same CLI session, you can use `unset AUTH_ENABLED`.

- `JWT_SECRET` - Secret used for signing and decoding JWT. Set this if you have `AUTH_ENABLED` set.

## About the JWT authentication

The API is in the process of implementing JWT.

There is a user model and a relation between question and user, but it allows for users (just in the form of JWT tokens) that do not exist in DB to create questions.

Payload:
```
{
  "sub": "1234567890",
  "name": "John Doe",
  "id": "1a1e7ee2-f3ee-4d3e-9749-5b3d4cd7c729", <-- important!
  "iat": 1516239022
}
```

This field `id` will be used to associate a question with a user.

There is one default user created called `admin` used as a super user for the state in which auth is disabled.

Headers:
```
{
  "alg": "HS256",
  "typ": "JWT"
}
```

Secrets should not be in source code or docs, but here is one dummy secret for testing:

`qwertyuiopasdfghjklzxcvbnm123456`

To test the API with JWT auth, generate a JWT token with a field `id` in the payload and some secret.

The secret has to be defined as env. variable `JWT_SECRET` too.

A user will be able to CRUD only their questions.

