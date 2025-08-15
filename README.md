# Go Lang To-Do App 📝

A simple **To-Do application** built with **Go** as the backend using **Gin framework**, **SQLite (via GORM)** for storage, and a frontend with **HTML, CSS, and JavaScript**.

This project was built **after completing Striver’s Complete Go Developer course** to explore Go and practice **full-stack development with Go**.

## Features ✨

- Add new tasks
- Delete tasks 🗑️
- Update task status (mark as done ✅)
- Persistent storage using **SQLite** via **GORM**
- Minimal frontend using **HTML, CSS, JS**
- Backend powered by **Gin framework**
- Serves frontend at **localhost:3001**

## Tech Stack 🛠️

- **Backend:** Go (Golang) with **Gin**
- **Database:** SQLite (via **GORM**)
- **Frontend:** HTML, CSS, JavaScript
- **ORM:** GORM

## Learning Context 🎓

- Built after learning Go through **Striver’s Complete Go Developer course**
- Focused on **CRUD operations, routing, database integration, and serving HTML templates**
- Practical experience combining **Go backend** with a **frontend**

## Installation ⚡

1. Clone the repository:

```bash
git clone https://github.com/your-username/todo-app.git
cd todo-app/backend
```

2. Install Go modules:

```bash
go mod tidy
```

3. Run the backend application (serves frontend on port 3000):

```bash
go run todoapp.go
```

4. Open your browser at:

```
http://localhost:3001
```

## Usage 🖱️

- **Add Task:** Enter task name and click **Add**
- **Delete Task:** Click the delete button next to a task
- **Update Status:** Check/uncheck the task to mark it as done/pending

All changes are persisted in **SQLite database (**``**)** automatically.

## Project Structure 📁

```
todo-app/
├── backend/
│   ├── controllers/     # Controllers for handling routes
│   ├── database/        # DB connection and migration
│   ├── models/          # Task model (GORM)
│   ├── renderer/        # Template rendering
│   ├── routes/          # Route definitions
│   ├── todoapp.go       # Main backend entry
│   ├── todo.db          # SQLite database
│   ├── go.mod           # Go modules
│   └── go.sum           # Go modules sum
├── frontend/
│   ├── assets/          # CSS, JS, images
│   ├── index.html       # Frontend HTML
│   ├── styles.css       # CSS
│   └── script.js        # JS
└── README.md            # This file

go-labs/                 # Folder with other Go practice projects
```



