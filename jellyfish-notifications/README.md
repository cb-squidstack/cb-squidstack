# Jellyfish Notifications

Manages user-facing notifications across the SquidStack platform.

## Features
- Sends real-time or delayed notifications
- Integration with orders, payments, and shipping
- Future support for email/SMS hooks

## API
- POST /notify
- GET /notifications/:userId

## Stack
- Go (Gin)
- In-memory queue or Redis
