# Tiger Kittens

## GQL Query

### Login

```
mutation {
  login(
    input: {username: "sample-mail@gmail.com", password: "Test1234"}
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
    input: {email: "sample-mail@gmail.com", username: "sample-username", password: "Test1234"}
  ) {
    message
    timestamp
  }
}
```

### Create Tiger

```
mutation CreateTiger($photo: Upload!) {
  createTiger(
    input: {
      dateOfBirth:"2024-01-01",
      lastLat:-6.300428,
      lastLong:107.167497,
      lastPhoto:$photo,
      lastSeen:"2024-05-01T00:00:00Z",
      name:"bengal tiger"
    }
  ) {
    message
    timestamp
  }
}

// Authorization: Bearer <access_token>
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

// Authorization: Bearer <access_token>
```

### List Sighting

```
{
  tigerSightings(input: {page: 1, perPage: 5}) {
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