# Octopus Payments

Processes and simulates payment transactions.

## Features
- Payment initiation and status tracking
- Simulated external payment gateway
- Support for mock failures and retries

## API
- POST /pay
- GET /payment/:id

## Stack
- Go (Gin)
- PostgreSQL

## Dependencies
- **Kraken Auth**: For validating user identity and ensuring secure transactions.
- **Cuttlefish Orders**: Collaborates to facilitate seamless payment processes during order placement.
