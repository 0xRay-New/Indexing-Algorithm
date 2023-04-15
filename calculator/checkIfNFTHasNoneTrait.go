package calculator

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (data *CollectionData) CheckIfNFTHasNoneTrait(nftTraits []string) []string {
	var noneTraits []string
	for _, nftTrait := range data.AllTraits {
		if !contains(nftTraits, nftTrait) {
			noneTraits = append(noneTraits, nftTrait)
		}
	}
	return noneTraits
}
