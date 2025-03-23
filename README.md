# Financial Assistance Scheme Management System

## Project Overview
This project is a Financial Assistance Scheme Management System built using **Go (Golang)**, **Gin** web framework, and **PostgreSQL** for database management. The system allows users to manage applicants, schemes, and applications with features like registration, updating, retrieval, and deletion of data.

## Tech Stack
- **Go (Golang)**
- **Gin** (for API routing)
- **PostgreSQL** (for database management)
- **GORM** (ORM for database interaction)

## Setup Instructions

### 1. Clone the Repository
```sh
git clone https://github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System.git
cd Financial-Assistance-Scheme-Management-System
```

### 2. Configure Environment Variables
Create a `.env` file in the root directory with the following values:

**.env.example**
```
PORT=8080
DB_HOST=localhost
DB_USER=postgres
DB_PASS=your_postgres_password
DB_NAME=mydb
DB_PORT=5432
```

Alternatively, you can copy `.env.example` as a template:
```sh
cp .env.example .env
```

### 3. Install Dependencies
```sh
go mod tidy
```

### 4. Database Setup
Ensure you have PostgreSQL installed and create the database.
```sql
CREATE DATABASE mydb;
```

### 5. Run Database Migrations
Migrations are handled automatically when the server starts. The system will auto-migrate the following models:
- Applicant
- Scheme
- Application
- Household Members

### 6. Run the Application
```sh
go run cmd/main.go  
```

## API Documentation

### Applicants
- **Create an Applicant**
  - **POST** `/api/applicants`
  - **Body:**
```json
{
    "name": "Mary",
    "employment_status": "unemployed",
    "sex": "female",
    "date_of_birth": "1984-10-06",
    "household": [
        {
            "name": "Gwen",
            "employment_status": "unemployed",
            "sex": "female",
            "date_of_birth": "2016-02-01",
            "relation": "daughter",
            "school_level": 2
        }
    ]
}
```

- **Get All Applicants**
  - **GET** `/api/applicants`

- **Get an Applicant by ID**
  - **GET** `/api/applicants/:id`

- **Update an Applicant**
  - **PUT** `/api/applicants/:id`
  - **Body:** Same format as Create Applicant

- **Delete an Applicant**
  - **DELETE** `/api/applicants/:id`

### Schemes
- **Create a Scheme**
  - **POST** `/api/schemes`
  - **Body:**
```json
{
  "name": "Retrenchment Assistance Scheme",
  "criteria": {
    "employment_status": "unemployed"
  },
  "benefits": [
    {
      "name": "SkillsFuture Credits",
      "amount": 500.00
    }
  ]
}
```

- **Get All Schemes**
  - **GET** `/api/schemes`

- **Get a Scheme by ID**
  - **GET** `/api/schemes/:id`

- **Get Eligible Schemes**
  - **GET** `/api/schemes/eligible/:applicantID`

- **Update a Scheme**
  - **PUT** `/api/schemes/:id`
  - **Body:** Same format as Create Scheme

- **Delete a Scheme**
  - **DELETE** `/api/schemes/:id`

### Applications
- **Register an Application**
  - **POST** `/api/applications`
  - **Body:**
```json
{
  "applicant_id": "<applicant_id>",
  "scheme_id": "<scheme_id>"
}
```

- **Get Applications**
  - **GET** `/api/applications` (List all applications)

- **Get Applications by ID**
  - **GET** `/api/applications/:id`

- **Update an Application**
  - **PUT** `/api/applications/:id`

- **Delete an Application**
  - **DELETE** `/api/applications/:id`

- **Delete Applications by Applicant ID**
  - **DELETE** `/api/applications/applicant/:applicant_id`

## Error Handling with ErrorMiddleware
This project uses middleware for unified error handling:

- `400 Bad Request` for validation issues
- `404 Not Found` for missing resources
- `500 Internal Server Error` for unexpected issues

## Testing Instructions
You can test the endpoints using tools such as **Postman** or **Thunder Client** in Visual Studio Code.

## Deployment
Currently, there is no automated deployment setup. For local testing, follow the above steps. 


