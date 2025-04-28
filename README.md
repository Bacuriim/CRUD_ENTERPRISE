# Crud Enterprise

## Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Project Structure](#project-structure)
4. [Requirements](#requirements)
5. [How to Run](#how-to-run)
6. [Team](#team)

---

## Overview

This project is a simple implementation of a CRUD (Create, Read, Update, Delete) system in Go. It manages two main entities: Employees (```funcionarios```) and Departments (```departamentos```). The data is stored in JSON files, and the application demonstrates basic operations for managing these entities.

---

## Features

- Employees CRUD:
   - Add new employees.
   - Update existing employees by ID.
   - Delete employees by ID.
   - List all employees.

- Departments CRUD:
   - Add new departments.
   - Update existing departments by ID.
   - Delete departments by ID.
   - List all departments.

- JSON Persistence:
   - All data is stored in JSON files (```data/funcionarios.json``` and ```data/departamentos.json```).
   - The JSON files are cleared at the end of the program execution.

---

## Project Structure

```bash
📦 CRUD_ENTERPRISE            # Main project folder
┣ 📂 data
┃ ┣ 📄 funcionarios.json      # JSON file for storing employees
┃ ┣ 📄 departamentos.json     # JSON file for storing departments
┣ 📂 models
┃ ┣ 🐹 funcionario.go         # CRUD logic for employees
┃ ┣ 🐹 departamento.go        # CRUD logic for departments
┣ 📂 utils
┃ ┣ 🐹 utils.go               # Utility functions (e.g., clearing JSON files)
┣ 🐹 main.go                  # Entry point of the application
┣ 🛠️ go.mod                   # Go module file
┣ 📄 README.md                # Project documentation          
```

---

## Requirements

- Go 1.20 or higher: [GoLang Download](https://go.dev/doc/install)
- A terminal or IDE to run the application: [VSCode](https://code.visualstudio.com/download)
   - VSCode Extensions:
      - Code Runner 
      - Go

---

## How to Run

1. Clone the repository:

````bash
git clone https://github.com/Bacuriim/CRUD_ENTERPRISE.git
````

2. Run the application:

````bash
go run main.go
````

or use the ```Code Runner``` VSCode Extension in main.go.

---

## Team

- Conrado Einstein
- Hiel Saraiva
- João Marcelo Pimenta