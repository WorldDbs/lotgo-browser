package MpoolPushMessageStruct

type MpoolPushMessage struct {
	Jsonrpc string `json:"jsonrpc"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Result struct {
		Message struct {
			Version    int    `json:"Version"`
			To         string `json:"To"`
			From       string `json:"From"`
			Nonce      int    `json:"Nonce"`
			Value      string `json:"Value"`
			GasLimit   int    `json:"GasLimit"`
			GasFeeCap  string `json:"GasFeeCap"`
			GasPremium string `json:"GasPremium"`
			Method     int    `json:"Method"`
			Params     string `json:"Params"`
			CID        struct {
				NAMING_FAILED string `json:"/"`
			} `json:"CID"`
		} `json:"Message"`
		Signature struct {
			Type int    `json:"Type"`
			Data string `json:"Data"`
		} `json:"Signature"`
		CID struct {
			NAMING_FAILED string `json:"/"`
		} `json:"CID"`
	} `json:"result"`
	ID int `json:"id"`
}
