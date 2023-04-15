package calculator

type CollectionResponse struct {
	Collection struct {
		PrimaryAssetContracts []struct {
			Address              string `json:"address"`
			SellerFeeBasisPoints int    `json:"seller_fee_basis_points"`
		} `json:"primary_asset_contracts"`
		Stats struct {
			Count float64 `json:"count"`
		} `json:"stats"`
		ImageURL string `json:"image_url"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
	} `json:"collection"`
}