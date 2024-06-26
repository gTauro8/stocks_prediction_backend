Financial Guru Web Application
This repository contains a web application built with Go (Golang) and Gin framework, designed for managing user financial questionnaires and providing investment recommendations.

Requirements
Before running the application, ensure you have the following installed:

Go: Version 1.16 or higher. Installation Guide
MongoDB: Version 4.0 or higher. Installation Guide
Postman: For testing API endpoints. Download Postman
Setup
Clone the repository:

bash
Copia codice
git clone https://github.com/yourusername/financial-guru.git
cd financial-guru
Set up environment variables:

Create a .env file in the root directory and add the following:

plaintext
Copia codice
JWT_SECRET=a4d1d2caca8527d49f147325c21ff7251fa6c31de1244136b9513c69ec592a52
Replace JWT_SECRET with your own secret key for JWT token generation.

Install dependencies:

bash
Copia codice
go mod tidy
Start MongoDB:

Make sure MongoDB is running locally on the default port 27017.

Run the application:

bash
Copia codice
go run main.go
The application will start running at http://localhost:8080.

Usage
Use Postman or any API testing tool to interact with the endpoints:

Register a new user:

POST http://localhost:8080/register
Body (JSON):
json
Copia codice
{
  "username": "yourusername",
  "password": "yourpassword"
}
Login with registered user:

POST http://localhost:8080/login
Body (JSON):
json
Copia codice
{
  "username": "yourusername",
  "password": "yourpassword"
}
Protected endpoints (require JWT token):

Include JWT token in the Authorization header: Bearer YOUR_JWT_TOKEN
Endpoints
Authentication
POST /register: Register a new user.
POST /login: Login with username and password to receive JWT token.
Questionnaire and Recommendations
POST /api/questionnaire: Save or update user's financial questionnaire responses.
POST /api/recommend: Get investment recommendations based on user's preferences.
POST /api/predict: Predict returns for recommended investments.
User Profile
GET /api/user/profile: Retrieve user's profile and financial questionnaire responses.
Contributing
Contributions are welcome! Feel free to open issues or pull requests for any improvements or feature requests.

