package keeper

type DomainType struct {
	Owner      string `json:"owner"`
	Domain     string `json:"domain"`
	Parent     string `json:"parent"`
	Expiration uint64 `json:"expiration"`
}
