package predictcontracts

// SignatureType represents the type of signature used for signing orders
type SignatureType int

// Signature type constants
const (
	SignatureTypeEOA SignatureType = 0 // Externally Owned Account
)

const COLLATERAL_TOKEN_DECIMALS = 18 // USDT on BNB has 18 decimals
const CONDITIONAL_TOKEN_DECIMALS = 6
