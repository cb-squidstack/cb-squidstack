# 🦑 SquidStack

SquidStack is a demo **microservices application** with a marine theme.  
It’s designed to showcase microservice patterns, authentication, role-based access, and feature management.

---

## 🌊 Components

| Service | Type | Status | Purpose | Database |
|---------|------|--------|---------|----------|
| [kraken-auth](https://github.com/cb-squidstack/kraken-auth/blob/main/README.md) | Service | ✅ Implemented | Authentication, JWT issuance, role/profile mgmt | **Yes** (`auth` schema) |
| [squid-ui](https://github.com/cb-squidstack/squid-ui/blob/main/README.md) | Frontend | 🟡 Partially Implemented | React frontend; login & admin screens | No |
| [cuttlefish-orders](https://github.com/cb-squidstack/cuttlefish-orders/blob/main/README.md) | Service | Stub (health only) | Order processing | **Yes** (`orders`) |
| [octopus-payments](https://github.com/cb-squidstack/octopus-payments/blob/main/README.md) | Service | Stub (health only) | Payment simulation | No |
| [clam-catalog](https://github.com/cb-squidstack/clam-catalog/blob/main/README.md) | Service | Stub (health only) | Product catalog | **Yes** (`catalog`) |
| [barnacle-reviews](https://github.com/cb-squidstack/barnacle-reviews/blob/main/README.md) | Service | Stub (health only) | Product reviews/ratings | **Yes** (`reviews`) |
| [squid-recommendations](https://github.com/cb-squidstack/squid-recommendations/blob/main/README.md) | Service | Stub (health only) | Recommendation engine | No |
| [nautilus-inventory](https://github.com/cb-squidstack/nautilus-inventory/blob/main/README.md) | Service | Stub (health only) | Inventory/stock tracking | **Yes** (`inventory`) |
| [urchin-analytics](https://github.com/cb-squidstack/urchin-analytics/blob/main/README.md) | Service | Stub (health only) | Event collection/analytics | **Yes** (`analytics`) |
| [jellyfish-notifications](https://github.com/cb-squidstack/jellyfish-notifications/blob/main/README.md) | Service | Stub (health only) | Notifications (email/SMS/in-app) | No |
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

```
                       ┌───────────────────────────┐
                       │         squid-ui          │
                       │  React frontend (no DB)   │
                       └────────────┬──────────────┘
                                    │  Login / Admin APIs
                                    ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                        kraken-auth (service + DB)                           │
│  • AuthN / JWT issuance, roles, profiles                                    │
│  • PostgreSQL (schema: auth + public)                                       │
│                                                                             │
│  Tables:                                                                    │
│   - users, auth_credentials, roles, user_roles                              │
│   - user_profiles (full_name, email, phone, address, country_code, roles[]) │
│   - countries                                                               │
└─────────────────────────────────────────────────────────────────────────────┘

Other services (stubs today, health check only)
───────────────────────────────────────────────────────────────────────────────
DB = Yes                                DB = No
────────────────                        ────────────────
• cuttlefish-orders  (schema: orders)   • octopus-payments
• clam-catalog       (schema: catalog)  • squid-recommendations
• barnacle-reviews   (schema: reviews)  • jellyfish-notifications
• nautilus-inventory (schema: inventory)
• urchin-analytics   (schema: analytics)
```


---

## 🔑 Key Design Principles

- **Separation of concerns:**  
  Each service owns its domain (auth, orders, catalog, reviews, etc).

- **Database per service:**  
  Only backend services own DBs (Postgres schemas).  
  `squid-ui` is stateless and stores only JWT + user snapshot in browser localStorage.  

- **JWT-based security:**  
  All service-to-service and frontend-to-service calls are secured with tokens from `kraken-auth`.

- **Feature management:**  
  Integrated with **CloudBees Unify** for flags controlling UI + rollout.

---

## 📎 Related Docs

Each service has its own README:

- [kraken-auth](https://github.com/cb-squidstack/kraken-auth/blob/main/README.md)  
- [squid-ui](https://github.com/cb-squidstack/squid-ui/blob/main/README.md)  
- [cuttlefish-orders](https://github.com/cb-squidstack/cuttlefish-orders/blob/main/README.md)  
- [octopus-payments](https://github.com/cb-squidstack/octopus-payments/blob/main/README.md)  
- [clam-catalog](https://github.com/cb-squidstack/clam-catalog/blob/main/README.md)  
- [barnacle-reviews](https://github.com/cb-squidstack/barnacle-reviews/blob/main/README.md)  
- [squid-recommendations](https://github.com/cb-squidstack/squid-recommendations/blob/main/README.md)  
- [nautilus-inventory](https://github.com/cb-squidstack/nautilus-inventory/blob/main/README.md)  
- [urchin-analytics](https://github.com/cb-squidstack/urchin-analytics/blob/main/README.md)  
- [jellyfish-notifications](https://github.com/cb-squidstack/jellyfish-notifications/blob/main/README.md)  
---
