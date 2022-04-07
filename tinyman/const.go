package tinyman

const (
	// TinyManURL is a tiny man web url
	TinyManURL = "https://tinyman.org"

	// AlgodTestnetHost is the algorand test net url
	AlgodTestnetHost = "https://testnet-api.algonode.cloud"

	// AlgodMainnetHost is the algorand main net url
	AlgodMainnetHost = "https://mainnet-api.algonode.cloud"

	// TestnetValidatorAppIdV1_1 is the Tinyman test net validator app id version 1.1
	TestnetValidatorAppIdV1_1 int = 62368684

	// MainnetValidatorAppIdV1_1 is the Tinyman main net validator app id version 1.1
	MainnetValidatorAppIdV1_1 int = 552635992

	// TestnetValidatorAppId is an alias for the current Tinyman test net validator app id
	TestnetValidatorAppId = TestnetValidatorAppIdV1_1

	// MainnetValidatorAppId is an alias for the current Tinyman main net validator app id
	MainnetValidatorAppId = MainnetValidatorAppIdV1_1

	// SwapFixedInput is a fixed-input swap type
	SwapFixedInput = "fixed-input"

	// SwapFixedOutput is a fixed-output swap type
	SwapFixedOutput = "fixed-output"

	// AlgoTokenName is the algo token name
	AlgoTokenName = "Algo"

	// AlgoTokenUnitName is the algo token unit name
	AlgoTokenUnitName = "ALGO"

	// AlgoTokenDecimals is the algo token decimals
	AlgoTokenDecimals = 6

	TotalLiquidityTokens = 0xFFFFFFFFFFFFFFFF

	LiquidityTokenDecimals = 6

	LiquidityTokenUnitName = "TM1POOL"

	LiquidityAssetUnitName = "TMPOOL11"
)
