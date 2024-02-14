package cmdmanager

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER.")
	fmt.Println("Press 0(zero) to EXIT.")

	var prices []string
	for {
		var price string
		_, err := fmt.Scanln(&price)
		if err != nil {
			return nil, errors.New("failed to get the user input")
		}

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	marshal, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return errors.New("failed to convert data to JSON format")
	}

	fmt.Println(marshal)

	return nil
}

func New() CMDManager {
	return CMDManager{}
}
