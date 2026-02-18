
# Pedalea
**Pedalea** is a Go written REST API, that simulates a bike rental service.

## Usage
The following commands are useful when working with the app:

```bash
# Run app
Make run
# Builds swagger file
make docs
# Makes mocks for testint
make mocks
# Runs tests
make test
# Creates the db and runs all migrations
make db
```

## Env
The app uses `godotenv` to load environment variables from a `.env` file.

```
DB_FILE=./pedalea.db
JWT_SECRET=secret
ADMIN_CREDENTIALS=secret
```
## Docs
Swagger documentation is available in the `/docs` directory generated using the `make docs` command.
