# opa-auth
Authorisation using Open Policy Agent (OPA)

# how to run

```
go run main.go
```

Use rest.http file to:
- `GET /bundle/data` to see bundle data
- `GET /bundle/policy` to see rego policy
- `POST /opa/eval` to evaluate provided input