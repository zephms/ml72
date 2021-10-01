//v0.1.1
package mlg

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const ONLY4NETERR int = 0
const NETERRorTO int = 1

func pureGet(seg []string) string {
	str2get := "http://127.0.0.1:" + strings.Join(seg, "/")
	resp, err := http.Get(str2get)
	if err!=nil {
		//fmt.Println(err)
		return "er86756"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if err!=nil{
		return "er56578"
	}
	return "ok" + string(body)
}

func purePost(seg []string, data map[string]string) string {
	str2post := "http://127.0.0.1:" + strings.Join(seg, "/")
	urlvalues := url.Values{}
	for k := range data {
		urlvalues.Add(k, data[k])
	}
	resp, err := http.PostForm(str2post, urlvalues)
	if err!= nil {
		return "er45347"
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		return "er877587"
	}
	return "ok"+string(body)
}


func loopPost(seg []string, data map[string]string ,loopFor int) string {
	loopFlag := true
	get := "er68564"
	for loopFlag {
		get = purePost(seg, data)
		if loopFor == ONLY4NETERR{
			if get[0:2] != "er"{
				loopFlag = false
			}

		} else {
			if get[0:2] == "er" || get[0:4]=="okto"{
				loopFlag = true
			} else {
				loopFlag = false
			}
		}
		time.Sleep(50*time.Millisecond)
	}
	return get
}

func loopGet(seg []string, loopFor int) string {
	loopFlag := true
	get := "er684876"
	for loopFlag {
		get = pureGet(seg)
		if loopFor == ONLY4NETERR{
			if get[0:2] != "er"{
				loopFlag = false
			}

		} else {
			if get[0:2] == "er" || get[0:4]=="okto"{
				loopFlag = true
			} else {
				loopFlag = false
			}
		}
		time.Sleep(50*time.Millisecond)
	}
	return get
}