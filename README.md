# id-generator
API for id generation written in Go


### Build the image
docker build --tag docker-id-generator .

### Tag images
docker image tag docker-id-generator:latest docker-id-generator:v1.0


The application starts at port 8080:

    GET /v1/ping Health check endpoint, returns 'pong' message



## Getting Started

# Clone Project
git clone https://github.com/oussaka/id-generator.git id-generator

# Change Directory
cd id-generator

Using Docker

# Build & Create Docker Containers
docker-compose up -d

Using Local Environment

# Copy Example Env file
cp ./env.example .env

# Change MongoDB URI and Database Name

# MONGO_URI=<mongo_uri>
# MONGO_DATABASE=<db_name>

# Download Modules
go mod download

# Build Project
go build -o go-starter id-generator

# Run the Project
./id-generator

The application starts at port 8080:

    GET /v1/ping Health check endpoint, returns 'pong' message

    POST /v1/auth/register Creates a user and tokens
    POST /v1/auth/refresh Refresh expired tokens
    POST /v1/auth/login Login a user

    POST /v1/users Create a new user
    GET /v1/users Get paginated list of users
    GET /v1/users/:id Get a one user details
    PUT /v1/users/:id Update a user
    DELETE /v1/users/:id Delete a user

## commands

### import new users from external API

``` ./cmd/import/importer import -startDate=2024-12-01 -color
