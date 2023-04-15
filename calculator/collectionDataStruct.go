package calculator

type CollectionData struct {
	Name                string                            // collection name
	Slug                string                            // collection slug (e.g. boredapeyachtclub)
	Address             string                            // collection eth address
	ImageURL            string                            // collection image url
	Fees                int                               // collection fees
	Count               int                               // collection count
	FirstIndex          int                               // first index of collection (either 0 or 1)
	LastIndex           int                               // last index of collection (either count or count-1)
	AllAssets           []AssetsResponse                  // All assets in the collection
	AllTraits           []string                          // e.g. ["hat", "shirt", "background"]
	Alleles             map[string]int                    // e.g. {"hat_red": 10, "hat_blue": 5, "hat_green": 3}
	AlleleCountPerTrait map[string]int                    // e.g. {"hat": 15, "shirt": 10, "background": 5}
	NFTScores           map[string]map[string]interface{} // unsorted nft scores
	SortedNFTScores     []NFT
	MetatraitCounts     map[int]int // count of metadata traits: e.g. {1 trait: 10, 2 traits: 5, 3 traits: 3}
}

type NFT struct {
	TokenID string
	Score   float64
}