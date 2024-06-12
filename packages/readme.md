## Project Structure

This project includes three modules:

- `cmd/`: Contains the main entry point of the application.
- `math/`, `time/`: Additional packages providing specific functionalities.

## Using Go Modules and Workspaces

The `math` package is integrated into `cmd` using the `go mod edit -replace` command. The `time` package is included via Go workspaces, utilizing the appropriate `go work init` command.

When using workspaces, the `go mod tidy` command may encounter issues with the `time` package. In such cases, `go mod tidy -e` can be used to resolve them.
