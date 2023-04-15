package calculator

import (
	"sort"
)
func (data *CollectionData) SortNFTS() {
	for k, v := range data.NFTScores {
		data.SortedNFTScores = append(data.SortedNFTScores, NFT{
			TokenID: k,
			Score: v["s"].(float64),
		})
	}
	sort.Slice(data.SortedNFTScores, func(i, j int) bool {
		return data.SortedNFTScores[i].Score > data.SortedNFTScores[j].Score
	})
}