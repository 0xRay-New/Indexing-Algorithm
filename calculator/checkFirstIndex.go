package calculator

import (
	"encoding/json"
	"log"
	"net/http"
)

func (data *CollectionData) CheckFirstIndex() error {
	for {
		req, err := http.NewRequest(http.MethodGet, "https://api.opensea.io/api/v1/assets?token_ids=0&order_direction=desc&offset=0&limit=2&collection="+data.Slug, nil)
		if err != nil {
			log.Println("error creating check first index request:", err)
			return err
		}
	
		req.Header.Set("X-API-KEY", "")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("error getting check first index response:", err)
			return err
		}
	
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			continue
		}
	
		var rdata AssetsResponse
		json.NewDecoder(resp.Body).Decode(&data)
		if len(rdata.Assets) == 0 {
			data.FirstIndex = 1
			data.LastIndex = data.Count
		} else {
			data.FirstIndex = 0
			data.LastIndex = data.Count - 1
		}
	
		return nil
	}
	
}