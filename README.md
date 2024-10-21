# Open Payments plugin for shopd

Submission for [Interledger Hackathon 2024](https://interledger.org/summit/hackathon)

Implement an Open Payments plugin for [shopd](https://shopd.link/): *"Hugo compatible e-commerce plugin for the Caddy Web Server"*
- Mostly static e-commerce (one executable plus static site and DB files)
- [Hugo](https://gohugo.io/) shortcodes are the application entry points
- RESTful Hypermedia API (HATEOAS with [htmx](https://htmx.org/))
- Self-hostable and platform independent by making use of [Caddy](https://caddyserver.com/)
- [Open Payments](https://openpayments.dev/) continues this theme of empowering the user and conforming to web standards. It allows *"small merchants, to use payment initiation APIs directly against their customers’ payment accounts"* and *"improves upon existing Open Banking standards as defined in the UK, EU, and other jurisdictions"*

# cgi

Makes use of refactored **open-payments-example** (interactive grant) for [Payment Redirect with Open Payments API](https://github.com/mozey/open-payments-example)
```bash
# Install node modules
pnpm install
# Start interactive grant,
# redirect and wait for user interaction
node redirect.js
# Continue grant
node continue.js
```

These scripts are called from the Go plugin, see `cgi.go`

The plugin implements the `PaymentRedirect` interface in `hooks.go`

For this demo the `ContinueGrant` func confirm the order. Usually `ProcessMsg` would subscribe to webhook events. Usually the message handler listens for webhooks. For this demo the continuation is done manually after the customer interaction.


## Slides

Draft: [2024-10-Interledger-Hackathon](https://github.com/shopd/ih2024/blob/main/2024-10-Interledger-Hackathon.pdf)

Presentation: [2024-10-Interledger-Hackathon-2](https://github.com/shopd/ih2024/blob/main/2024-10-Interledger-Hackathon-2.pdf)


## Demo

This repo is a submodule, running the Go code requires a working shopd dev environment.

Watch the [demo video](https://youtu.be/C4_YlobWVJQ)

PS. If anyone would like to collaborate on the shopd beta release, please visit the site link above and click *"Join Waiting List"*
