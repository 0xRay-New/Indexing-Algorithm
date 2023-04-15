package calculator

func GetCollectionData(slug string) CollectionData {
	return CollectionData{
		Slug:                slug,
		Alleles:             make(map[string]int),
		AlleleCountPerTrait: make(map[string]int),
		NFTScores:           make(map[string]map[string]interface{}),
		MetatraitCounts:     make(map[int]int),
	}
}