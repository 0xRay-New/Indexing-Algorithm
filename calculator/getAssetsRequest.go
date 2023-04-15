package calculator

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)




func (data *CollectionData) GetAssetsRequest() error {
	
	if data.Count <= 10000 {
		for counter := 0; counter <= data.LastIndex; counter += 50 {
			log.Println("indexing", counter, "for collection", data.Slug)
			for {
				url := "https://api.opensea.io/api/v1/assets?limit=50&order_direction=asc&offset="+strconv.Itoa(counter)+"&collection="+data.Slug
	
				req, err := http.NewRequest(http.MethodGet, url, nil)
				if err != nil {
					log.Println("error creating get assets request:", err)
					return err
				}
		
				req.Header.Set("X-API-KEY", "")
		
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Println("error getting get assets response:", err)
					return err
				}
		
				defer resp.Body.Close()
				if resp.StatusCode != 200 {
					continue
				}
	
				var rdata AssetsResponse
				json.NewDecoder(resp.Body).Decode(&rdata)
	
				data.AllAssets = append(data.AllAssets, rdata)
	
				for i := range rdata.Assets {
					data.MetatraitCounts[len(rdata.Assets[i].Traits)]++
				}
	
				break
			}
		}
	} else {
		log.Println("Using backup indexing method for collection ", data.Slug)
		count := data.Count
		val := 0
		for count > 0 {
			for {
				time.Sleep(1*time.Second)
				url := "https://api.opensea.io/api/v1/assets?limit=30&order_direction=asc&collection="+data.Slug
				for i := val; i < val+30; i++ {
					url += "&token_ids="+strconv.Itoa(i)
				}
				val += 30
				req, err := http.NewRequest(http.MethodGet, url, nil)
				if err != nil {
					log.Println("error creating get assets request:", err)
					return err
				}
		
				req.Header.Set("X-API-KEY", "")
		
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Println("error getting get assets response:", err)
					return err
				}
		
				defer resp.Body.Close()
				if resp.StatusCode != 200 {
					log.Println("error getting get assets response:", resp.StatusCode)
					val -= 30
					continue
				}

				var rdata AssetsResponse
				json.NewDecoder(resp.Body).Decode(&rdata)
	
				data.AllAssets = append(data.AllAssets, rdata)
	
				for i := range rdata.Assets {
					data.MetatraitCounts[len(rdata.Assets[i].Traits)]++
				}

				count -= len(rdata.Assets)

				log.Println("Got ", val, "for collection", data.Slug)

				break
				
			}
		}
	}
	

	if data.MetatraitCounts[0] != 0 {
		delete(data.MetatraitCounts, 0)
	}
	top := float64((len(data.Alleles)+len(data.MetatraitCounts)) * data.Count)
	for i := range data.AllAssets{
		for j := range data.AllAssets[i].Assets {
			nftScore := float64(0)
			nftTraits := []string{}
			
			for x := range data.AllAssets[i].Assets[j].Traits {
				
				nftTraits = append(nftTraits, strings.ToLower(data.AllAssets[i].Assets[j].Traits[x].TraitType))

				alpha := float64(data.Alleles[strings.ToLower(data.AllAssets[i].Assets[j].Traits[x].TraitType+"_"+data.AllAssets[i].Assets[j].Traits[x].Value)])
				beta := float64(data.AlleleCountPerTrait[strings.ToLower(data.AllAssets[i].Assets[j].Traits[x].TraitType)])
				
				bottom := alpha*beta
				nftScore += top/bottom
			}

			// check nonetraits
			missingTraits := data.CheckIfNFTHasNoneTrait(nftTraits)
			for y := range missingTraits {
				alpha := float64(data.Alleles[strings.ToLower(missingTraits[y]+"_nonetrait")])
				beta := float64(data.AlleleCountPerTrait[strings.ToLower(missingTraits[y])])
				bottom := alpha*beta
				
				nftScore += top/bottom
			}
			alpha := float64(data.MetatraitCounts[len(data.AllAssets[i].Assets[j].Traits)])
			beta := float64(len(data.MetatraitCounts))
			bottom := alpha*beta
			nftScore += top/bottom
			if nftScore > 0 && !math.IsInf(nftScore, 1) {
				data.NFTScores[data.AllAssets[i].Assets[j].TokenID] = make(map[string]interface{})
				data.NFTScores[data.AllAssets[i].Assets[j].TokenID]["s"] = math.Round(nftScore*100)/100
			}
		}
	}
	log.Println("done")
	return nil
}
