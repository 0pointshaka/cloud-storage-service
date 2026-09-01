# Contributing to Cloud Storage Service

Thank you for your interest in contributing to Cloud Storage Service! This document provides guidelines and instructions for contributing to the project.

## 🎯 Code of Conduct

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.

## 📋 Getting Started

### Prerequisites

Before you start contributing, ensure you have the following installed:

**For Backend Development (Go):**
- Go 1.21 or later - [Download](https://golang.org/dl/)
- Git for version control

**For Frontend Development (React):**
- Node.js 16 or later - [Download](https://nodejs.org/)
- npm or yarn package manager

**For All Contributors:**
- A GitHub account
- Git installed on your machine
- Familiarity with basic Git workflow (fork, clone, commit, push, pull request)

### Development Environment Setup

#### Windows Users

1. Clone the repository:
```batch
git clone https://github.com/0pointshaka/cloud-storage-service.git
cd cloud-storage-service
```

2. Check prerequisites:
```batch
setup.bat
```

3. Build the project:
```batch
build.bat
```

4. Start development servers:
```batch
run-dev.bat
```

#### Mac/Linux Users

1. Clone the repository:
```bash
git clone https://github.com/0pointshaka/cloud-storage-service.git
cd cloud-storage-service
```

2. Set up backend:
```bash
cd backend
go mod download
export JWT_SECRET="dev-secret-key"
go run main.go
```

3. In another terminal, set up frontend:
```bash
cd frontend
npm install
npm run dev
```

## 🔄 Contributing Workflow

### 1. Fork the Repository

Click the "Fork" button in the top-right corner of the repository page.

### 2. Create a Feature Branch

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/issue-number
```

**Branch naming conventions:**
- `feature/` - for new features
- `fix/` - for bug fixes
- `docs/` - for documentation changes
- `refactor/` - for code refactoring
- `test/` - for test additions
- `chore/` - for maintenance tasks

### 3. Make Your Changes

#### Backend (Go) Development

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code
- Write unit tests for new functionality
- Keep functions small and focused
- Add meaningful comments for exported functions

Example:
```go
// HandleUploadFile processes file uploads and stores them in the database
func HandleUploadFile(c *gin.Context) {
    // Implementation here
}
```

#### Frontend (React) Development

- Follow the [Airbnb React/JSX Style Guide](https://github.com/airbnb/javascript/tree/master/react)
- Use functional components and hooks
- Keep components focused and reusable
- Write meaningful component documentation
- Use proper prop validation

Example:
```jsx
/**
 * FileUploadButton - Handles file selection and upload
 * @param {function} onUpload - Callback when file is uploaded
 * @param {boolean} loading - Loading state indicator
 */
function FileUploadButton({ onUpload, loading }) {
    // Implementation here
}
```

### 4. Commit Your Changes

Write clear, descriptive commit messages:

```bash
git commit -m "feat: add file upload progress indicator"
# or
git commit -m "fix: resolve JWT token expiration issue"
```

**Commit message format:**
```
<type>: <subject>

<body>

<footer>
```

**Types:**
- `feat:` - A new feature
- `fix:` - A bug fix
- `docs:` - Documentation only changes
- `refactor:` - Code change that neither fixes a bug nor adds a feature
- `perf:` - Code change that improves performance
- `test:` - Adding or updating tests
- `chore:` - Changes to build process, dependencies, etc.

**Example:**
```
feat: implement file sharing functionality

- Add share endpoint to API
- Create sharing handler
- Add database migration for shared_files table
- Add frontend UI for share dialog

Fixes #123
```

### 5. Push to Your Fork

```bash
git push origin feature/your-feature-name
```

### 6. Create a Pull Request

1. Go to your fork on GitHub
2. Click "New Pull Request"
3. Select your branch
4. Fill in the PR template with:
   - Clear description of changes
   - Related issues (use `#123` format)
   - Screenshots (if UI changes)
   - Testing instructions
   - Any breaking changes

**Pull Request Title Format:**
```
[Backend/Frontend] Brief description of changes
```

## 🧪 Testing

### Backend Testing (Go)

```bash
cd backend
go test ./...
go test -v ./...          # Verbose output
go test -cover ./...      # Coverage report
```

### Frontend Testing (React)

```bash
cd frontend
npm test                   # Run tests
npm test -- --coverage    # Coverage report
```

**Requirements:**
- All new features must include tests
- Bug fixes should include a test that would have caught the bug
- Maintain or improve code coverage
- Tests must pass before merge

## 📝 Code Style

### Go Code Style

- Use `gofmt` for formatting
- Variable names should be concise
- Use error handling consistently
- Follow the [Effective Go](https://golang.org/doc/effective_go) guide

```go
// Good
func (u *User) HasAccess(file *File) bool {
    return u.ID == file.UserID
}

// Avoid
func (u *User) CheckIfUserHasAccessToFile(file *File) bool {
    // ...
}
```

### React Code Style

- Use functional components
- Use meaningful component and variable names
- Proper prop types and validation
- Consistent formatting with Prettier

```jsx
// Good
function FileList({ files, onDelete }) {
    return (
        <ul>
            {files.map(file => (
                <FileItem key={file.id} file={file} onDelete={onDelete} />
            ))}
        </ul>
    );
}

// Avoid
function fl({ f, od }) {
    return <ul>{f.map(file => <li key={file.id}>{file.name}</li>)}</ul>;
}
```

## 📚 Documentation

- Update README.md if adding new features
- Add comments to complex logic
- Update API documentation in LANGUAGE_STATS.md
- Create/update wiki pages for complex features

## 🐛 Reporting Bugs

When reporting bugs, include:

1. **Description**: Clear description of the bug
2. **Steps to Reproduce**: Step-by-step instructions
3. **Expected Behavior**: What should happen
4. **Actual Behavior**: What actually happens
5. **Environment**: OS, Node version, Go version, etc.
6. **Screenshots/Logs**: If applicable

## 🎨 Feature Requests

When suggesting features, include:

1. **Use Case**: Why you need this feature
2. **Description**: How it should work
3. **Alternatives**: Any alternative solutions
4. **Additional Context**: Any other relevant information

## 📖 Documentation

We welcome documentation improvements! You can:

- Fix typos
- Improve existing documentation
- Add missing documentation
- Create tutorials or guides
- Improve code comments

## ✅ Review Process

1. **Automated Checks**: GitHub Actions runs tests and linters
2. **Code Review**: Maintainers review the code
3. **Discussion**: Feedback and suggestions may be provided
4. **Approval**: Once approved, your PR will be merged

## 🚀 Deployment

Maintainers handle the release and deployment process. After merge:

1. Changes are reviewed for release notes
2. Version number is updated
3. Release is tagged
4. Changes are deployed to production

## 📞 Questions or Need Help?

- Check existing [GitHub Issues](https://github.com/0pointshaka/cloud-storage-service/issues)
- Read the [README.md](README.md)
- Check [LANGUAGE_STATS.md](LANGUAGE_STATS.md) for project structure
- Create a new issue with the `question` label

## 🙏 Thank You!

Thank you for contributing to Cloud Storage Service! Your efforts help make this project better for everyone.

---

**Happy Contributing! 🎉**
