# Api docs

## 1.get wallet balance
#### request
```
127.0.0.1:8087/v1/chain_info/balance?address=0xAb7a82c767aC2D5F3Db625e91afE553F0Dae25f4
```
#### response
```
{
    "success": true,
    "msg": "success",
    "obj": {
        "pol_balance": "2342319981366312",
        "usdt_balance": "29900000",
        "fcc_balance": "180000000"
    },
    "code": "0000"
}
```
- pol_balance: polygon token balance
- usdt_balance: usdt token balance
- fcc_balance: fcc token balance


## 2.get sign info

#### request
```
127.0.0.1:8087/v1/chain_info/sign_info?address=0xAb7a82c767aC2D5F3Db625e91afE553F0Dae25f4
```

#### response
```
{
    "success": true,
    "msg": "success",
    "obj": {
        "nonce": "2",
        "native_token_gas_limit": "21000",
        "erc20_token_gas_limit": "150000",
        "max_fee_per_gas": "36097846677",
        "max_priority_fee_per_gas": "30000000000",
        "gas_price": "36097846677"
    },
    "code": "0000"
}
```

- nonce: address nonce
- native_token_gas_limit: pol token transfer gaslimit
- erc20_token_gas_limit: erc20 token transfer gaslimit

## 3 send tx to polygon network

#### request
```
127.0.0.1:8087/v1/chain_info/send_tx?rawTx=0x02f8b38189478505d21dba008506c088e200830165cc94c2132d05d31c914a87c6611c10748aeb04b58e8f80b844a9059cbb000000000000000000000000ab7a82c767ac2d5f3db625e91afe553f0dae25f40000000000000000000000000000000000000000000000000000000000989680c080a0aaa42b2a0f7cd2096f68b71923ddccfb02e8c192d5fd4294aff6ad2c4513d0a1a03e93b5676d0bd5c232253d89b56b0ee7a289d737c3509a4b61c96efaa07bc8a4
```
- rawTx: offline sign transaction
#### response
```
{
    "success": false,
    "msg": "send tx fail, please try again later",
    "obj": null,
    "code": "0002"
}
```

## 4.submit wallet address for airdrop

#### request
```
127.0.0.1:8087/v1/wallet/createWallet?address=0xee2E207D30383430a815390431298EBa3c1C8c20&device=1222121111
```

#### response
```
{
    "success": true,
    "msg": "success",
    "obj": "0xb49e5019de0d90d279cbeb2f7e2da86c5ed8fe1c298e0c7a4d85c141e01a3bb9",
    "code": "0000"
}
```