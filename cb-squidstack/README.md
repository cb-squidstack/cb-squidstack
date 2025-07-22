

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
