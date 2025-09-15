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

### Mermaid (optional — GitHub only)

```mermaid
flowchart TB
  UI["squid-ui<br/>React frontend<br/>(no DB)"]
  AUTH["kraken-auth<br/>Auth/JWT + Profiles<br/>PostgreSQL<br/>schemas: auth, public"]

  UI -->|Login / Admin APIs| AUTH

  subgraph DBYES["Services with DB (planned)"]
    ORD["cuttlefish-orders<br/>schema: orders"]
    CAT["clam-catalog<br/>schema: catalog"]
    REV["barnacle-reviews<br/>schema: reviews"]
    INV["nautilus-inventory<br/>schema: inventory"]
    ANA["urchin-analytics<br/>schema: analytics"]
  end

  subgraph DBNO["Stateless services (planned)"]
    PAY["octopus-payments"]
    REC["squid-recommendations"]
    NOT["jellyfish-notifications"]
  end

  classDef db fill:#eef7ff,stroke:#88a,stroke-width:1px,color:#111;
  classDef ndb fill:#f9f9f9,stroke:#bbb,stroke-width:1px,color:#111;

  class AUTH,ORD,CAT,REV,INV,ANA db;
  class PAY,REC,NOT ndb;

🔑 Key Design Principles
	•	Separation of concerns:
Each service owns its domain (auth, orders, catalog, reviews, etc).
	•	Database per service:
Only backend services own DBs (Postgres schemas).
squid-ui is stateless and stores only JWT + user snapshot in browser localStorage.
	•	JWT-based security:
All service-to-service and frontend-to-service calls are secured with tokens from kraken-auth.
	•	Feature management:
Integrated with CloudBees Unify for flags controlling UI + rollout.

⸻

📎 Related Docs

Each service has its own README:
	•	kraken-auth
	•	squid-ui
	•	cuttlefish-orders
	•	octopus-payments
	•	clam-catalog
	•	barnacle-reviews
	•	squid-recommendations
	•	nautilus-inventory
	•	urchin-analytics
	•	jellyfish-notifications
