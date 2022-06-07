package data

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	v1 "ngBills/api/billsResult/v1"
	nnres "ngBills/internal/data/nn/nnmodell"
	unit "ngBills/internal/data/nn/unit"
	"strconv"
	"strings"
	"time"
)

func getInputModel(request *v1.GetTrueResultRequest) nnres.NnModel {
	inputModel := nnres.NnModel{}
	inputModel.Res = ""
	inputModel.Appkey = httpGetSign(request)

	//设置时间戳
	currentTicks := time.Now().Unix()
	//DateTime dtFrom = new DateTime(1970, 1, 1, 0, 0, 0, 0);
	//long currentMillis = (currentTicks - dtFrom.Ticks) / 10000;
	inputModel.Timetamp = strconv.FormatInt(currentTicks, 10)
	inputModel.Custno = request.CompanyNo
	//获取加密后的信
	sign := unit.Md532(inputModel.Appkey + inputModel.Custno + inputModel.Timetamp)
	inputModel.Sign = sign
	inputModel.TaxNo = request.TaxNo
	inputModel.InvoiceCode = request.InvoCode
	inputModel.InvoiceNo = request.InvoNo
	inputModel.InvoiceDate = request.InDate
	inputModel.UserTax = request.TaxNo
	if invoiceType(request.InvoCode) == true {
		inputModel.OptionField = request.CheckCode
	} else {
		inputModel.OptionField = request.Amount
	}

	bytes, _ := json.Marshal(nnres.NnneedModel{
		IntefaceType: "1",
		IsSandbox:    "0",
		UserTax:      inputModel.UserTax,
		TaxNo:        inputModel.TaxNo,
		InvoiceCode:  inputModel.InvoiceCode,
		InvoiceNo:    inputModel.InvoiceNo,
		InvoiceDate:  inputModel.InvoiceDate,
		OptionField:  inputModel.OptionField,
	})

	inputModel.Res = NNSerectSign(string(bytes), sign)

	return inputModel
}

func httpGetSign(request *v1.GetTrueResultRequest) string {
	apiUrl := "http://gateway.ns820.com/external/enterprise/appKey.do"
	data := url.Values{}
	data.Set("enpNo", request.CompanyNo)

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		panic(err)
	}
	fmt.Println(string(b))
	sign := nnres.Sign{}
	json.Unmarshal(b, &sign)

	return sign.Data
}

func HttpGet(request *v1.GetTrueResultRequest) nnres.NNResult {

	input := getInputModel(request)
	temurl := "http://gateway.ns820.com/i2d/sdp/api/Invoice/InvoiceInspection"

	//?appKey={0}&timestamp={1}
	data := url.Values{}
	data.Set("appKey", input.Appkey)
	data.Set("timestamp", input.Timetamp)

	u, err := url.ParseRequestURI(temurl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode

	fmt.Println(u.String())
	b := httpDo(u.String(), input.Res, input.Sign)
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		panic(err)
	}
	nnresModel := nnres.NNResult{}
	json.Unmarshal([]byte(b), &nnresModel)
	if nnresModel.Data != "" && nnresModel.Code == 1 {
		databute, _ := base64.StdEncoding.DecodeString(nnresModel.Data)
		datas, _ := unit.EcbDecrypt(databute, []byte(input.Sign[0:16]))
		nnresModel.Data = string(datas)
	}
	return nnresModel
}

func NNSerectSign(str string, key string) string {
	if len(key) > 16 {
		bytes := unit.EcbEncrypt([]byte(str), []byte(key[0:16]))
		return base64.StdEncoding.EncodeToString(bytes)
	} else {
		return ""
	}
}

func invoiceType(invoCode string) bool {
	isPP := false
	if len(invoCode) == 10 {
		flag := invoCode[7:8] //取值判断第8位的值
		if flag == "6" || flag == "3" {
			isPP = true
		}
	}
	if len(invoCode) == 12 {
		firstFlag := invoCode[0:1] //取值判断第1位的值
		sumFlag := firstFlag
		if sumFlag == "0" {
			if invoCode[10:12] != "13" {
				isPP = true
			} else {
				isPP = false
			}

		}
	}
	return isPP
}

func httpDo(url string, data string, sign string) string {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "text/plain")
	req.Header.Add("sign", sign)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}
