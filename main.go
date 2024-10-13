package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const apiKey = "6cc2830faef093a8f3d3d435"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: currency-converter <amount> <from_currency> <to_currency> ")
		return
	}

	amountStr := os.Args[1]
	fromCurrency := os.Args[2]
	toCurrency := os.Args[3]

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	rate, err := getExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if rate == 0 {
		fmt.Println("Invalid currency code of APi error.")
	}

	convertedAmount := amount * rate
	fmt.Printf("%.2f %s is equal to %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)

}

func getExchangeRate(currency string, toCurrency string) (float64, error) {
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", apiKey, currency)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		ConversionRates map[string]float64 `json:"conversion_rates"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	rate, exists := result.ConversionRates[toCurrency]
	if !exists {
		return 0, fmt.Errorf("currency code %s not found", toCurrency)
	}
	return rate, nil
}
