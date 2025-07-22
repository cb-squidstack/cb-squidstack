# Cuttlefish Orders

Handles order placement, validation, and status updates.

## Features
- Order creation and retrieval
- Integration with payments and inventory
- Per-user order history

## API
- POST /orders
- GET /orders/:id
- GET /orders/user/:userId

## Stack
- Go (Gin)
- PostgreSQL

## Dependencies
- **Kraken Auth**: For validating and managing user authentication.
- **Nautilus Inventory**: To check stock levels and update inventory during order processing.
- **Octopus Payments**: For processing payment transactions.
