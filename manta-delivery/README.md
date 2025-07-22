# Manta Delivery

Coordinates delivery logistics and shipping status.

## Features
- Assigns carriers
- Updates delivery status
- Simulates shipping delays

## API
- POST /ship
- GET /delivery/:orderId

## Stack
- Go (Gin)
- PostgreSQL

## Dependencies
- **Cuttlefish Orders**: To fetch order and shipping details for coordination.
