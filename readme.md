# kvstore
**kvstore** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve

set msg:
kvstored tx kvstore set-entry key helloworld --from alice --chain-id kvstore

get msg:
kvstored q kvstore get-entry key
```

