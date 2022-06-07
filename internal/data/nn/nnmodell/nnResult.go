package nnmodell

type NNResult struct {
	Code int    `json:"Code,omitempty"`
	Msg  string `json:"Msg,omitempty"`
	Data string `json:"Data,omitempty"`
}

//public string xfsh { get; set; }
//public string xfmc { get; set; }
//public string xfdzdh { get; set; }
//public string xfyhzh { get; set; }
//public string je { get; set; }
//public string se { get; set; }
//public string hjje { get; set; }
//public string jqm { get; set; }
//public string jym { get; set; }
//public string bz { get; set; }
//public string zfbz { get; set; }
//public string hongc { get; set; }

//public class NNFPYZMXResult
//{
//public string spmc { get; set; }
//
//public string ggxh { get; set; }
//public string jldw { get; set; }
//public string spsl { get; set; }
//public string dj { get; set; }
//public string je { get; set; }
//public string slv { get; set; }
//public string se { get; set; }
//}
type NNFYZMXResult struct {
	Spmc string `json:"spmc,omitempty"`
	Spsl string `json:"spsl,omitempty"`
	Ggxh string `json:"ggxh,omitempty"`
	Jldw string `json:"jldw,omitempty"`
	Dj   string `json:"dj,omitempty"`
	Je   string `json:"je,omitempty"`
	Slv  string `json:"slv,omitempty"`
	Se   string `json:"se,omitempty"`
}
type NNFpYzResult struct {
	NNResult
	Fphm    string `json:"fphm,omitempty"`
	Fpdm    string `json:"fpdm,omitempty"`
	Fpzl    string `json:"fpzl,omitempty"`
	Kprq    string `json:"kprq,omitempty"`
	Gfmc    string `json:"gfmc,omitempty"`
	Gfsh    string `json:"gfsh,omitempty"`
	Gfdzdh  string `json:"gfdzdh,omitempty"`
	Gfyhzh  string `json:"gfyhzh,omitempty"`
	Xfsh    string `json:"xfsh,omitempty"`
	Xfmc    string `json:"xfmc,omitempty"`
	Xfdzdh  string `json:"xfdzdh,omitempty"`
	Xfyhzh  string `json:"xfyhzh,omitempty"`
	Je      string `json:"je,omitempty"`
	Se      string `json:"se,omitempty"`
	Hjje    string `json:"hjje,omitempty"`
	Jqm     string `json:"jqm,omitempty"`
	Jym     string `json:"jym,omitempty"`
	Bz      string `json:"bz,omitempty"`
	Zfbz    string `json:"zfbz,omitempty"`
	Hongc   string `json:"hongc,omitempty"`
	Details []NNFYZMXResult
}
