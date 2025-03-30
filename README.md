# InternShip Task

The task for QuantStrategix Intership

## How To Run

#### Clone the repo

```sh
$ git clone
$ cd
```

#### Download Dependencies

```sh
$ go mod download
```

#### Fill Environment Variables

```sh
$ cp .env.example .env
```

#### Start Containers

```sh
$ docker compose up --build
```

## API Endpoints

[Postman Documentation]{https://documenter.getpostman.com/view/26244894/2sB2cPjkco}

| route                      | description                               |
| -------------------------- | ----------------------------------------- |
| POST /register             | Register user                             |
| POST /login                | User login                                |
| POST /subjects             | Add subject                               |
| GET /subjects              | Get all subjects                          |
| POST /tutor/profile        | Add/Update tutor details                  |
| GET /tutor?subject=        | To get all tutor of that subject          |
| POST /bookings             | To book a tutor session                   |
| GET /bookings              | To get all the booked session of the user |
| PATCH /bookings/:id/status | To update the status of booking           |

## POST /register

### Request

```json
{
  "name": "tutor",
  "email": "tutor@gmail.com",
  "password": "1234",
  "role": "tutor"
}
```

### Response

```json
{
  "message": "User registered successful",
  "status": true
}
```

## POST /login

### request

```json
{
  "email": "tutor@gmail.com",
  "password": "1234"
}
```

### response

```json
{
  "message": "Successfully logged in",
  "status": true,
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDMyOTc2MDEsInJvbGUiOiJ0dXRvciIsInN1YiI6InR1dG9yQGdtYWlsLmNvbSJ9.AAAMJkHUprm8tlDEM_VMhL7ApHeDUWDcJXj-RL0hTcM"
  }
}
```

## POST /subjects

### request

```json
{
  "name": "DSA"
}
```

### response

```json
{
  "message": "Subject added",
  "status": true
}
```

## GET /subjects

### response

```json
{
  "message": "All subject retrived",
  "status": true,
  "data": [
    {
      "ID": "e298ed85-1530-48c4-85a0-3fcf51d3012a",
      "Name": "DSA"
    }
  ]
}
```

## POST /tutor/profile

### request

```json
{
  "bio": "testing",
  "subject_names": ["DSA"]
}
```

### response

```json
{
  "message": "Details added",
  "status": true
}
```

## GET /tutor?subject=DSA

### response

```json
{
  "message": "Got the tutor list",
  "status": true,
  "data": [
    {
      "UserID": "b17532a1-2ee2-4ba4-a727-012c2c612857",
      "Bio": "testing",
      "SubjectsName": ["DSA"]
    }
  ]
}
```

## POST /bookings

### request

```json
{
  "tutor_id": "b17532a1-2ee2-4ba4-a727-012c2c612857",
  "subject_id": "e298ed85-1530-48c4-85a0-3fcf51d3012a",
  "booking_time": "tuesday"
}
```

### response

```json
{
  "message": "Booking done",
  "status": true
}
```

## GET /bookings

### response

```json
{
  "message": "Booking fetched",
  "status": true,
  "data": [
    {
      "ID": "844088ce-cfe0-4871-94d9-f52f936fccc0",
      "StudentID": "f982dc6e-4966-466b-9a3f-9be6c0b277b7",
      "Student": {
        "ID": "f982dc6e-4966-466b-9a3f-9be6c0b277b7",
        "Name": "student",
        "Email": "student@gmail.com",
        "Password": "$2a$10$UQYreZOcbdIrG7ugryM9aOrZuNiOBkwy8vtEFOOd7an0TzYV/9KZS",
        "Role": "student"
      },
      "TutorID": "b17532a1-2ee2-4ba4-a727-012c2c612857",
      "Tutor": {
        "ID": "b17532a1-2ee2-4ba4-a727-012c2c612857",
        "Name": "tutor",
        "Email": "tutor@gmail.com",
        "Password": "$2a$10$SBoyzTsTSsoxghCWfl7xkew3ukvKzLOM9cPWGKixWeFvXxLIFBV4e",
        "Role": "tutor"
      },
      "SubjectID": "e298ed85-1530-48c4-85a0-3fcf51d3012a",
      "Subject": {
        "ID": "e298ed85-1530-48c4-85a0-3fcf51d3012a",
        "Name": "DSA"
      },
      "BookingTime": "tuesday",
      "Status": "pending"
    }
  ]
}
```

## PATCH /bookings/:id/status

### request

```json
{
  "status": "confirmed"
}
```

### response

```json
{
  "message": "Booking updated",
  "status": true
}
```
