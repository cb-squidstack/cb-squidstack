# Squid Recommendations

Suggests products based on user activity and review data.

## Features
- Personalized recommendations per user
- Popular product suggestions
- Integration with review and order data

## API
- GET /recommendations/:userId

## Stack
- Go (Gin)
- Internal in-memory scoring engine

## Dependencies
- **Barnacle Reviews**: To analyze review data for recommendation building.
- **Cuttlefish Orders**: For understanding user behaviors and ordering patterns.
