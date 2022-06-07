package nnmodell

type nnbaseModel struct {
	Res      string `json:"res,omitempty"`
	Sign     string `json:"sign,omitempty"`
	Timetamp string `json:"timetamp,omitempty"`
	Appkey   string `json:"appkey,omitempty"`
	Custno   string `json:"custno,omitempty"`
}

type NnModel struct {
	nnbaseModel
	NnneedModel
}

type NnneedModel struct {
	IntefaceType string `json:"intefaceType,omitempty"`
	IsSandbox    string `json:"isSandbox,omitempty"`
	UserTax      string `json:"userTax,omitempty"`
	TaxNo        string `json:"taxNo,omitempty"`
	InvoiceCode  string `json:"invoiceCode,omitempty"`
	InvoiceNo    string `json:"invoiceNo,omitempty"`
	InvoiceDate  string `json:"invoiceDate,omitempty"`
	OptionField  string `json:"optionField,omitempty"`
}

type Sign struct {
	Data string `json:"data,omitempty"`
}
