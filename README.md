# SME Exclusive - B2B Procurement and Supply Chain Management Platform

## Table of Contents
1. [Overview](#overview)
2. [Key Features](#key-features)
3. [Technologies Used](#technologies-used)
4. [Installation](#installation)
   - [Clone the Repository](#clone-the-repository)
   - [Set up the Project](#set-up-the-project)
   - [Configure Environment Variables](#configure-environment-variables)
   - [Run Locally](#run-locally)
   - [Run with Docker](#run-with-docker)
---

## Overview
SME Exclusive is a B2B SaaS platform designed to streamline procurement and supply chain processes, initially focusing on spare parts procurement. It connects buyers (SMEs) with suppliers, allowing them to manage RFQs (Request for Quotations), bid on procurement requests, and access financial institution support for order financing. The platform is built with extensibility in mind, featuring a plugin architecture that will enable the addition of various industries and product categories in future iterations.

The platform supports user authentication via Firebase and features a robust role-based access control (RBAC) system, multi-tenancy support, and integration with external systems. It also provides basic analytics and reporting, ensuring that users have the tools to track procurement activity, orders, and financial transactions.

## Key Features

- **User and Entity Registration and Management**  
  - Google Sign-In through Firebase for authentication  
  - Separate registration for users and entities (Suppliers, Customers, Financial Institutions)  
  - Entity profile creation, verification, and management  

- **Role-Based Access Control (RBAC) and Multi-Tenancy**  
  - Flexible RBAC system to manage user permissions  
  - Support for multi-tenancy, allowing users to switch between roles  
  - Admin interface for role and permission management  

- **RFQ Creation and Management**  
  - Buyers can create and manage RFQs with detailed specifications  
  - Bid deadlines and supplier category settings  
  - RFQ search, filtering, and categorization capabilities  

- **Supplier Bidding System**  
  - Suppliers can submit bids for RFQs and items  
  - Ability to collaborate with NBFCs for financing bids  
  - Bid comparison tools for buyers with customizable evaluation criteria  

- **Order Management**  
  - Convert winning bids into orders and track their status  
  - Order modification, cancellation, and document management  
  - Order-level performance analytics  

- **Financial Institution Integration**  
  - Financial institutions can offer order financing to buyers  
  - System for suppliers to request bids from NBFCs  
  - Secure communication between entities  

- **Basic Analytics and Reporting**  
  - Dashboards and reports for RFQ activity, orders, and financial transactions  
  - Data visualization tools and custom report generation  
  - Export functionality for CSV and PDF formats  

- **Plugin Architecture for Extensibility**  
  - Support for future industry and product category extensions  
  - Plugin management system for easy installation and deactivation  
  - Documentation and guidelines for third-party plugin development  

- **API Integration for External Systems**  
  - RESTful APIs for integration with external systems  
  - OAuth 2.0 authentication and webhook support for real-time notifications  

## Technologies Used
- **Go** - Backend development
- **Fiber** - Web framework
- **MongoDB** - Database
- **Firebase** - User authentication and management
- **Kafka** - Event streaming
- **WebSocket**  - Real time communication


## Installation

### 1. Clone the repository:

```bash
git clone https://github.com/akhilckenshi/SMEExclusive.git
```

### 2. Set up the Project:
Navigate to the root directory of the project and run the following commands to set up the Go module:

```bash
cd exclusive
go mod tidy
```
These commands will create go.sum file, which manage dependencies for the project.


### 3. Configure Environment Variables:

  
3.1. Create a **`.env`** file in the root directory:

```bash
DB_URI="your_db_uri"
DB_NAME="your_db_name"
JWT_SECRET="your_jwt_secret"
PORT="your_port"
```
    
3.2. Create a **`firebase-config.json`** file in the root directory:

```json  
{
    "type": "your_account_type",
    "project_id": "your_project_id",
    "private_key_id": "your_private_key_id",
    "client_email": "your_client_email",
    "client_id": "your_client_id",
    "auth_uri": "your_auth_uri",
    "token_uri": "your_token_uri",
    "auth_provider_x509_cert_url": "your_auth_provider_x509_cert_url",
    "client_x509_cert_url": "your_client_x509_cert_url",
    "universe_domain": "your_universe_domain"
}
```
### 4. Run Locally

Follow these steps to run the project locally:

4.1. Navigate to the root folder of the project:
```bash
cd /path/to/root/folder
```
4.2. Run the project:
```bash
go run main.go
```

### 5. Run with Docker

Follow these steps to run the project with Docker:

5.1. Build the Docker Image

To build the Docker image, navigate to the project root directory and run:

```bash
docker build -t smeexclusive .
```
5.2. Run the Docker Container

To start the container, use the following command:

```bash
docker run -d -p 3000:3000 --name smeexclusive-container smeexclusive
```

5.3. Stop the Running Container

To stop the container, run:

```bash
docker stop smeexclusive-container
```

5.4. Restart the Container

If you need to restart the stopped container, execute:

```bash
docker start smeexclusive-container
```

5.5. Remove the Container

If you need to remove the container before creating a new one:

```bash
docker rm smeexclusive-container
```

5.6. Remove the Docker Image

To delete the Docker image, use:

```bash
docker rmi smeexclusive
```