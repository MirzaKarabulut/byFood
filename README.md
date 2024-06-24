# ByFood API

This is a simple API for a book store, developed using the Go programming language and the Gin web framework. It provides endpoints for managing books and processing URLs.

## Table of Contents

- [Setup Instructions](#setup-instructions)
- [Project Structure](#project-structure)
- [Endpoint Usage](#endpoint-usage)

## Setup Instructions

### Prerequisites

- Go
- PostgreSQL

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
2. Navigate to the project directory:
  cd ByFood

3. Install dependencies:
go mod tidy

4. Run the application:
   go run main.go

5. Navigate to the frontend directory:
   cd frontend

6. Install Next.js dependencies:
   npm install next@latest react@latest react-dom@latest

8. Run the frontend application:
   npm run dev


## Project Structure

byFood/

 - backend/
 
   - controllers/       # Business logic handlers
   
   - models/            # Database models
   
   - migrate/           # Database migration files
   
   - docs/              # API documentation files
   
   - initializers/      # Setup scripts (e.g., database connection)
   
   - go.mod             # Go module definitions
   
   - go.sum             # Go module checksums
   
   - main.go            # Entry point of the backend application
   
  
 - frontend/
 
   - public/            # Static files like `index.html`, images, etc.
   
   - src/               # Contains components, pages, and styles.
   
      - components/    # React components
      
      - pages/         # React pages
      
      - styles/        # CSS and styling files
      
   - package.json       # Next.js dependencies and scripts
   
 .gitignore             # Specifies intentionally untracked files to ignore
 
 README.md              # Project overview and setup instructions
 


  ## Endpoint Usage

# Books
- GET /books : 
  Retrieve all books.
- POST /books : 
  Create a new book. Requires a JSON body.
- GET /books/id : 
  Retrieve a book by its ID.
- PUT /books/id :
  Update a book by its ID. Requires a JSON body.
- DELETE /books/id :
  Delete a book by its ID.
  
# URL Processing
- POST /process-url:
  Process a given URL. Requires a JSON body with the URL.
  
# Documentation
- GET /docs/index.html :
 Access the Swagger UI for API documentation.
 
   
