DOING WIKI BACK WIP - FORKER

ORIGINAL PROJECT https://github.com/dwilliamsuk/EthermineRPC-Go

# EthermineRPC-Go
Discord Rich Presence for https://ethermine.org/ written in Go

![EthermineRPC Example]![image](https://user-images.githubusercontent.com/52573108/161124175-1d204f48-a5f5-4b5e-b5c2-8e9bf8ea412a.png)

## Configuration
Edit config.yaml as follows:

| Key | Value |
| ------ | ------ |
| updateTime | Refresh Time (in minutes) |
| minerID | Your Ethereum Miner Address |
| clientID | Your RPC Client ID (Check Important Notes) |
| LargeImage | RPC Art Asset Name To Use (Check Important Notes) |
| LargeImageText | Large Image Asset Text |
| SmallImage | Small image under the big one |
| SmallImageText | Small Image Asset Text |
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
