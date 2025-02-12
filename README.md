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

## Feedback

@⁨Ten⁩ your feedback is ready

*Backend Review*
- Incomplete/incorrect graceful shutdown implementation, which could cause issues during process termination. - ✅
- Errors during database initialization and migration are simply returned without logging, making it harder to diagnose issues - ✅
- Missing /users/count endpoint for retrieving the user count. - ✅
- Incorrect use of JSON binding for the GET /users/{id} endpoint, leading to potential data binding issues. - ✅
- Empty GetPostsRequest struct, which may cause confusion or errors when handling requests. - ✅
- The GetUserCount method is not exposed via an API endpoint, limiting its accessibility. - ✅
- Duplicate logging key in structured logging middleware, which could create confusion in logs. - ✅
- No CI/CD configuration files present, preventing automated testing and deployment processes. - ✅
- No authentication or secure error reporting evident, which could compromise security. - ✅
- No rate limiting middleware, leaving the API vulnerable to abuse or overload. - ✅
- No tests are provided, which limits the ability to ensure the functionality and stability of the application. - ✅
- Several endpoints and production-readiness features are missing, which could affect the application's robustness in a live environment.

*Frontend Review*
- Error handling and validation were not done gracefully. Creating a user with unexpected fields did not trigger error or success messages, leaving users unaware of the action's outcome.
- The form was not cleared after a user was created, leading to a poor user experience as previous input remained visible. - ✅
- The delete functionality is not working, which affects the application's core functionality. - ✅
- Spacing is inconsistent across the interface, and the font on the user listing page is too large, affecting visual clarity. - 
-  Responsiveness is perfect, adapting well to different screen sizes. - ✅
- Pages were broken down into reusable components, enhancing code maintainability. - ✅
- Pagination is not functioning properly, hindering smooth data navigation. - ✅
- Give the title and body of the post a word limit, or simply break longer words with an elipsis(....) as longer words ruins the responsiveness of the code when tested.