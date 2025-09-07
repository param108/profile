# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is "Profile" - an open source platform for blogging using tweets. It's a Twitter-like social media platform with both frontend and backend components. The live deployment is at https://ui.tribist.com.

## Repository Structure

- **`/api`** - Go backend API server (main application logic)
- **`/twitterlike`** - Next.js frontend React application 
- **`/dashboard`** - Additional dashboard component
- **`/docs`** - Documentation and contribution guidelines
- **`/deployment`** - Deployment configurations
- **`/dev_setup`** - Development environment setup

## Development Commands

### Backend API (Go)
```bash
cd api
make build        # Build the Go application
make test         # Run tests with linting check  
make dev          # Build and run development server with local DB
make db           # Create and migrate local database
make db-destroy   # Destroy local database
make migrate      # Run database migrations
make schema       # Update schema.sql from current DB state
```

### Frontend (Next.js)
```bash
cd twitterlike
npm run dev       # Start development server
npm run build     # Build for production
npm run start     # Start production server
npm run lint      # Run ESLint
npm run test      # Run Jest tests in watch mode
npm run citest    # Run all Jest tests (CI mode)
```

## Architecture

### Backend (API)
- **Language**: Go 1.19+
- **Framework**: Gorilla Mux for routing, GORM for database
- **Database**: PostgreSQL
- **Structure**: Modular packages for different entities (users, tweets, threads, etc.)
- **Entry Point**: `api/server/main.go` - CLI application with commands (serve, migrate)
- **Key Packages**:
  - `server/` - HTTP server and routing logic
  - `store/` - Database models and operations  
  - `users/`, `tweets/`, `threads/` - Domain-specific handlers
  - `utils/` - Shared utilities and authentication

### Frontend (Next.js)
- **Framework**: Next.js 13+ with App Router
- **Language**: TypeScript + React
- **Styling**: Tailwind CSS
- **Testing**: Jest with React Testing Library
- **Key Directories**:
  - `app/` - Next.js app router pages and components
  - `app/components/` - Reusable React components
  - `__test__/` - Jest test files

### Database
- PostgreSQL with GORM ORM
- Migrations in `api/db/migrations/`
- Schema maintained in `api/db/schema.sql`
- **Important**: Always update schema.sql when adding migrations

## Environment Setup

### API Environment Files
- `dev.env` - Local development
- `test.env` - Testing
- `ci.env` - CI/CD
- Copy `env.sample` for reference

### Required Tools
- Go 1.19+
- Node.js/npm
- PostgreSQL
- goimports (for Go formatting)

## Development Workflow

1. **Database Changes**: When adding DB migrations, always run `make schema` to update `schema.sql` before creating PRs
2. **Testing**: Backend tests require `goimports` formatting compliance
3. **CI/CD**: GitHub Actions handle testing and deployment for both API and frontend
4. **Versioning**: Both components generate version files from git branch/commit info

## Key Implementation Notes

- The API serves as a CLI tool with multiple commands (serve, migrate, etc.)
- Frontend uses Next.js App Router pattern
- Authentication handled via JWT in the backend
- The platform supports threads (longer posts) and regular tweets
- User management includes groups and permissions
- to build backend do `cd api; make build`