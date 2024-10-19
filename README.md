# Open Payments plugin for shopd

Submission for [Interledger Hackathon 2024](https://interledger.org/summit/hackathon)

Implement an Open Payments plugin for [shopd](https://shopd.link/): *"Hugo compatible e-commerce plugin for the Caddy Web Server"*

Makes use of refactored **open-payments-example** with two step interaction for [Payment Redirect with Open Payments API](https://github.com/mozey/open-payments-example)


## Testing

Install dependencies with [PNPM](https://pnpm.io/)
```bash
cd cgi
pnpm install
```

Requires [NodeJS](https://nodejs.org/en) and [Go](https://go.dev/)

Create wallets and [developer key](https://wallet.interledger-test.dev/settings/developer-keys)

Configure env vars
```bash
cp .env.sample.sh .env
# Set env
```

Run tests
```bash
go test ./...
```


## TODO
x Create shopd plugin (refactor existing yoco implementation)
x Implement PaymentRedirect hook
- Run commands for CGI scripts
- Make test run stand-alone
- Check plugin works with shopd
- Record demo video
- Sleep zzz
