package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Links struct {
	Homepage          []string `json:"homepage"`
	Whitepaper        string   `json:"whitepaper"`
	BlockchainSite    []string `json:"blockchain_site"`
	OfficialForumURL  []string `json:"official_forum_url"`
	TwitterScreenName string   `json:"twitter_screen_name"`
	FacebookUsername  string   `json:"facebook_username"`
	ReposURL          struct {
		Github []string `json:"github"`
	} `json:"repos_url"`
}

type CurrencyPrice struct {
	BDT float32 `json:"bdt"`
	ETH float32 `json:"eth"`
	USD float32 `json:"usd"`
}

type MarketData struct {
	CurrentPrice          CurrencyPrice `json:"current_price"`
	FullyDilutedValuation CurrencyPrice `json:"fully_diluted_valuation"`
	TotalVolume           CurrencyPrice `json:"total_volume"`
	High24                CurrencyPrice `json:"high_24h"`
	Low24                 CurrencyPrice `json:"low_24h"`
	PriceChange24h        float64       `json:"price_change_percentage_24h"`
	PriceChange7d         float64       `json:"price_change_percentage_7d"`
	PriceChange14d        float64       `json:"price_change_percentage_14d"`
	PriceChange30d        float64       `json:"price_change_percentage_30d"`
	PriceChange60d        float64       `json:"price_change_percentage_60d"`
	PriceChange200d       float64       `json:"price_change_percentage_200d"`
	PriceChange1y         float64       `json:"price_change_percentage_1y"`
	TotalSupply           float64       `json:"total_supply"`
	MaxSupply             float64       `json:"max_supply"`
	CirculatingSupply     float64       `json:"circulating_supply"`
}

type DeveloperData struct {
	Forks                        int `json:"forks"`
	Stars                        int `json:"stars"`
	Subscribers                  int `json:"subscribers"`
	TotalIssues                  int `json:"total_issues"`
	ClosedIssues                 int `json:"closed_issues"`
	PullRequestsMerged           int `json:"pull_requests_merged"`
	PullRequestContributors      int `json:"pull_request_contributors"`
	CodeAdditionsDeletions4Weeks struct {
		Additions int `json:"additions"`
		Deletions int `json:"deletions"`
	} `json:"code_additions_deletions_4_weeks"`
	CommitCount4Weeks              int   `json:"commit_count_4_weeks"`
	Last4WeeksCommitActivitySeries []int `json:"last_4_weeks_commit_activity_series"`
}

type Description struct {
	EN string `json:"en"`
}



type CoinInfo struct {
	ID               string         `json:"id"`
	Symbol           string         `json:"symbol"`
	Name             string         `json:"name"`
	WebSlug          string         `json:"web_slug"`
	Platforms        map[string]any `json:"platforms"`
	BlockTime        int            `json:"block_time_in_minutes"`
	HashingAlgorithm string         `json:"hashing_algorithm"`
	Categories       []string       `json:"categories"`
	Description      Description    `json:"description"`
	Links            Links          `json:"links"`
	GenesisDate      string         `json:"genesis_date"`
	MarketCapRank    int            `json:"market_cap_rank"`
	MarketData       MarketData     `json:"market_data"`
	DeveloperData    DeveloperData  `json:"developer_data"`
}

func fetchJSON(url string, target any) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to fetch data %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading the response body %v", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	return nil
}

func displayBasicInfo(data CoinInfo) {
	fmt.Println("\nName: ", data.Name)
	fmt.Println("ID: ", data.ID)
	fmt.Println("Symbol: ", data.Symbol)
	fmt.Println("Web Slug: ", data.WebSlug)
	fmt.Println("Platforms: ", data.Platforms)
	fmt.Println("BlockTime: ", data.BlockTime)
	fmt.Println("Hash Algorithm: ", data.HashingAlgorithm)
	fmt.Println("Genesis Data: ", data.GenesisDate)
	fmt.Println("Market Cap Rank: ", data.MarketCapRank)
}

