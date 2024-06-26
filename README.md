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


## Usage

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


## Endpoints

** Authentication
- POST /register: Register a new user.
- POST /login: Login with username and password to receive JWT token.

  
** Questionnaire and Recommendations 
- POST /api/questionnaire: Save or update user's financial questionnaire responses.
- POST /api/recommend: Get investment recommendations based on user's preferences.
- POST /api/predict: Predict returns for recommended investments.
  
** User Profile
- GET /api/user/profile: Retrieve user's profile and financial questionnaire responses.

  All /api/ methods neded Autorization Berear token


## Contributing
- Contributions are welcome! Feel free to open issues or pull requests for any improvements or feature requests.

      





