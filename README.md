# Crud Enterprise

## Table of Contents

- [Crud Enterprise](#crud-enterprise)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Features](#features)
  - [Project Structure](#project-structure)
  - [Requirements](#requirements)
  - [How to Run](#how-to-run)
  - [Team](#team)

---

## Overview

This project is a simple implementation of a CRUD (Create, Read, Update, Delete) system in Go. It manages five main entities: Bosses (```chefes_departamentos```), Departments (```departamentos```), Employees (```funcionarios```), Employees_Departments (```funcionarios_departamentos```) and Projects (```projetos```). The data is momently persisted in JSON files and storage in Text files. The application demonstrates basic operations for managing these entities.

---

## Features

- Bosses CRUD:
   - Add new bosses to existing departments.
   - Update existing bosses by ID.
   - Delete bosses by ID.
   - List all bosses.

- Departments CRUD:
   - Add new departments to existing bosses.
   - Update existing departments by ID.
   - Delete departments by ID.
   - List all departments.

- Employees CRUD:
   - Add new employees to existing departments.
   - Update existing employees by ID.
   - Delete employees by ID.
   - List all employees.

- Employees_Projects CRUD:
   - Add new employees to new projects.
   - Update existing employees to existing projects by ID.
   - Delete employees from projects by ID.
   - List all employees from projects.

- Projects CRUD:
   - Add new projects to existing employees and departments.
   - Update existing projects by ID.
   - Delete projects by ID.
   - List all projects.

- JSON Persistence:
   - All data is momently stored in JSON files (```data/json/chefes_departamentos.json```), (```data/json/departamentos.json```), (```data/json/funcionarios.json```), , (```data/json/funcionarios_projetos.json```) and (```data/json/projetos.json```), and stored in text files.
   - The JSON files are cleared at the end of the program execution.

- Text Storage:
   - All data is stored in Text files (```data/txt/chefes_departamentos.txt```), (```data/txt/departamentos.txt```), (```data/txt/funcionarios.txt```), , (```data/txt/funcionarios_projetos.txt```) and (```data/txt/projetos.txt```).

---

## Project Structure

```bash
📦 CRUD_ENTERPRISE   # Main project folder
┣ 📂 controllers  
┃ ┣ 🐹 screen_controller.go     
┣ 📂 data
┃ ┣ 📂 json
┃ ┃ ┣ 💾 chefes_departamentos.json
┃ ┃ ┣ 💾 departamentos.json
┃ ┃ ┣ 💾 funcionarios.json
┃ ┃ ┣ 💾 funcionarios_projetos.json
┃ ┃ ┣ 💾 projetos.json
┃ ┣ 📂 txt
┃ ┃ ┣ 📄 chefes_departamentos.txt
┃ ┃ ┣ 📄 departamentos.txt
┃ ┃ ┣ 📄 funcionarios.txt
┃ ┃ ┣ 📄 funcionarios_projetos.txt
┃ ┃ ┣ 📄 projetos.txt
┣ 📂 models
┃ ┣ 🐹 chefe_departamento.go             # CRUD logic for bosses
┃ ┣ 🐹 departamento.go                   # CRUD logic for departments
┃ ┣ 🐹 funcionario.go                    # CRUD logic for employees
┃ ┣ 🐹 funcionarios_projetos.go          # CRUD logic for employees and projects
┃ ┣ 🐹 projetos.go                       # CRUD logic for projects
┣ 📂 utils
┃ ┣ 🐹 gui_utils.go                      # Utility functions (e.g., clearing Entry Widgets)
┃ ┣ 🐹 utils.go                          # Utility functions (e.g., clearing JSON files)
┣ 📂 views
┃ ┣ 📂 main_view
┃ ┃ ┣ 🐹 main_page.go                    # GUI main page
┃ ┣ 📂 table_view 
┃ ┃ ┣ 🐹 chefes_departamentos_page.go    # GUI for bosses
┃ ┃ ┣ 🐹 departamentos_page.go           # GUI for departments
┃ ┃ ┣ 🐹 funcionarios_page.go            # GUI for employees
┃ ┃ ┣ 🐹 projetos_page.go                # GUI for projects
┣ 🛠️ go.mod                              # Go module file
┣ 🛠️ go.sum                              # Go Sum file
┣ 🐹 main.go                             # Entry point of the application
┣ 📄 README.md                           # Project documentation          
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

2. Install the dependencies:

````bash
go mod tidy
````

3. Run the application:

````bash
go run main.go
````

or use the ```Code Runner``` VSCode Extension in main.go.

---

## Team

- Conrado Einstein
- Hiel Saraiva
- João Marcelo Pimenta