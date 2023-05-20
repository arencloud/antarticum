## Project mockups
- All project mockups goes here
- Please update this file to add description for each of them
- Mockups have been made with Penpot and draw.io.

## High-Level Design

The high-level design of the Hospital API engine encompasses the architecture and key components. It follows a modular and layered approach to ensure scalability, maintainability, and reusability. Each building block serves a specific purpose in the overall system design. For a detailed understanding, please refer to the design documentation available in the [mockups](/mockups) folder.

### Token-based Authentication

The API engine incorporates a token-based authentication mechanism to secure the endpoints and control access. It uses industry-standard authentication protocols, such as JSON Web Tokens (JWT), to authenticate and authorize requests.

### RESTful Design Principles

The API engine adheres to the principles of Representational State Transfer (REST), providing a standardized and intuitive interface for clients to interact with the API. It employs HTTP methods (GET, POST, PUT, DELETE) to perform various operations on resources and follows RESTful URL patterns.

### Relational Database Integration

The API engine integrates with a relational database system to store and retrieve data efficiently. It utilizes an ORM (Object-Relational Mapping) framework to facilitate seamless interaction with the database. The schema design ensures proper normalization, relationships, and indexes for optimal performance.

### Error Handling and Response Formats

The API engine includes a robust error handling mechanism to provide meaningful and consistent error responses. It follows best practices for error status codes, error message formats, and error response structures. It also supports multiple response formats, including JSON and XML, to cater to different client requirements.

### Caching Layer

To enhance performance and reduce unnecessary database queries, the API engine incorporates a caching layer. It caches frequently accessed data, such as static resources or frequently retrieved information, to minimize response time and improve overall system performance.

For a visual representation of the system architecture and design decisions, refer to the hospital_api_mockup.jpg file available in the [mockups](/mockups) folder. Additionally, you can explore the detailed API endpoint specifications in the [hospital endpoints' svg](/mockups).
