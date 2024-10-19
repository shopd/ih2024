# Open Payments plugin for shopd

Submission for [Interledger Hackathon 2024](https://interledger.org/summit/hackathon)

Implement an Open Payments plugin for [shopd](https://shopd.link/): *"Hugo compatible e-commerce plugin for the Caddy Web Server"*


# cgi

Makes use of refactored **open-payments-example** (interactive grant) for [Payment Redirect with Open Payments API](https://github.com/mozey/open-payments-example)
```bash
# Install node modules
pnpm install
# Start interactive grant
node redirect.js
# Continue grant
node continue.js
```

These script are executed from the Go plugin, see `cgi.go`

The plugin implements the `PaymentRedirect` interface in `hooks.go`

The `ContinueGrant` func calls `ProcessMsg`. Usually the message handler listens for webhooks. For this demo the continuation is done manually after the customer interaction


## Demo

Link to demo video here...


## TODO

x Create shopd plugin (refactor existing yoco implementation)
x Implement PaymentRedirect hook
- Run commands for CGI scripts
- Check plugin works with shopd
- Record demo video
- Sleep zzz
