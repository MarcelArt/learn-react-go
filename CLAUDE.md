# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Learning Management System (LMS) for Elementary School to College, built with:
- **Backend**: Go with Fiber framework, GORM ORM, PostgreSQL database
- **Frontend**: React + TypeScript in the `./web` directory

## Architecture

### Backend (Go)
- Uses **Repository Pattern** with clean architecture
- **JWT Authentication** with access tokens (5 min) and refresh tokens (1-30 days)
- **Fiber** web framework with Swagger documentation
- **GORM** for database operations
- **Polygo CLI** for scaffolding (project was generated using polygo)

### Frontend (React)
- **React 19** with TypeScript
- **TanStack Router** for routing
- **Tailwind CSS** + **Shadcn** components
- **Vite** as build tool

## Development Commands

### Backend
```bash
# Run development server with hot reload
make dev

# Generate Swagger documentation
make swag

# Run the application
make go

# Database migrations
go run main.go migrate up    # Create/update tables
go run main.go migrate down  # Drop tables

# Start server only
go run main.go serve
```

### Frontend (in ./web directory)
```bash
# Start development server
npm run dev

# Build for production
npm run build

# Run tests
npm run test

# Preview production build
npm run serve
```

## Project Structure

### Backend Structure
- `models/` - Data models with GORM tags and DTOs
- `repositories/` - Data access layer with interfaces
- `handlers/api/` - HTTP handlers (REST API)
- `routes/` - Route definitions
- `middlewares/` - Custom middleware (auth, etc.)
- `config/` - Environment configuration
- `utils/` - Utility functions (JWT, error handling)
- `cmd/` - Command management (serve, migrate)

### Frontend Structure
- `web/src/` - React source code
- `web/src/routes/` - Route components (TanStack Router)
- `web/src/components/` - Reusable components
- `web/public/` - Static assets

## Key Technical Details

### Authentication System
- JWT-based with access/refresh token pairs
- Access token expires in 5 minutes
- Refresh token: 1 day (normal) or 30 days (remember me)
- Device tracking for security

### Database
- PostgreSQL with GORM
- Multi-tenancy support (schools)
- Schema includes: Users, Roles, Permissions, Courses, Assignments, etc.

### Scaffolding with Polygo
The project uses polygo CLI for scaffolding new endpoints:
```bash
# Install polygo
go install github.com/MarcelArt/polygo@latest

# Scaffold new model with CRUD
polygo add ModelName
```

## Environment Variables

Backend requires `.env` file with:
- `PORT` - Server port (default: 8080)
- `DB_*` - Database connection details
- `JWT_SECRET` - JWT signing secret
- `SERVER_ENV` - Environment (development/production)

## API Documentation

Swagger UI available at: `http://localhost:8080/swagger/index.html`

## Testing

- Backend: No specific test runner configured (add as needed)
- Frontend: Vitest with React Testing Library

## Task Log

### Completed Tasks (2025-09-27)

**User Management System Implementation**
- ✅ Created School model and API endpoints
- ✅ Created Role model and API endpoints
- ✅ Created Permission model and API endpoints
- ✅ Created UserRole and RolePermission relationship models
- ✅ Updated User model to include school relationship
- ✅ Created multi-tenancy middleware
- ✅ Created authentication middleware with role-based access
- ✅ Created school registration flow
- ✅ Tested the user management system

**Features Implemented:**
- Multi-tenancy architecture with school data isolation
- Role-Based Access Control (RBAC) with many-to-many relationships
- JWT authentication with access/refresh tokens
- School registration with automatic default role creation
- Database migrations for all models
- Comprehensive API endpoints with proper authentication
- Transaction-based school registration flow

**Testing Results:**
- School registration working correctly
- User authentication and token refresh functional
- Multi-tenancy data isolation verified
- User creation and management working
- Database migrations successful

---

## Task Logging Note

**IMPORTANT**: Whenever I complete a task requested by the user, I must add a detailed entry to the Task Log section above. The entry should include:
- Date of completion
- Brief description of the task
- Key features implemented
- Testing results
- Any relevant technical details

This ensures a complete record of all development work and helps maintain project context.