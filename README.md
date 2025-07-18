
# cb-squidstack

## Purpose
This repository orchestrates all services within the cb-squidstack system, providing Docker Compose and a Makefile for managing the stack.

## Service Dependency Map

| Service               | Depends On                                        |
|------------------------|--------------------------------------------------|
| **squid-ui**           | kraken-auth, cuttlefish-orders, octopus-payments, jellyfish-notifications |
| **squid-recommendations** | barnacle-reviews, cuttlefish-orders            |
| **barnacle-reviews**   | kraken-auth                                      |
| **octopus-payments**   | kraken-auth, cuttlefish-orders                   |
| **cuttlefish-orders**  | kraken-auth, nautilus-inventory, octopus-payments |
| **manta-delivery**     | cuttlefish-orders                                 |
| **nautilus-inventory** | Provides /inventory, /health                     |

## Running the Stack
```bash
make up
make healthcheck
```

## Health Endpoints
- squid-ui: `/squidui`
- nautilus-inventory: `/inventory`, `/inventory/:productId`, `/health`
