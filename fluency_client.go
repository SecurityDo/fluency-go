package fluency

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/SecurityDo/fluency-go/fsb"
)

type HTTPService struct {
	url         string
	client      *http.Client
	DebugFlag   bool
	token       string
	tokenHeader string //Fluencytoken

}

func NewHTTPService(url string) *HTTPService {

	s := &HTTPService{
		url: url,
	}
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 10 * time.Second,
	}
	timeout := time.Duration(600 * time.Second)
	s.client = &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
	return s
}

func (r *HTTPService) Close() {
	r.client.CloseIdleConnections()
}
func (r *HTTPService) GetClient() *http.Client {
	return r.client
}

func (r *HTTPService) GetUrl() string {
	return r.url
}

func (r *HTTPService) SetToken(token string) {
	r.token = token
}
func (r *HTTPService) SetTokenHeader(tokenHeader string) {
	r.tokenHeader = tokenHeader
}

func (r *HTTPService) Call(prefix string, functionName string, args *fsb.JNode) (result *fsb.JNode, err error) {
	remoteReq := new(fsb.CallRequest)
	remoteReq.Function = functionName
	remoteReq.Kargs = args
	reqStr, err := json.Marshal(remoteReq)
	if err != nil {
		return nil, fmt.Errorf("args is not a valid json structure")
	}
	//localUrl := fmt.Sprintf("http://localhost:%d/fsb/remoteCall", port)
	//apiPrefix := r.prefix
	//if(prefix != ""){apiPrefix=prefix}
	fullUrl := fmt.Sprintf("%s/%s/%s", r.url, prefix, functionName)
	if prefix == "" {
		fullUrl = fmt.Sprintf("%s/%s", r.url, functionName)
	}
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(reqStr))
	req.Header.Set("Content-Type", "application/json")
	if (r.token != "") && (r.tokenHeader != "") {
		req.Header.Set(r.tokenHeader, r.token)
	}

	if r.DebugFlag {
		if dump, err := httputil.DumpRequestOut(req, false); err == nil {
			fmt.Printf("Request:\n-----------------------------------------\n")
			fmt.Printf("%s", dump)
			pretty, _ := json.MarshalIndent(remoteReq, "", "   ")
			fmt.Printf("%s\n\n", pretty)
		}
	}
	resp, err := r.client.Do(req)
	if err != nil {
		fmt.Printf("Failed to call local http service %s: %s\n", r.url, err.Error())
		return result, fmt.Errorf("Failed to call local http service %s: %s\n", r.url, err.Error())
	}
	defer resp.Body.Close()
	if r.DebugFlag {
		if dump, err := httputil.DumpResponse(resp, true); err == nil {
			fmt.Printf("Response:\n-----------------------------------------\n")
			fmt.Printf("%s\n\n", dump)
		}
	}

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	if resp.StatusCode != 200 {
		fmt.Printf("HTTP ERROR from local http service %s: %s\n", r.url, resp.Status)
		return result, fmt.Errorf("HTTP Error from Local HTTP service %s: %s", r.url, resp.Status)
	}
	var res fsb.CallResponse
	//var obj map[string]interface{}
	//err = json.Unmarshal(body,&res)
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	err = decoder.Decode(&res)
	if err != nil {
		fmt.Printf("Failed to parse response body -> %s\n", err.Error())
		return result, fmt.Errorf("HTTP Error from Local HTTP service: %s", resp.Status)
	}
	if r.DebugFlag {
		pretty, _ := json.MarshalIndent(res, "", "   ")
		fmt.Printf("%s\n\n", pretty)
	}

	if res.Verdict == "ERROR" {
		fmt.Println("RPC call return with ERROR: ", res.Error)
		return result, fmt.Errorf("RPC call return with ERROR: %s", res.Error)
	} else if res.Verdict == "EXCEPTION" {
		fmt.Println("RPC call return with EXCEPTION: ", res.Exception)
		return result, fmt.Errorf("RPC call return with EXCEPTION: %s", res.Exception)
	}

	if len(res.Attachments) > 0 {
		fmt.Printf("Call has attachment\n")
		PrettyPrintJSON(res.Attachments[0])
	}

	return res.Response, nil

}

type FluencyClient struct {
	serviceClient *HTTPService
}

func NewFluencyClient(siteURL string, token string) *FluencyClient {
	s := &FluencyClient{
		serviceClient: NewHTTPService(siteURL),
	}
	s.serviceClient.SetTokenHeader("Fluencytoken")
	s.serviceClient.SetToken(token)
	return s

}

func (r *FluencyClient) GenericCall(prefix string, functionName string, jNode *fsb.JNode) (res *fsb.JNode, err error) {
	res, err = r.serviceClient.Call(prefix, functionName, jNode)

	if err != nil {
		fmt.Printf("%s/%s/%s call failed", r.serviceClient.GetUrl(), prefix, functionName)
		return res, err
	}
	return res, err
}
