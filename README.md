<!--
parent:
  order: false
-->

<div align="center">
  <h1> FishCake Service </h1>
</div>

<div align="center">
  <a href="https://github.com/FishcakeLab/fishcake-service/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/FishcakeLab/fishcake-service.svg" />
  </a>
  <a href="https://github.com/FishcakeLab/fishcake-service/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/FishcakeLab/fishcake-service.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/FishcakeLab/fishcake-service">
    <img alt="GoDoc" src="https://godoc.org/github.com/FishcakeLab/fishcake-service?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/FishcakeLab/fishcake-service">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/FishcakeLab/fishcake-service"/>
  </a>
</div>

This is a contract event listen service for fishcake which can sync polygon blocks and parse fishcake contract's events in blocks, It has three modules, they are synchronizer, event and rest api.

- Synchronizer: Synchronize contract events related to FishCake and insert them into the database with batches
- Event: Extract contract events from the data and process them according to business requirements and then enter them into the business database
- Api: Provide API for FishCake frontend


**Note**: Requires [Go 1.21+](https://golang.org/dl/)

## Architecture

[![FishCake](https://github.com/FishcakeLab/fishcake-service/blob/main/assets/fishcake.png)](https://github.com/FishcakeLab)


## Installation

For prerequisites and detailed build instructions please read the [Installation](https://github.com/FishcakeLab/fishcake-service/) instructions. Once the dependencies are installed, run:

```bash
make 
```

Or check out the latest [release](https://github.com/FishcakeLab/fishcake-service).


## Setup And Run

- Config yaml 
```
migrations: "./migrations"
#polygon_rpc: "https://polygon-mumbai.infura.io/v3/f22a689113714ae8a41d02b18e1b1f7d"
polygon_rpc: "https://polygon-mainnet.g.alchemy.com/v2/CIZMD_P-HQWswRx6-n8uar5ONYLN02sb"
rpc_url: "193.203.215.185:8189"
polygon_chain_id: "80001"
http_host: "127.0.0.1"
http_port: 8087
db_host: "127.0.0.1"
db_port: 5432
db_name: "fishcake"
db_user: "fishcake"
db_password: ""
metrics_host:
metrics_port:
start_block: 58013559
eventStartBlock: 58013559
contracts:
    - "0xE967Df5072C00051a7Df26D8E736b8769fB991b5"
    - "0x4907e1fA441673CC004415784430179B8B4938Cf"
```

- Create database and migrate

```
create database fishcake;
./fishcake migrate
```

- Run service

```
./fishcake index
```

- Run api 

```
./fishcake api
```

