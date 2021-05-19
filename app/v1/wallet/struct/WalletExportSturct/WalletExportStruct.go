package WalletExportSturct

type WalletExport struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Type       string `json:"Type"`
		PrivateKey string `json:"PrivateKey"`
	} `json:"result"`
	ID int `json:"id"`
}
