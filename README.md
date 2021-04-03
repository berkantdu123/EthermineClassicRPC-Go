# EthermineRPC-Go
Discord Rich Presence for https://ethermine.org/ written in Go

![EthermineRPC Example](https://i.jaffasite.ga/2021/04/etherminerpc-go.png)

## Configuration
Edit config.yaml as follows:

| Key | Value |
| ------ | ------ |
| updateTime | Refresh Time (in minutes) |
| minerID | Your Ethereum Miner Address |
| clientID | Your RPC Client ID (Check Important Notes) |
| LargeImage | RPC Art Asset Name To Use (Check Important Notes) |
| LargeImageText | Large Image Asset Text |

## Usage

```bash
go run main.go
```

__OR__

```bash
go build main.go
````
and use the built binary

## Important Notes
You __NEED__ to create an application [here](https://discord.com/developers/applications/) and grab your client ID, then create an "Art Asset" that will be displayed next to your status.

I'd recommend you __keep the refresh time at or above 2 minutes__ as __Ethermine only update information every 2 minutes__

I use [api.ethermine.org](https://ethermine.org) and [min-api.cryptocompare.com](https://cryptocompare.com) to get information on your mining stats and current pricing for Ethereum in USD. __I am not associated with either of these sites and cannot be held responsible for their content or actions.__
