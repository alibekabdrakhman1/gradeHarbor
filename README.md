# GradeHarbor

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Description

Welcome to the GradeHarbor Backend! This project is designed to managing school grading system, focusing on user roles such as admin, teacher, parent, and student. The project is built using Golang with a microservices architecture including Auth, User, and Class microservices. Communication between microservices is facilitated through gRPC, and PostgreSQL is used as the database. Additionally, Apache Kafka is employed for sending confirmation codes during user registration.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Usage](#usage)
5. [Microservices](#microservices)
    - [Auth](#auth)
    - [User](#user)
    - [Class](#class)
6. [Communication](#communication)
7. [Database](#database)
8. [Docker Compose](#docker-compose)
9. [Contributing](#contributing)
10. [License](#license)

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Installation

1. Clone the repository:

   ```bash
   git clone [https://github.com/your_username/school-grading-system.git](https://github.com/alibekabdrakhman1/gradeHarbor)https://github.com/alibekabdrakhman1/gradeHarbor
   cd school-grading-system
   docker-compose build

