# GradeHarbor

## Description

Welcome to the GradeHarbor Backend! This project is designed to managing school grading system, focusing on user roles such as admin, teacher, parent, and student. The project is built using Golang with a microservices architecture including Auth, User, and Class microservices. Communication between microservices is facilitated through gRPC, and PostgreSQL is used as the database. Additionally, Apache Kafka is employed for sending confirmation codes during user registration.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Installation

1. Clone the repository:

   ```bash
   git clone [https://github.com/alibekabdrakhman1/gradeHarbor]
   cd school-grading-system
   docker-compose build

