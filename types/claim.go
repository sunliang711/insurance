package types

// Claim TODO
type Claim struct {
	Name        string `json:"name"`
	CardType    string `json:"card_type"`
	Number      string `json:"number"`
	Sex         string `json:"sex"`
	Birthday    string `json:"birthday"`
	Phone       string `json:"phone"`
	City        string `json:"city"`
	ClaimDate   string `json:"claim_date"`
	ClaimReason string `json:"claim_reason"`
	ClaimType   string `json:"claim_type"`
	Description string `json:"description"`
}
