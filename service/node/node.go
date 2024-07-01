package node

import (
	"bytes"
	"eiyaro-htmx-explorer/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func baseQuery(
	method string,
	path string,
	data []byte,
) ([]byte, error) {

	baseURL := os.Getenv("NODE_HOST")
	authToken := os.Getenv("NODE_ACCESS_TOKEN")
	authData := []byte(authToken)
	auth64 := base64.StdEncoding.EncodeToString(authData)

	request, err := http.NewRequest(method, fmt.Sprintf("%s/%s", baseURL, path), bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating http request", err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("basic %s", auth64))

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making http call", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll((resp.Body))
	if err != nil {
		fmt.Println("Failed to read response data", err)
		return nil, err
	}

	return body, nil
}

func rpcCall(path string, data []byte) (models.RPCMapResponse, error) {

	body, err := baseQuery("post", path, data)

	var responseData models.RPCMapResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Failed to parse response data", err)
		return nil, err
	}

	return responseData, nil
}

func GetLastBlock() (float64, error) {

	data, err := rpcCall("net-info", nil)

	if err != nil {
		fmt.Println("Failed to get block count", err)
		return -1, err
	}

	return data.M("data").V("highest_block").(float64), nil
}

func GetHashRate() (float64, error) {

	data, err := rpcCall("get-hash-rate", []byte("{}"))

	if err != nil {
		fmt.Println("Failed to get hash rate", err)
		return -1, err
	}

	raw := data.M("data").V("hash_rate").(float64)
	rate := raw / 1000000

	fmt.Println("HashRate:", rate)

	return rate, nil
}

func GetDifficulty() (float64, error) {

	netinfo, err := rpcCall("net-info", nil)

	if err != nil {
		fmt.Println("Failed to get net info", err)
		return -1, err
	}

	data, err := rpcCall("get-difficulty", []byte(
		fmt.Sprintf("{ \"block_height\": %.0f }", netinfo.M("data").V("current_block").(float64)),
	))

	if err != nil {
		fmt.Println("Failed to get hash rate", err)
		return -1, err
	}

	raw := data.M("data").V("difficulty").(string)
	diff, err := strconv.ParseFloat(raw, 64)

	if err != nil {
		fmt.Println("Failed to parse difficulty value", err)
		return -1, err
	}

	rate := diff / 1000000000

	return rate, nil
}

func GetPendingXfers() (float64, error) {

	data, err := rpcCall("list-unconfirmed-transactions", nil)

	if err != nil {
		fmt.Println("Failed to get pending count", err)
		return -1, err
	}

	return data.M("data").V("total").(float64), nil
}

func GetBlock(height int64) (models.Block, error) {

	var bockData models.RPCResponse[models.Block]

	body, err := baseQuery("post", "get-block", []byte(
		fmt.Sprintf("{ \"block_height\": %d }", height),
	))

	if err != nil {
		fmt.Println("Failed to get block data", err)
		return bockData.Data, err
	}

	err = json.Unmarshal(body, &bockData)
	if err != nil {
		fmt.Println("Failed to parse response data", err)
		return bockData.Data, err
	}

	return bockData.Data, nil
}

func GetBlocks(from int64, count int64) ([]models.Block, error) {

	var blocks []models.Block

	for i := from; i > (from - count); i-- {
		block, err := GetBlock(i)
		if err != nil {
			fmt.Println("Failed to get block", i)
			continue
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}
