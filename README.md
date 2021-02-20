# go-service-gin
go started template for building microservice using framework gin

---

## Architecture
**Domain-Driven-Design** is the concept that the structure and language of software code should match the business domain. <br><br>
DDD has 4 layers in the architecture:

- **Interface**: This layer responsibles for the interaction with user, whether software presents information or recieves information from user.
- **Application**: This is a thin layer between interface and domain, it could call domain services to serve the application purposes.
- **Domain**: The heart of the software, this layer holds domain logic and business knowledge.
- **Infrastructure**: A supporting layer for the other layers. This layer contains supporting libraries or external services like database or UI supporting library.

---

## Table of Content