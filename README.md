# Questions API

## Env. variables

`API_PORT` - the port on which the api will run
`AUTH_ENABLED` - set it to anything to enable auth.
`JWT_SECRET` - Secret used for signing and decoding JWT. Set this if you have `AUTH_ENABLED` set.

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

To test the API with JWT auth, generate a JWT token with a field `id` in the payload and some secret. The secret has to be defined as env. variable too.

A user will be able to CRUD only their questions.