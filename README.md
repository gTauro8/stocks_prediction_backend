# Financial Guru Web Application

This repository contains a web application built with Go (Golang) and Gin framework, designed for managing user financial questionnaires and providing investment recommendations.

## Requirements

Before running the application, ensure you have the following installed:

- **Go**: Version 1.16 or higher. [Installation Guide](https://golang.org/doc/install)
- **MongoDB**: Version 4.0 or higher. [Installation Guide](https://docs.mongodb.com/manual/installation)
- **Postman**: For testing API endpoints. [Download Postman](https://www.postman.com/downloads/)

## Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/financial-guru.git
   cd financial-guru

2. ** Set up environment variables:
    Create a json (Config.json) file in the root directory and add the following:
     ```JSON
       {
          "JWT_SECRET": "a4d1d2caca8527d49f147325c21ff7251fa6c31de1244136b9513c69ec592a52"
       }
     ```
   
     ```bash
     node -e "console.log(require('crypto').randomBytes(32).toString('hex'))"
     
3. ** Install dependencies:
   ```bash
   go mod tidy

3. ** Start MongoDB
   Make sure MongoDB is running locally on the default port 27017.


##Usage

Use Postman or any API testing tool to interact with the endpoints:

- ** Register a new user:

POST http://localhost:8080/register
Body:
  ```JSON
      {
        "username": "yourusername",
        "password": "yourpassword"
      }
 ```
 
  - ** Login user:

POST http://localhost:8080/login
Body:
  ```JSON
      {
        "username": "yourusername",
        "password": "yourpassword"
      }
  ```

Response: auth_token


      