func displayGeneralData(data CoinInfo) {
	fmt.Println("Name: ", data.Name)
	fmt.Println("Description:", data.Description.EN)
	fmt.Println("\nHomepage:", data.Links.Homepage)
	fmt.Println("\nWhitepaper:", data.Links.Whitepaper)
	fmt.Println("\nBlockchain Site:", data.Links.BlockchainSite)
	fmt.Println("\nOfficial Forum URL:", data.Links.OfficialForumURL)
	fmt.Println("\nTwitter Screen Name:", data.Links.TwitterScreenName)
	fmt.Println("\nFacebook Username:", data.Links.FacebookUsername)
	fmt.Println("\nGithub Repos:", data.Links.ReposURL.Github)
}

func displayPrices(data CoinInfo) {
	displayPrice(data.MarketData.CurrentPrice, "Current Price")
	displayPrice(data.MarketData.FullyDilutedValuation, "Fully Diluated Valuation")
	displayPrice(data.MarketData.TotalVolume, "Total Volume")
	displayPrice(data.MarketData.High24, "24 Hour High")
	displayPrice(data.MarketData.Low24, "24 Hour Low")

	fmt.Printf("\nTotal Supply: %f\n", data.MarketData.TotalSupply)
	fmt.Printf("Max Supply: %f\n", data.MarketData.MaxSupply)
	fmt.Printf("Circulating Supply: %f\n", data.MarketData.CirculatingSupply)

	fmt.Printf("\n Price Change Percentage 24 hour: %f\n", data.MarketData.PriceChange24h)
	fmt.Printf("\n Price Change Percentage 7 Days: %f\n", data.MarketData.PriceChange7d)
	fmt.Printf("\n Price Change Percentage 14 Days: %f\n", data.MarketData.PriceChange14d)
	fmt.Printf("\n Price Change Percentage 30 Days: %f\n", data.MarketData.PriceChange30d)
	fmt.Printf("\n Price Change Percentage 60 Days: %f\n", data.MarketData.PriceChange60d)
	fmt.Printf("\n Price Change Percentage 200 Days: %f\n", data.MarketData.PriceChange200d)
	fmt.Printf("\n Price Change Percentage 1 Year: %f\n", data.MarketData.PriceChange1y)
}

func displayPrice(data CurrencyPrice, title string) {
	fmt.Printf("\n%s", title+"\n")
	fmt.Printf("USD: %f\n", data.USD)
	fmt.Printf("BDT: %f\n", data.BDT)
	fmt.Printf("ETH: %f\n", data.ETH)
}

func displayDeveloperData(response DeveloperData) {
	fmt.Println("\nDeveloper Data: ")
	fmt.Println("Forks:", response.Forks)
	fmt.Println("Stars:", response.Stars)
	fmt.Println("Subscribers:", response.Subscribers)
	fmt.Println("Total Issues:", response.TotalIssues)
	fmt.Println("Closed Issues:", response.ClosedIssues)
	fmt.Println("Pull Requests Merged:", response.PullRequestsMerged)
	fmt.Println("Pull Request Contributors:", response.PullRequestContributors)
	fmt.Println("Code Additions (4 Weeks):", response.CodeAdditionsDeletions4Weeks.Additions)
	fmt.Println("Code Deletions (4 Weeks):", response.CodeAdditionsDeletions4Weeks.Deletions)
	fmt.Println("Commit Count (4 Weeks):", response.CommitCount4Weeks)
	fmt.Println("Last 4 Weeks Commit Activity Series:", response.Last4WeeksCommitActivitySeries)
}

func main() {
	coinName := flag.String("coin", "", "name of coin")
	operation := flag.String("type", "", "type of information")

	flag.Parse()

	url := "https://api.coingecko.com/api/v3/coins/" + *coinName

	var responseBody CoinInfo
	err := fetchJSON(url, &responseBody)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch *operation {
	case "basic":
		displayBasicInfo(responseBody)
	case "general":
		displayGeneralData(responseBody)
	case "price":
		displayPrices(responseBody)
	case "dev":
		displayDeveloperData(responseBody.DeveloperData)
	}
}
