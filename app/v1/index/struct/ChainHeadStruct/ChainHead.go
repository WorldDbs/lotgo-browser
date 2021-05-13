package ChainHeadStruct

type ChainHead struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Cids []struct {
			NAMING_FAILED string `json:"/"`
		} `json:"Cids"`
		Blocks []struct {
			Miner  string `json:"Miner"`
			Ticket struct {
				VRFProof string `json:"VRFProof"`
			} `json:"Ticket"`
			ElectionProof struct {
				WinCount int    `json:"WinCount"`
				VRFProof string `json:"VRFProof"`
			} `json:"ElectionProof"`
			BeaconEntries interface{} `json:"BeaconEntries"`
			WinPoStProof  []struct {
				PoStProof  int    `json:"PoStProof"`
				ProofBytes string `json:"ProofBytes"`
			} `json:"WinPoStProof"`
			Parents []struct {
				NAMING_FAILED string `json:"/"`
			} `json:"Parents"`
			ParentWeight    string `json:"ParentWeight"`
			Height          int    `json:"Height"`
			ParentStateRoot struct {
				NAMING_FAILED string `json:"/"`
			} `json:"ParentStateRoot"`
			ParentMessageReceipts struct {
				NAMING_FAILED string `json:"/"`
			} `json:"ParentMessageReceipts"`
			Messages struct {
				NAMING_FAILED string `json:"/"`
			} `json:"Messages"`
			BLSAggregate struct {
				Type int    `json:"Type"`
				Data string `json:"Data"`
			} `json:"BLSAggregate"`
			Timestamp int `json:"Timestamp"`
			BlockSig  struct {
				Type int    `json:"Type"`
				Data string `json:"Data"`
			} `json:"BlockSig"`
			ForkSignaling int    `json:"ForkSignaling"`
			ParentBaseFee string `json:"ParentBaseFee"`
		} `json:"Blocks"`
		Height int `json:"Height"`
	} `json:"result"`
	ID int `json:"id"`
}
