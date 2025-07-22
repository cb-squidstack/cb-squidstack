# Urchin Analytics

Collects and exposes usage metrics and system-wide trends.

## Features
- Tracks orders, payments, and inventory trends
- Simple Prometheus-compatible metrics
- Aggregated daily reports

## API
- GET /metrics
- GET /analytics/dashboard

## Stack
- Go (Gin)
- In-memory store + Prometheus exporter

## CI/CD Workflow Setup

The `urchin-analytics` project is equipped with a CI/CD pipeline integrated through CloudBees. This pipeline uses Kaniko for building Docker images and ArgoCD for deployment.

### Workflow Overview
- **Build Stage**: This step builds a Docker image using Kaniko and pushes it to the specified Docker registry.
- **Manage Charts**: Helm charts are packaged and pushed to a Nexus repository.
- **Deployment Stages**: The build artifact is deployed in the development, pre-production, and production environments using ArgoCD.

### Environment Variables and Secrets
The following variables and secrets must be set in your CloudBees Unify configuration:

- **Secrets**:
  - `DOCKER_PASSWORD`: The password for Docker Hub.
  - `NEXUS_PASSWORD`: The password for Nexus repository.
  - `argocd_cred`: The credential for ArgoCD access.

- **Variables**:
  - `DOCKER_USERNAME`: The username for Docker Hub.
  - `NEXUS_USERNAME`: The username for Nexus repository.
  - `argocd_url`: The ArgoCD server URL.
  - `argocd_user`: The ArgoCD user.

Ensure these variables are correctly set to facilitate the workflow execution and deployment process.

## Helm Chart Details

The `urchin-analytics` Helm chart is used to manage Kubernetes resources for deploying the application. The chart includes the following configuration:

- **Chart Name**: `urchin-analytics`
- **Chart Version**: `0.1.0`
- **Dependencies**: The chart depends on Prometheus for monitoring and metrics collection.
  - **Prometheus Repo**: [Prometheus Community Helm Charts](https://prometheus-community.github.io/helm-charts)

Place the Helm chart in the `chart/` directory and update the `chart/urchin-analytics/Chart.yaml` as needed for further customization.