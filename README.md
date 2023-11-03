# Prototype of Image Compatibility Schema

## Run Schema Validation

```bash
go run cmd/compat/compat.go validate-schema examples/gpu-passthrough.json
```

## Run Host Validation
```bash
go run cmd/compat/compat.go validate-host examples/gpu-passthrough.json
```

## Print Version
```bash
go run cmd/compat/compat.go version
```
