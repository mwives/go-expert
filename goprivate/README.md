## Downloading a Private Go Module from a Private Repository

To download a private Go module, set the `GOPRIVATE` environment variable to the repository URL:

```bash
export GOPRIVATE=github.com/your-organization/*
```

This ensures Go bypasses the public proxy for private modules.

## Git Authentication Issues

For troubleshooting Git authentication issues, see the [Go documentation](https://golang.org/doc/faq#git_https).

## Using `go mod vendor`

Use `go mod vendor` to store dependencies locally, ensuring offline access and consistent builds:

```bash
go mod vendor
```

Build with the vendor directory by adding:

```bash
go build -mod=vendor
```
