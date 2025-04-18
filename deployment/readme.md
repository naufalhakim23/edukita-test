# Kubernetes Deployment Guide for Edukita LMS

This guide explains how to deploy the Edukita LMS application on Kubernetes using the provided YAML files.

## Prerequisites

- Kubernetes cluster (EKS, GKE, AKS, or any other Kubernetes platform)
- kubectl command-line tool configured to communicate with your cluster
- Docker image repository (Docker Hub, ECR, GCR, etc.)

## Deployment Steps

### 1. Apply the Configuration and Secrets

```bash
kubectl apply -f configmap.yaml
kubectl apply -f secret.yaml
```

### 2. Create the PostgreSQL Database

```bash
kubectl apply -f postgres-pvc.yaml
kubectl apply -f postgres-deployment.yaml
kubectl apply -f postgres-service.yaml
```

### 3. Run Database Migrations

```bash
kubectl apply -f migration-job.yaml
```

Wait for the migration job to complete:

```bash
kubectl get jobs edukita-lms-migrations
```

### 4. Deploy the Application

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

### 5. Set Up External Access

```bash
kubectl apply -f ingress.yaml
```

### 6. Configure Auto-scaling (Optional)

```bash
kubectl apply -f horizontal-pod-autoscaler.yaml
```

## Verifying the Deployment

### Check Deployment Status

```bash
kubectl get deployments
kubectl get pods
kubectl get services
```

### Check Application Logs

```bash
kubectl logs -l app=edukita-lms
```

### Access the Application

Once the Ingress is configured and DNS is propagated, you can access the application at:

```
https://lms.edukita.example.com
```

## Troubleshooting

### Checking Pod Status

```bash
kubectl describe pod <pod-name>
```

### Checking Service Status

```bash
kubectl describe service edukita-lms
```

### Checking Ingress Status

```bash
kubectl describe ingress edukita-lms-ingress
```

### Database Connection Issues

If the application cannot connect to the database:

1. Verify the PostgreSQL pod is running:
   ```bash
   kubectl get pods -l app=postgres
   ```

2. Check PostgreSQL logs:
   ```bash
   kubectl logs -l app=postgres
   ```

3. Verify the secret contains the correct database connection string:
   ```bash
   kubectl describe secret edukita-lms-secrets
   ```

## Updating the Application

To update the application with a new version:

1. Update your Docker image with the new version
2. Update the deployment:
   ```bash
   kubectl set image deployment/edukita-lms edukita-lms-app=edukita-lms:new-version
   ```
   Or apply the updated deployment.yaml file:
   ```bash
   kubectl apply -f deployment.yaml
   ```