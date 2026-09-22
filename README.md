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

## Rule: everything under BASE_PATH

**This repo has no HTTP surface.** It is a Cobra CLI - it runs a command and exits, so there are no routes, no
redirects and no asset URLs, and `BASE_PATH` is unused today - the fleet may
still inject it, and the app is free to ignore it.

The rule applies to anything HTTP added later. The fleet serves apps behind a
proxy at `BASE_PATH=/direct/<agent>:<port>`, and that prefix is forwarded
**unchanged** - nginx does not strip it - so a handler registered at `/health`
is never reached. If you add an HTTP server here:

1. Add a `basePath()` helper that reads `BASE_PATH` and normalises it to `""`
   (standalone) or `/leading/no-trailing-slash`.
2. Mount every route on the group/router returned from it - never on the root
   router. Route patterns stay group-relative (`"/health"`, not the full
   prefixed path).
3. Prefix `basePath()` onto every absolute URL you hand a client: redirect
   targets, links and form actions in rendered HTML, and asset paths.

The gin, echo, fiber, chi and buffalo templates each carry a worked example of
this helper.
