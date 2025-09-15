

## Project Structure

This repository (`cb-squidstack`) serves as the orchestration layer for all SquidStack services.

### Directory Layout
```
/Users/brown/git_orgs/
├── cb-squidstack/             # This repo: orchestration, compose, makefile, scripts
│   ├── docker-compose.yaml
│   ├── Makefile
│   ├── README.md
│   └── scripts/
│       └── healthcheck.sh
├── squid-ui/                   # Separate service repos
├── nautilus-inventory/
├── manta-delivery/
├── squid-recommendations/
├── barnacle-reviews/
├── cuttlefish-orders/
├── octopus-payments/
├── urchin-analytics/
├── jellyfish-notifications/
├── kraken-auth/
├── clam-catalog/
```

## Usage

Ensure you are inside the `cb-squidstack` directory to run commands:

```bash
make up
make healthcheck
make down
```

## Notes

- Each service is in its own repo at the same directory level as `cb-squidstack`.
- `docker-compose.yaml` references them using `../<service>`.
- This keeps orchestration clean and separate from app code.

# SquidStack Architecture

SquidStack is a demo microservices application with marine-themed services.

## System Diagram

```mermaid
flowchart TB
  UI[squid-ui\nReact frontend\n(no DB)]
  AUTH[kraken-auth\nAuth/JWT + Profiles\nPostgreSQL\nschemas: auth, public]

  UI -->|Login / Admin APIs| AUTH

  subgraph DBYES[Services with DB (planned)]
    ORD[cuttlefish-orders\nschema: orders]
    CAT[clam-catalog\nschema: catalog]
    REV[barnacle-reviews\nschema: reviews]
    INV[nautilus-inventory\nschema: inventory]
    ANA[urchin-analytics\nschema: analytics]
  end

  subgraph DBNO[Stateless services (planned)]
    PAY[octopus-payments]
    REC[squid-recommendations]
    NOT[jellyfish-notifications]
  end

  classDef db fill:#eef7ff,stroke:#88a,stroke-width:1px;
  classDef ndb fill:#f9f9f9,stroke:#bbb,stroke-width:1px;

  class AUTH,ORD,CAT,REV,INV,ANA db;
  class PAY,REC,NOT ndb;
