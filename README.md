# Questions API

Hi! :)
Here is what I did for the time being.

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


## A few notes about the implementation

- By default, the API starts with authentication disabled. This means you can CRUD questions without any setup!

- Authentication can be enabled with 2 environment variables. Read on about that below.

- JWT auth. works with any generated token, but there is a field in the payload called `id` which is important.

- The generated JWT token should be put as a `Authorization` header with a value of `Bearer insert_token_here`

- A user authenticated with JWT will live in their own world in Question API - they can only work with their own data.

- Because of the ordering requirements and difficulty to track the options - whether they are deleted, updated, created, I decided to delete them when updating a question and recreate them from the user request.

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

