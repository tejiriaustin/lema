# Lema - User Management System

A full-stack user management system built with Go and React.

## Features

- User management with address information
- Post creation and management per user 
- Paginated data tables
- RESTful API
- Responsive UI with Tailwind CSS

## Tech Stack

**Backend:**
- Go 1.21+
- GORM
- Gin
- SQLite
- Zap Logger

**Frontend:**
- React 18
- TypeScript
- React Query
- Tailwind CSS
- Vite

## API Endpoints
```
POST    /users              // Create user
GET    /users              // Paginated user list
GET    /users/:id          // Single user with address
GET    /posts?userId=:id   // User's posts
POST   /posts              // Create post
DELETE /posts/:id          // Delete post
```

## Installation

1. Clone repository:
```bash
git clone https://github.com/tejiriaustin/lema.git
```

2. Backend Setup
 - Set .env file with the following variables


3. Database Schema

- Users (id, name, email, created_at, updated_at, version)
- Addresses (id, user_id, street, city, state, zipcode)
- Posts (id, user_id, title, body, created_at, updated_at)

4. Contributing

- Fork repository
- Create feature branch
- Commit changes
- Push to branch
- Create pull request
