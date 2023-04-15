package calculator

type AssetsResponse struct {
	Assets []struct {
		Traits []struct {
			TraitType string `json:"trait_type"`
			Value     string `json:"value"`
		} `json:"traits"`
		TokenID string `json:"token_id"`
	} `json:"assets"`
}