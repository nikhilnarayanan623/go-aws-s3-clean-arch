# Go AWS S3 File Upload and Download - Clean Architecture Project
This project is a simple application built using Go programming language that demonstrates file upload and download functionalities using Amazon S3. The project follows a clean code architecture, promoting modularity and maintainability.

# Project Overview
The go-aws-s3-clean-arch project showcases how to upload and download files to/from Amazon S3 using a clean architecture approach. It emphasizes separation of concerns and modularity, making it easier to expand and maintain the application.

# Used Packages
The project utilizes the following packages:
1. [AWS SDK](https://github.com/aws/aws-sdk-go): A comprehensive SDK for integrating Go applications with Amazon Web Services, providing functionalities for interacting with Amazon S3 and other AWS services.
2. [GIN](github.com/gin-gonic/gin): A web framework written in Go that combines high performance with an API similar to Martini.
3. [JWT](github.com/golang-jwt/jwt): A Go implementation of JSON Web Tokens for secure authentication and authorization.
4. [GORM](https://gorm.io/index.html) with [PostgreSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL): A powerful ORM library for Go that simplifies database operations.
5. [Wire](https://github.com/google/wire): A code generation tool for dependency injection, making it easier to connect components.
6. [Viper](https://github.com/spf13/viper): A configuration solution for Go applications, supporting various formats and 12-Factor app principles.
7. [swag](https://github.com/swaggo/swag) with [gin-swagger](https://github.com/swaggo/gin-swagger) and [swaggo files](github.com/swaggo/files): Converts Go annotations to Swagger Documentation 2.0 for API documentation.

Please refer to the respective package documentation for more information on their usage and integration.
Setup Instructions

To use and test the go-aws-s3-clean-arch project, follow these steps:
Prerequisites

Ensure the following software is installed on your system:

    Go (https://golang.org/doc/install)
    PostgreSQL (https://www.postgresql.org/download/)
    AWS account and S3 bucket (https://aws.amazon.com/s3/)

## 1. Clone the Repository

Clone the go-aws-s3-clean-arch repository to your local system:
```
git clone https://github.com/your-username/go-aws-s3-clean-arch.git
cd go-aws-s3-clean-arch
```
## 2. Install Dependencies

Install the required dependencies using either the provided Makefile command or Go's built-in module management:

```
# Using Makefile
make deps
# OR using Go
go mod tidy
```
## 3. Configure Environment Variables

Create a .env file in the project's root directory and set the following variables:
```.env
# Database
DB_HOST="your-database-host"
DB_NAME="your-database-name"
DB_USER="your-database-user"
DB_PASSWORD="your-database-password"
DB_PORT="your-database-port"

# AWS S3
AWS_ACCESS_KEY_ID="your-aws-access-key-id"
AWS_SECRET_ACCESS_KEY="your-aws-secret-access-key"
AWS_REGION="your-aws-region"
AWS_BUCKET_NAME="your-s3-bucket-name"

# Server
PORT="server-port"
```
## 4. Generate Wire Injectors

Run the following command to generate the Wire injectors:
```
make wire
```
## 5. Run the Application

```
make run
```
