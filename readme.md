# Tiger Kittens

## Setup



## REST API

Postman collection can be found in the root of this project in file named `tiger-kittens.postman_collection`

### List of available endpoint:

### Register `POST /v1/user/register`

Create new user

### Login `POST /v1/user/login`

Login to existing user

### Create Tiger `POST /v1/tiger`

Create new tiger

### List Tiger `GET /v1/tiger?page=1&per_page=5`

List all tigers according to pagination

### Create Sighting `POST /v1/tiger/{tigerID}sighting`

Create new sighting, will update lastLat, lastLong, lastSeen and lastPhoto in related tigerID

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