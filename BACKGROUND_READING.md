# Background Reading

## Overview

[rafiki-low-level-intro](https://interledger.org/developers/blog/rafiki-low-level-intro/)

[openpayments overview](https://openpayments.dev/introduction/overview/) *"Open Payments aims to improve upon existing Open Banking standards as defined in the UK, EU, and other jurisdictions. Existing Open Banking ecosystems are dominated by aggregators and intermediaries, making it impossible for independent third-parties, such as small merchants, to use payment initiation APIs directly against their customers’ payment accounts. Open Payments allows for scenarios where clients can dynamically register and engage with the APIs without needing to pre-register with the ASE. This allows for a truly distributed and federated payment ecosystem with global reach and no dependence on any particular underlying account type or settlement system. Open Payments is also a significantly simpler standard with a small number of resource types and a more secure and powerful authorization protocol."*

Above is **what makes Open Payments better than Credit or Debit cards**. From the perspective of a small business, it works more like cash? But better than cash because they don't have to keep small change, do daily POS cashouts, or pay card processing transaction fees? Account Servicing Entities are registered Financial Service Providers, and are regulated entities within the countries they operate. ASEs provide wallet services to their customers, possibly for a flat monthly subscription and/or transaction fee?

[how-does-interledger-work](https://interledger.org/developers/get-started/#how-does-interledger-work) *"ILPv4 has three packet types - Prepare, Fulfill, and Reject... Interledger enables secure, multi-hop payments using Hashed Timelock Agreements. As of Interledger v4, these conditions are not enforced by the ledger, as it would be too costly and slow. Instead, participants in the network use these hashlocks to perform accounting with their peers"*

[Interledger Graveyard](https://interledger.org/developers/blog/simplifying-interledger-the-graveyard-of-possible-protocol-features/) *"These were 12 promising features, and countless hours were spent perfecting them. But they were sacrificed for simplicity’s sake... Getting people to agree on any standard is notoriously difficult, so we have worked to make Interledger as simple as possible... the core protocol would only be finished when there was nothing more to take out... we celebrate the life and death of these features for bringing us closer to payments interoperability"*


## Wallet Addresses

[wallet-addresses](https://openpayments.dev/introduction/wallet-addresses/) *"Not all URLs are wallet addresses, but all wallet addresses are URLs... If the URL is a wallet address, the response will provide details about the underlying Open Payments-enabled account... the authServer, assetCode, etc"*

```bash
# https://github.com/rs/curlie
curlie --request GET \
   --url https://ilp.interledger-test.dev/9734d1e \
   --header 'accept: application/json'
```


## TigerBeetle

[tigerbeetle](https://tigerbeetle.com/) uses Zig 😎

[tigerbeetle-go](https://github.com/tigerbeetle/tigerbeetle-go)


## Use cases

[Open payments use cases](https://interledger.org/summit/open-payments-use-cases)<sup>[1]</sup>
- *"Simplify cross-border and cross-currency transactions: Develop tools that make international and multi-currency payments effortless and accessible"*


## Test Network

Use these instead
- [Test Wallet](https://wallet.interledger-test.dev/)
- [Test E-Commerce](https://boutique.interledger-test.dev/)


## Rafiki

[Rafiki](https://rafiki.money) *"Test Wallet is a Rafiki playground, where you can add multiple accounts and make Interledger transactions with play money"* Don't use this one for the hackathon...

[where-did-rafiki-money-go](https://interledger.org/developers/blog/where-did-rafiki-money-go/)
*"Our aim with the new design was to entirely separate Test Wallet from the idea of Rafiki"*


## Examples

[Interactive payment example](https://github.com/interledger/open-payments-example)
