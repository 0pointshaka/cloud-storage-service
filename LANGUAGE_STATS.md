# Language Statistics & Architecture

Repository: **0pointshaka/cloud-storage-service**  
Repository ID: **1353019409**

## 📊 Language Composition

This is a **full-stack application** built with two primary languages:

### Backend: Go (🟦 Go)
- **Purpose**: RESTful API server, database management, file operations
- **Percentage**: ~45-50% of total codebase
- **Key Files**:
  - `backend/main.go` - Application entry point
  - `backend/config/` - Database and environment setup
  - `backend/models/` - Data models (User, File, SharedFile)
  - `backend/handlers/` - HTTP request handlers
  - `backend/routes/` - API route definitions
  - `backend/middleware/` - Authentication middleware

### Frontend: JavaScript/React (🟨 JavaScript)
- **Purpose**: User interface, client-side logic, API communication
- **Percentage**: ~45-50% of total codebase
- **Key Files**:
  - `frontend/src/App.jsx` - Main React component
  - `frontend/src/pages/` - Page components
  - `frontend/src/styles/` - CSS styling
  - `frontend/vite.config.js` - Vite bundler configuration

### Configuration & Scripts: Other (~5%)
- **Batch Scripts** (Windows): `.bat` files for automation
- **Configuration Files**: `.gitignore`, `go.mod`, `package.json`
- **Documentation**: `README.md`, `CONTRIBUTING.md`, `LANGUAGE_STATS.md`

## 🏗 Architecture Overview

```
┌─────────────────────────────────────────────────────┐
│                  Browser / Client                    │
└──────────────────────┬──────────────────────────────┘
                       │
                       │ HTTP/REST
                       │
┌──────────────────────▼──────────────────────────────┐
│           Frontend (React + Vite)                    │
│  ┌────────────────────────────────────────────────┐ │
│  │  Components                                    │ │
│  │  - LoginPage                                  │ │
│  │  - DashboardPage                              │ │
│  │  - FileUpload                                 │ │
│  │  - FileList                                   │ │
│  │  - ShareDialog                                │ │
│  └────────────────────────────────────────────────┘ │
│  ┌────────────────────────────────────────────────┐ │
│  │  Services (Axios)                              │ │
│  │  - API calls to backend                       │ │
│  │  - Error handling                             │ │
│  │  - Request/Response interceptors              │ │
│  └────────────────────────────────────────────────┘ │
└──────────────────────┬──────────────────────────────┘
                       │
                       │ HTTP/REST (port 8080)
                       │
┌──────────────────────▼──────────────────────────────┐
│            Backend (Go + Gin Framework)              │
│  ┌────────────────────────────────────────────────┐ │
│  │  Handlers (HTTP Endpoints)                     │ │
│  │  - /api/auth/* - Authentication               │ │
│  │  - /api/files/* - File management             │ │
│  │  - /api/share/* - File sharing                │ │
│  └────────────────────────────────────────────────┘ │
│  ┌────────────────────────────────────────────────┐ │
│  │  Middleware                                    │ │
│  │  - JWT authentication                         │ │
│  │  - CORS handling                              │ │
│  │  - Error logging                              │ │
│  └────────────────────────────────────────────────┘ │
│  ┌────────────────────────────────────────────────┐ │
│  │  Models (GORM)                                 │ │
│  │  - User                                        │ │
│  │  - File                                        │ │
│  │  - SharedFile                                 │ │
│  └────────────────────────────────────────────────┘ │
└──────────────────────┬──────────────────────────────┘
                       │
                       │ SQL queries
                       │
┌──────────────────────▼──────────────────────────────┐
│              Database (SQLite)                       │
│  ┌────────────────────────────────────────────────┐ │
│  │  Tables                                        │ │
│  │  - users (id, username, email, password)      │ │
│  │  - files (id, user_id, name, path, size)      │ │
│  │  - shared_files (id, file_id, shared_with_id) │ │
│  └────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────┘
                       │
                       │ File I/O
                       │
┌──────────────────────▼──────────────────────────────┐
│            File Storage (uploads/)                   │
└─────────────────────────────────────────────────────┘
```

## 🔧 Technology Stack by Component

### Backend (Go)

| Component | Technology | Purpose |
|-----------|-----------|---------|
| **Framework** | Gin | Fast HTTP web framework |
| **Database** | SQLite + GORM | Data persistence |
| **Authentication** | JWT + bcrypt | User auth & password security |
| **File I/O** | io/os | File upload/download handling |
| **Logging** | Standard log | Error and info logging |

### Frontend (JavaScript/React)

| Component | Technology | Purpose |
|-----------|-----------|---------|
| **UI Framework** | React 18 | User interface |
| **Build Tool** | Vite | Fast development & production build |
| **HTTP Client** | Axios | API communication |
| **Styling** | CSS3 | Component styling |
| **State** | React Hooks | Component state management |

### Development & Deployment

| Component | Technology | Purpose |
|-----------|-----------|---------|
| **Windows Automation** | Batch Scripts | Build, run, setup automation |
| **Version Control** | Git/GitHub | Code management |
| **Package Manager** | npm | Frontend dependencies |
| **Module System** | Go Modules | Backend dependencies |

