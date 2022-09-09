package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestGetType(t *testing.T) {
	//color := GetPassType("c_conf", 0)
	//fmt.Println(color)
	c := map[string]string{"d": "a"}
	//WriteHash("c.json", c)
	fmt.Println(len(c["a"]))
}

func TestGetTokenName(t *testing.T) {
	client := &http.Client{}
	//addr := strings.ToLower("0xb6E040C9ECAaE172a89bD561c5F73e1C48d28cd9")
	//var data = strings.NewReader(fmt.Sprintf(`{"operationName":"getNamesFromSubgraph","variables":{"address":"%s"},"query":"query getNamesFromSubgraph($address: String!) {\n  domains(first: 1000, where: {resolvedAddress: $address}) {\n    name\n    __typename\n  }\n}\n"}`, addr))
	req, err := http.NewRequest("GET", "http://127.0.0.1:30030/api/gettokenname", nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("tokenid", "1")
	q.Add("name", "did")
	req.URL.RawQuery = q.Encode()
	log.Println(req.URL.String())
	//req.Header.Set("authority", "api.thegraph.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("content-type", "application/json")
	//req.Header.Set("origin", "https://app.ens.domains")
	//req.Header.Set("referer", "https://app.ens.domains/")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ens := make(map[string]string)
	json.Unmarshal(bodyText, &ens)
	log.Println(ens)
}

func TestGetOpenSea(t *testing.T) {
	client := &http.Client{}
	//addr := strings.ToLower("0xb6E040C9ECAaE172a89bD561c5F73e1C48d28cd9")
	//var data = strings.NewReader(fmt.Sprintf(`{"operationName":"getNamesFromSubgraph","variables":{"address":"%s"},"query":"query getNamesFromSubgraph($address: String!) {\n  domains(first: 1000, where: {resolvedAddress: $address}) {\n    name\n    __typename\n  }\n}\n"}`, addr))
	req, err := http.NewRequest("GET", "https://api.opensea.io/api/v1/asset/0x22871b977AAe43d44FE50dF03f632134c3e3e490/1/?force_update=true", nil)
	if err != nil {
		log.Fatal(err)
	}
	//q := req.URL.Query()
	//q.Add("tokenid", "1")
	//q.Add("name", "did")
	//req.URL.RawQuery = q.Encode()
	//log.Println(req.URL.String())
	//req.Header.Set("authority", "api.thegraph.com")
	req.Header.Set("X-API-KEY", "")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("content-type", "application/json")
	//req.Header.Set("origin", "https://app.ens.domains")
	//req.Header.Set("referer", "https://app.ens.domains/")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bodyText))
	ens := make(map[string]bool)
	json.Unmarshal(bodyText, &ens)
	log.Println(ens)
}
func TestO(t *testing.T) {
	client := &http.Client{}
	var data = strings.NewReader(`{"id":"ToolbarAssetRefreshMutation","query":"mutation ToolbarAssetRefreshMutation(\n  $asset: AssetRelayID!\n) {\n  assets {\n    refresh(asset: $asset)\n  }\n}\n","variables":{"asset":"QXNzZXRUeXBlOjYxOTI0NzYyMQ=="}}`)
	req, err := http.NewRequest("POST", "https://opensea.io/__api/graphql/", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "opensea.io")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "_gcl_au=1.1.825860495.1657894523; ajs_user_id=0xd07bdb622a7e9d519a17c4c097bc479012761880; ajs_anonymous_id=a41ba829-1bb4-4c89-a6b3-faa3f721a6ef; cf_clearance=OVZLpKPuBxmb4_Iiirlfr5Ee.3R6BiZYp60eT.6VyT0-1658128208-0-150; opensea_logged_out=false; SFaSaPpU=A0y3NgKCAQAAjJVX0skfKdCSEtkcsrBZxeF0yFSiB8ofBljj0kd1_E0CYQ3xASPVJRuuci7ywH9eCOfvosJeCA|1|1|cbffaac61b6594bdd53b0eeb6a9a72bd5ea5d17e; amp_d28823=S75ubUZMBSaNb8YgTABnmL.MHhkMDdiZGI2MjJhN2U5ZDUxOWExN2M0YzA5N2JjNDc5MDEyNzYxODgw..1gb260vl7.1gb2610l1.s.10.1s; dpr=1; __os_session=eyJpZCI6ImIxNTk4MTcxLTVkMWUtNDM5ZC1hNzZlLTI5ZjI4OGFiYmQ0MyJ9; __os_session.sig=O4YNj-6uUE_GRBDkWJX-yqsGoyhSDOoFbIpWMshxrYM; _gid=GA1.2.100163569.1662630605; active_wallet=%22MetaMask%22; os-wallet={%22accounts%22:[{%22address%22:%220xd07bdb622a7e9d519a17c4c097bc479012761880%22%2C%22imageUrl%22:%22https://storage.googleapis.com/opensea-static/opensea-profile/11.png%22%2C%22nickname%22:null%2C%22relayId%22:%22QWNjb3VudFR5cGU6NDM1NjAzNTE2%22%2C%22isCompromised%22:false%2C%22isStaff%22:false%2C%22user%22:{%22relayId%22:%22VXNlclR5cGU6MzQwMjY3NjY=%22%2C%22username%22:%22asimovmeta%22%2C%22publicUsername%22:%22asimovmeta%22%2C%22hasAffirmativelyAcceptedOpenseaTerms%22:true%2C%22email%22:null}%2C%22metadata%22:{%22isBanned%22:false}%2C%22walletName%22:%22MetaMask%22}]%2C%22activeAccount%22:{%22address%22:%220xd07bdb622a7e9d519a17c4c097bc479012761880%22%2C%22imageUrl%22:%22https://storage.googleapis.com/opensea-static/opensea-profile/11.png%22%2C%22nickname%22:null%2C%22relayId%22:%22QWNjb3VudFR5cGU6NDM1NjAzNTE2%22%2C%22isCompromised%22:false%2C%22isStaff%22:false%2C%22user%22:{%22relayId%22:%22VXNlclR5cGU6MzQwMjY3NjY=%22%2C%22username%22:%22asimovmeta%22%2C%22publicUsername%22:%22asimovmeta%22%2C%22hasAffirmativelyAcceptedOpenseaTerms%22:true%2C%22email%22:null}%2C%22metadata%22:{%22isBanned%22:false}%2C%22walletName%22:%22MetaMask%22}}; csrftoken=DFU9M1p9AOVHXhkzWsMkHfejAH9uqDCr; __cf_bm=BjONg.as0CUT.v8t.zF0k2_sAW5vD5YMlswR646WpME-1662714763-0-AS7r/U3NyA7JArrr6fdoYmHae+SwakRPKTygFP4isRjlQuY6D3qft9qwMlFB/eeCGM2B9aI2tIDVnbeGOTukZic=; _ga_9VSBF2K4BX=GS1.1.1662710037.37.1.1662715283.0.0.0; _ga=GA1.1.595409396.1657894523; _uetsid=9ded5bc02f5b11eda8c1d519c849f125; _uetvid=9102e420044811edb0a98fb802b27146; sessionid=eyJzZXNzaW9uSWQiOiJiMTU5ODE3MS01ZDFlLTQzOWQtYTc2ZS0yOWYyODhhYmJkNDMifQ:1oWaDb:a4jOAlu4tYfFKGkDL6AuMBnqaxVv0nPf4d3T6uzB9Jo; amp_ddd6ec=htzXRoXbK9-uvIf63t8gey.MHhkMDdiZGI2MjJhN2U5ZDUxOWExN2M0YzA5N2JjNDc5MDEyNzYxODgw..1gcgmtu7r.1gcgotod0.17d.15u.2db; _dd_s=rum=0&expire=1662716314663")
	req.Header.Set("origin", "https://opensea.io")
	req.Header.Set("referer", "https://opensea.io/")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("x-app-id", "opensea-web")
	req.Header.Set("x-build-id", "aab93fd3316a5d979d54cf3a5c661be6e6969e7b")
	req.Header.Set("x-signed-query", "aca8db96dceca728d85769b3c7fc210d9d7f8eae6448bd2a81bbf8f7f6a48370")
	req.Header.Set("x-viewer-address", "0xd07bdb622a7e9d519a17c4c097bc479012761880")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
