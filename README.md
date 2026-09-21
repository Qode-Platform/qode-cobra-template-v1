# Cobra template

Provisioned from [`Qode-Platform/fleet-template-v1`](https://github.com/Qode-Platform/fleet-template-v1) - the fleet
lifecycle contract with a Cobra starter on top.

## Verified

Built and tested locally on Go 1.23.4 (toolchain auto-upgraded to 1.25):
`go build ./...` and `go test ./...` both pass.

## Fleet lifecycle

| step | command |
|---|---|
| install | `go mod download` |
| build | `go build -o ./.bin/app ./cmd/app` |
| start | `(none - not a service)` |

## Notes

- NOT A SERVICE: a CLI has nothing listening on $PORT, so START_CMD is empty and bin/run stops at the start step.
- Run it with: ./.bin/app greet world
