# Tiger Kittens

## How to run

1. This requires docker & docker-compose to be installed, and port 9000 to be available
2. Run `make db-up` to spawn postgres container
3. Run `make db-migrate` to sync schema migration to latest version
4. Run `make api-up` to build and run API service
5. For development, run `make serve` instead to run directly from code, for development use go version 1.21.10

## Email SMTP Setup

To enable sending email there is some env that need to be configured:
```
// .env.local
EMAIL_ENABLED=false # set to true, if you want to enable sending through google email
EMAIL_APP_PASSWORD= # generate app password from your google email in this link https://myaccount.google.com/apppasswords
EMAIL_SENDER= # put your google email here, ex: my-email@gmail.com
EMAIL_SERVER_ADDR=smtp.gmail.com:587
EMAIL_SERVER_HOST=smtp.gmail.com
```

This email feature is tested using GMAIL only during development.

## REST API

Postman collection can be found in the root of this project in file named `tiger-kittens.postman_collection`

### List of available endpoint:

### Register `POST /v1/user/register`

Create new user
```
Validation:
- email is required, must be an email, max 64 char
- password is required, between 8-20 char, must contain lower case, upper case and number
- username is required, between 6-64 char, only alphanum
```

### Login `POST /v1/user/login`

Login to existing user
```
Validation:
- password is required, between 8-20 char, must contain lower case, upper case and number
- username is required, between 6-64 char (can be filled with email)
```

### Create Tiger `POST /v1/tiger`

Create new tiger
```
Validation:
- lastLat is required, between -90 - 90
- lastLong is required, between -180 - 180
- lastSeen is required, RFC3339 or ISO8601
- lastPhoto is required, can only handle jpeg and png
- date_of_birth is required, date only (2006-01-02)
- name is required, between 3-64 char
- need HTTP Header 'Authorization: Bearer <access_token>'
```

### List Tiger `GET /v1/tiger?page=1&per_page=5`

List all tigers according to pagination

### Create Sighting `POST /v1/tiger/{tigerID}/sighting`

Create new sighting, will update lastLat, lastLong, lastSeen and lastPhoto in related tigerID
```
Validation:
- lat is required, between -90 - 90
- long is required, between -180 - 180
- photo is required, can only handle jpeg and png
- can only add new sighting if range from previous sighting > 5km, this is calculated using Haversine formula from lat & long value
- need HTTP Header 'Authorization: Bearer <access_token>'
```

### List Sighting `GET /v1/tiger/{tigerID}/sighting`

List all sighting of 1 tiger according to pagination

## GQL Query

GQL query can be accessed in `http://localhost:9000/gql/query`

Playground can be accessed in `http://localhost:9000/gql/playground`, but for upload files it is recommended to use Altair GQL Client https://altairgraphql.dev/  

### Login

```
mutation {
  login(
    input: {
      username: "sample-mail@gmail.com", 
      password: "Test1234"
    }
  ) {
    accessToken
    timestamp
  }
}
```

### Register

```
mutation {
  register(
    input: {
      email: "sample-mail@gmail.com", 
      username: "sample-username", 
      password: "Test1234"
    }
  ) {
    message
    timestamp
  }
}
```

### Create Tiger

```
// Authorization: Bearer <access_token>
mutation CreateTiger($photo: Upload!) {
  createTiger(
    input: {
      dateOfBirth:"2024-01-01",
      lastLat:-6.300428,
      lastLong:107.167497,
      lastPhoto:$photo,
      lastSeen:"2024-05-01T00:00:00Z",
      name:"sample-tiger"
    }
  ) {
    message
    timestamp
  }
}
```

### List Tiger

```
{
  tigers(input: {page: 1, perPage: 5}) {
    data {
      ID
      dateOfBirth
      lastLat
      lastLong
      lastSeen
      lastPhoto
      name
      createdAt
      updatedAt
    }
    pagination {
      page
      perPage
      totalPage
      totalItem
    }
  }
}
```

### Create Sighting

```
// Authorization: Bearer <access_token>
mutation CreateSighting($photo: Upload!) {
  createSighting(
    input: {
      tigerID:"6",
      lat:-6.400428,
      long:107.167497,
      photo:$photo
    }
  ) {
    message
    timestamp
  }
}
```

### List Sighting

```
{
  tigerSightings(input: {tigerID: "3", page: 1, perPage: 5}) {
    data {
      uploadedBy
      tigerName
      photo
      lat
      long
      createdAt
    }
    pagination {
      page
      perPage
      totalPage
      totalItem
    }
  }
}
```
