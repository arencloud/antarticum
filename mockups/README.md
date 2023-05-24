## Project mockups
- All project mockups goes here
- Please update this file to add description for each of them
- Mockups have been made with Penpot and draw.io.

## High-Level Design

The high-level design of the Hospital API engine encompasses the architecture and key components. It follows a modular and layered approach to ensure scalability, maintainability, and reusability. Each building block serves a specific purpose in the overall system design. For a detailed understanding, please refer to the design documentation available in the [mockups](/mockups) folder.

### Token-based and Other Authentication Options

The API engine currently supports token-based authentication as the default authentication mechanism. This choice was made to provide a secure and standardized approach to user authentication. However, the API engine is designed to be extensible, allowing the integration of other authentication methods such as OAuth, OpenID Connect, or even custom authentication schemes based on specific requirements.

### RESTful Design Principles

The API engine adheres to the principles of Representational State Transfer (REST), providing a standardized and intuitive interface for clients to interact with the API. It employs HTTP methods (GET, POST, PUT, DELETE) to perform various operations on resources and follows RESTful URL patterns.

### Database Flexibility

While the API engine currently integrates with a relational database system, it is not limited to only relational databases. The choice of a relational database was made to leverage the advantages of structured data and the well-established ecosystem of relational database management systems. However, the API engine is designed to be adaptable and can be extended to support other database technologies, such as NoSQL databases, depending on specific project needs.

### Error Handling and Response Formats

The API engine includes a robust error handling mechanism to provide meaningful and consistent error responses. It follows best practices for error status codes, error message formats, and error response structures. It also supports multiple response formats, including JSON and XML, to cater to different client requirements.

### Caching Layer

To enhance performance and reduce unnecessary database queries, the API engine incorporates a caching layer. It caches frequently accessed data, such as static resources or frequently retrieved information, to minimize response time and improve overall system performance.

### Pluggable Layers for Extensibility

The API engine incorporates pluggable layers, allowing for easy customization and extension of its functionalities. Pluggable layers provide a way to add or modify components without impacting the core functionality. For example, you can easily plug in additional modules for integrating with third-party services, implementing custom business logic, or supporting specific hospital management features. This design approach promotes modularity, code reusability, and scalability.

### Advantages of Pluggable Layers

The pluggable layers in the API engine offer several advantages:

- **Flexibility**: Pluggable layers allow developers to add or modify functionality without modifying the core codebase, promoting flexibility and adaptability to changing requirements.
- **Modularity**: The API engine's architecture, based on pluggable layers, promotes modular development, making it easier to maintain and test individual components.
- **Code Reusability**: Pluggable layers enable the reuse of existing components, reducing development time and effort.
- **Scalability**: The pluggable architecture allows for the addition of new functionality or services as the system grows, enabling scalability without impacting the existing codebase.

For a visual representation of the system architecture and design decisions, refer to the hospital_api_mockup.jpg file available in the [mockups](/mockups) folder. Additionally, you can explore the detailed API endpoint specifications in the [hospital endpoints' svg](/mockups).
