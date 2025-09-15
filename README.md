# 🦑 SquidStack

SquidStack is a demo **microservices application** with a marine theme.  
It’s designed to showcase microservice patterns, authentication, role-based access, and feature management.

---

## 🌊 Components

| Service                | Type      | Status                | Purpose                                   | Database |
|-------------------------|-----------|-----------------------|-------------------------------------------|----------|
| [kraken-auth](../kraken-auth/README.md) | Service   | ✅ Implemented         | Authentication, JWT issuance, role/profile mgmt | **Yes** (`auth` schema) |
| [squid-ui](../squid-ui/README.md)       | Frontend | 🟡 Partially Implemented | React frontend; login & admin screens     | No       |
| [cuttlefish-orders](../cuttlefish-orders/README.md) | Service   | Stub (health only)    | Order processing                          | **Yes** (`orders`) |
| [octopus-payments](../octopus-payments/README.md)   | Service   | Stub (health only)    | Payment simulation                        | No       |
| [clam-catalog](../clam-catalog/README.md)           | Service   | Stub (health only)    | Product catalog                           | **Yes** (`catalog`) |
| [barnacle-reviews](../barnacle-reviews/README.md)   | Service   | Stub (health only)    | Product reviews/ratings                   | **Yes** (`reviews`) |
| [squid-recommendations](../squid-recommendations/README.md) | Service | Stub (health only)    | Recommendation engine                     | No       |
| [nautilus-inventory](../nautilus-inventory/README.md)       | Service | Stub (health only)    | Inventory/stock tracking                  | **Yes** (`inventory`) |
| [urchin-analytics](../urchin-analytics/README.md)           | Service | Stub (health only)    | Event collection/analytics                | **Yes** (`analytics`) |
| [jellyfish-notifications](../jellyfish-notifications/README.md) | Service | Stub (health only)    | Notifications (email/SMS/in-app)          | No       |

---

## 🗂️ Current Status

- **Implemented:**  
  - `kraken-auth` (full DB + JWT + admin APIs)

- **Partially implemented:**  
  - `squid-ui` (login/auth integrated, admin user view)

- **Stubs with healthcheck only:**  
  - `cuttlefish-orders`, `octopus-payments`, `clam-catalog`,  
    `barnacle-reviews`, `squid-recommendations`, `nautilus-inventory`,  
    `urchin-analytics`, `jellyfish-notifications`

---

## 📐 Architecture

### ASCII overview (always works)
