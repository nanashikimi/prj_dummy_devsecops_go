# dummy-devsecops-go

Minimal Go HTTP service used as a playground for DevSecOps experiments.

## Note

This isn't a full project, just a training for learning DevSecOps.

## Endpoints

- `GET /healthz` – returns `ok`
- `GET /?name=YourName` – returns `hello, YourName`

## Run

```bash
go run ./...