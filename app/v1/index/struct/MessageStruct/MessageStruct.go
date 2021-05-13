package MessageStruct

type Message struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		BlsMessages []struct {
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
		} `json:"BlsMessages"`
		SecpkMessages []struct {
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
		} `json:"SecpkMessages"`
		Cids []struct {
			NAMING_FAILED string `json:"/"`
		} `json:"Cids"`
	} `json:"result"`
	ID int `json:"id"`
}
