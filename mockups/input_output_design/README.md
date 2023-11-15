# Pluggable Interface for Handling Multiple Input/Output Formats

This repository demonstrates the implementation of a pluggable interface in an API to support multiple input/output formats, including JSON, XML, and YAML. The interface allows users to interact with the API using their preferred data format while maintaining a common data structure internally.

## Overview

Modern APIs often need to cater to a variety of data formats to accommodate different client needs. This project addresses this challenge by providing a flexible system for processing requests and responses in different formats, while still adhering to a uniform data structure within the application.

## Features

- Supports JSON, XML, and YAML input and output formats.
- Utilizes a middleware for format detection, conversion, and content negotiation.
- Maintains a common data structure that acts as an intermediary between different formats.
- Provides flexibility for users to specify their preferred input/output format.


## Examples

- **JSON Input:**
  ```json
  { "name": "John Doe", "age": 30, "email": "john@example.com" }
  ```

- **XML Input:**
  ```xml
  <user>
    <name>John Doe</name>
    <age>30</age>
    <email>john@example.com</email>
  </user>
  ```

- **YAML Input:**
  ```yaml
  name: John Doe
  age: 30
  email: john@example.com
  ```
