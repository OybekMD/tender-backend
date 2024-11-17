# Tender Management System

## Overview

The **Tender Management System** is a backend application built to manage procurement processes (tenders) and bids. It allows **Clients** to post tenders, while **Contractors** can submit bids. Clients can evaluate bids and award tenders to contractors based on submitted bids. The system includes features like user authentication, role management, tender creation, bid submission, bid evaluation, and history tracking.

### Key Features:
- **User Authentication & Role Management**: Registration and login for Clients and Contractors with JWT-based authentication.
- **Tender Posting**: Clients can create, manage, and delete tenders.
- **Bid Submission**: Contractors can submit bids with price, delivery time, and comments.
- **Bid Evaluation & Tender Awarding**: Clients can review and award tenders based on bids.
- **Bid Filtering & Sorting**: Clients can filter and sort bids based on price and delivery time.
- **Real-Time Updates** (Optional): WebSocket integration for real-time notifications.
- **Rate Limiting**: Prevention of spammy bid submissions.
- **Caching**: Improves performance by caching tender and bid data.

## Technologies Used
- **Backend**: Golang (Go)
- **Database**: PostgreSQL (or MongoDB)
- **API**: RESTful APIs
- **Real-time Notifications**: WebSocket (optional)
- **Documentation**: Swagger for API documentation
- **Containerization**: Docker for easy deployment
- **Authentication**: JWT (JSON Web Tokens)

## Database Schema

- **User**: `id, username, password, role (client/contractor), email`
- **Tender**: `id, client_id, title, description, deadline, budget, status`
- **Bid**: `id, tender_id, contractor_id, price, delivery_time, comments, status`
- **Notification**: `id, user_id, message, relation_id, type, created_at`

## Project Setup

### Prerequisites
- Docker (for containerization)
- Make (for managing build and run tasks)

### Steps to Run the Project

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourusername/tender-management-system.git
   cd tender-management-system
