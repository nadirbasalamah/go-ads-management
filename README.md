# go-ads-management

REST API application to manage advertisements. Written in Go with Echo Framework.

## Core Features

- Basic authentication.
- Advertisement management.
- Generate advertisement based on target audience, product and campaign platform.
- Review advertisement.

## Tech Stack

- Go
- MySQL
- Echo
- Pinata (File Storage)
- Github Actions
- Docker

## How to Use

1. Clone this repository.

2. Copy the configuration file.

```sh
cp .env.example .env
```

3. Fill the configuration inside `.env` file.

4. Generate the admin account.

```sh
go run helper/admin/generate.go
```

5. Run the application.

```sh
go run main.go
```

## Notes for Using with Docker

1. Make sure to set the `APP_MODE` in the `.env` into `production`.

2. Adjust the `DB_HOST` to use `mysql-service`.

3. Run the application.

```sh
docker compose up -d
```

4. Generate the admin account.

```sh
go run helper/admin/generate.go
```

5. Stop the application.

```sh
docker compose down
```

## Documentation

The application documentation is available [here](https://documenter.getpostman.com/view/5781191/2sAYBVgqzx).

## Additional Notes

1. In order to use recommendation features, make sure to insert the OpenAI API key.
2. This application uses [Pinata](https://pinata.cloud/) for file storage. Please insert the required credentials in the `.env` file.