## 💾 Database Schema

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE files (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    path TEXT NOT NULL,
    size INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE shared_files (
    id INTEGER PRIMARY KEY,
    file_id INTEGER NOT NULL,
    shared_by_id INTEGER NOT NULL,
    shared_with_id INTEGER NOT NULL,
    created_at TIMESTAMP,
    FOREIGN KEY (file_id) REFERENCES files(id),
    FOREIGN KEY (shared_by_id) REFERENCES users(id),
    FOREIGN KEY (shared_with_id) REFERENCES users(id)
);
```

## 📡 API Flow Example: File Upload

```
1. Frontend (React)
   └─> User selects file
   └─> FileUpload component prepares FormData
   └─> Axios sends POST request to /api/files/upload
   └─> Includes JWT token in Authorization header

2. Backend (Go)
   └─> Gin receives request
   └─> AuthMiddleware validates JWT token
   └─> FileHandler processes upload
   └─> Validates file
   └─> Saves to disk (/uploads/)
   └─> Creates File record in SQLite

3. Database (SQLite)
   └─> File entry created with user_id, name, path, size
   └─> Returns new file ID

4. Backend Response
   └─> Returns success with file metadata
   └─> Includes file ID, size, created_at

5. Frontend (React)
   └─> Receives response
   └─> Updates file list
   └─> Shows success message
```

## 📊 Code Statistics

### Backend (Go)
- **Total Lines**: ~1,500+ (estimated)
- **Main Packages**:
  - `main` - Entry point
  - `config` - Configuration
  - `models` - Data structures
  - `handlers` - Business logic
  - `routes` - Routing
  - `middleware` - Middleware functions
- **Dependencies**: Gin, GORM, JWT, bcrypt

### Frontend (React)
- **Total Lines**: ~1,500+ (estimated)
- **Components**: 5+ main components
- **Pages**: 2 main pages (Login, Dashboard)
- **Dependencies**: React, Axios, Vite

### Build & Deployment
- **Batch Scripts**: 4 files (setup, build, run, run-dev)
- **Configuration Files**: 5+ files

## 🚀 Development Workflow by Language

### Go Development
1. Edit `.go` files in `backend/` directory
2. Run `go fmt` to format code
3. Run tests: `go test ./...`
4. Run server: `go run main.go`
5. Changes require restart

### React Development
1. Edit `.jsx` files in `frontend/src/` directory
2. Vite automatically reloads on save (hot reload)
3. Run tests: `npm test`
4. Build: `npm run build`
5. Output to `frontend/dist/`

## 📦 Dependency Management

### Go Dependencies
```bash
cd backend
go mod download      # Download dependencies
go mod tidy         # Clean up unused dependencies
go mod graph        # View dependency graph
```

### JavaScript Dependencies
```bash
cd frontend
npm install         # Install dependencies
npm update         # Update dependencies
npm audit          # Check security vulnerabilities
```

## 🔒 Security Architecture

1. **Authentication**
   - JWT tokens issued on login
   - Token expires after 24 hours
   - Tokens validated on each protected endpoint

2. **Authorization**
   - Users can only access their own files
   - Sharing is explicit and verified
   - Database constraints enforce user associations

3. **Password Security**
   - Passwords hashed with bcrypt
   - Salt added automatically
   - Never stored in plain text

4. **Communication**
   - HTTP endpoints (HTTPS in production)
   - CORS configured for frontend origins
   - Request validation on all endpoints

## 🎯 Code Organization Best Practices

### Backend (Go)
- **Models**: Define database schemas
- **Handlers**: Implement business logic
- **Middleware**: Handle cross-cutting concerns
- **Routes**: Organize endpoints
- **Config**: Centralize configuration

### Frontend (React)
- **Components**: Reusable UI elements
- **Pages**: Full-page views
- **Services**: API communication layer
- **Styles**: Component-specific CSS
- **Hooks**: Custom logic hooks

## 📈 Scalability Considerations

**Current Limitations:**
- SQLite not suitable for production with high concurrency
- Files stored locally (single server)
- No caching layer

**For Production:**
- Migrate to PostgreSQL or MySQL
- Implement S3/Cloud storage
- Add Redis caching
- Use load balancer
- Implement CDN for frontend assets

## 🔄 Data Flow

### Authentication Flow
```
User Input → Frontend Form → Axios POST → Backend Handler 
→ Database Check → Password Verification → JWT Generation 
→ Response to Frontend → Store Token → Redirect to Dashboard
```

### File Upload Flow
```
File Selection → FormData Preparation → Axios POST with JWT 
→ Backend Auth Middleware → File Validation → Save to Disk 
→ Database Record Creation → Response with File Metadata 
→ Update Frontend State → Show Success Message
```

### File Sharing Flow
```
Select File → Enter Username → Axios POST → Backend Verification 
→ User Lookup → Create SharedFile Record → Response 
→ Update Frontend → Show Shared File in User's Account
```

## 🧪 Testing Strategy

### Backend Testing
- **Unit Tests**: Test individual functions and handlers
- **Integration Tests**: Test database interactions
- **API Tests**: Test endpoints with various inputs
- **Coverage Target**: >80% code coverage

### Frontend Testing
- **Component Tests**: Test React components in isolation
- **Integration Tests**: Test component interactions
- **E2E Tests**: Test full user workflows
- **Coverage Target**: >70% code coverage

---

**Repository**: [0pointshaka/cloud-storage-service](https://github.com/0pointshaka/cloud-storage-service)  
**Repository ID**: 1353019409  
**Created**: 2026-09-01  
**Last Updated**: 2026-09-01
