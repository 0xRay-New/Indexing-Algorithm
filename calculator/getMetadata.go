package calculator

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"
)

// gets basic collection data from opensea
func (data *CollectionData) GetMetadata() error {
	for {
		req, err := http.NewRequest(http.MethodGet, "https://api.opensea.io/api/v1/collection/"+data.Slug, nil)
		if err != nil {
			log.Println("error creating get metadata request:", err)
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("error getting get metadata response:", err)
			return err
		}

		defer resp.Body.Close()
		if resp.StatusCode == 404 {
			log.Println("Collection not found:", data.Slug)
			return errors.New("collection not found")
		}
		if resp.StatusCode != 200 {
			log.Println(resp.StatusCode)
			time.Sleep(1*time.Second)
			continue
		}

		traits := make(map[string]map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&traits)

		data.Name = traits["collection"]["name"].(string)
		data.Address = traits["collection"]["primary_asset_contracts"].([]interface{})[0].(map[string]interface{})["address"].(string)
		data.ImageURL = traits["collection"]["image_url"].(string)
		data.Fees = int(traits["collection"]["primary_asset_contracts"].([]interface{})[0].(map[string]interface{})["seller_fee_basis_points"].(float64)) + 250
		data.Count = int(traits["collection"]["stats"].(map[string]interface{})["count"].(float64))

		
		for i, v := range traits["collection"]["traits"].(map[string]interface{}) {
			count := 0
			data.AlleleCountPerTrait[strings.ToLower(i)] = len(v.(map[string]interface{}))
			for k, x := range v.(map[string]interface{}) {
				data.Alleles[strings.ToLower(i)+"_"+strings.ToLower(k)] = int(x.(float64))
				count += int(x.(float64))
			}
			if count != data.Count {
				log.Println("count mismatch:", count, data.Count)
				// means that this trait has a nonetype
				data.Alleles[strings.ToLower(i)+"_nonetrait"] = data.Count - count
				data.AlleleCountPerTrait[strings.ToLower(i)]++
			}
			data.AllTraits = append(data.AllTraits, strings.ToLower(i))
		}
		return nil
	}
}