package main

import (
  "fmt"
  "github.com/hugolgst/rich-go/client"
  "time"
  "net/http"
  "io/ioutil"
  "github.com/tidwall/gjson"
  "github.com/spf13/viper"
  "os"
  "strconv"
)

var (
  updateTime int
  minerID string
  clientID string
  LargeImage string
  LargeImageText string
)

func init() {
	viper.SetConfigName("config")
  viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, %s", err)
	}
  updateTime = viper.GetInt("config.updateTime")
  minerID = viper.GetString("config.minerID")
  clientID = viper.GetString("config.clientID")
  LargeImage = viper.GetString("config.LargeImage")
  LargeImageText = viper.GetString("config.LargeImageText")
  fmt.Println("-- EthermineRPC By Dan (https://github.com/TheJaffaMeme) -- ")
}

func main() {
  err := client.Login(clientID)
  if err != nil {
    panic(err)
  }

  ethAmount := getEthermineMonies()
  usdAmount := getEthPricing(ethAmount)

  ethAmountStr := fmt.Sprintf("%.5f", ethAmount)
  usdAmountStr := fmt.Sprintf("%.2f", usdAmount)

  fmt.Println("ETH = "+ ethAmountStr)
  fmt.Println("USD = "+ usdAmountStr)

  now := time.Now()
  later := time.Now().Add(time.Minute * time.Duration(updateTime))

  err = client.SetActivity(client.Activity{
	State:      usdAmountStr+" USD",
	Details:    ethAmountStr+" ETH",
	LargeImage: LargeImage,
	LargeText:  LargeImageText,
  Timestamps: &client.Timestamps{
		Start: &now,
    End: &later,
  },
  })

  if err != nil {
	panic(err)
  }

  fmt.Println("Sleeping for "+strconv.Itoa(updateTime)+" minutes...")
  time.Sleep(time.Minute * time.Duration(updateTime))
  main()
}

func owie() {
  fmt.Println("Error! Check Network Or Try Again Later.")
  time.Sleep(time.Second * time.Duration(10))
  os.Exit(3)
}

func getEthermineMonies() (float64) {
  url := "https://api.ethermine.org/miner/"+minerID+"/dashboard"
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {owie()}
  res, err := http.DefaultClient.Do(req)
  if err != nil {owie()}
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {owie()}
  resp := string(body)
  unpaidmoniesraw := gjson.Get(resp, "data.currentStatistics.unpaid").Num
  var unpaidmonies float64 = unpaidmoniesraw / 1000000000000000000
  return unpaidmonies
}

func getEthPricing(eth float64) (float64) {
  url := "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD"
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {owie()}
  res, err := http.DefaultClient.Do(req)
  if err != nil {owie()}
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {owie()}
  resp := string(body)
  ethtousd := gjson.Get(resp, "USD").Num
  ethinusd := eth * ethtousd
  return ethinusd
}
