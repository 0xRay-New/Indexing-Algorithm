package common

type CollectionStruct struct {
	Address        string  `json:"address"`
	CollectionName string  `json:"collection_name"`
	ImageURL       string  `json:"image_url"`
	Fees           float64 `json:"fees"`
	IndexedData    []byte  `json:"indexed_data"`
	Count          int
	IndexCount     int
}

var IndexedCollections = []string{}

func CheckIfCollectionInIndexedCollections(collectionName string) bool {
	for _, collection := range IndexedCollections {
		if collection == collectionName {
			return true
		}
	}
	return false
}

var MongoDbConnectionURL = "mongodb url here"