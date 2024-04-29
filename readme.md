# kvstore
**kvstore** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve

msg:
kvstored tx kvstore set-entry key helloworld --from alice --chain-id kvstore

query:
kvstored q kvstore get-entry key
```
ru desc: 
* proto/kvstore/kvstore/entry.proto - cущность с которой работает модуль kvstore(Entry)
* proto/kvstore/kvstore/query.proto - получение даннных из блокчейна(GetEntry)
* proto/kvstore/kvstore/tx.proto - изменение состояния блокчейна(SetEntry)

для работы с kvstore можно взять пример из [osmosis](https://mapofzones.com/home/osmosis-1/overview?columnKey=ibcVolume&period=24h)
