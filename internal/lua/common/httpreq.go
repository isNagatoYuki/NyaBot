package common

import (
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/internal/utils"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"log"
)

type HttpReqFunc struct{}

func (h *HttpReqFunc) Get(url string, params string) string {
	resp, err := utils.GetRequest(url, params)
	if err != nil {
		log.Println("[Lua HTTP] Error while sending GET request: ", err)
		return ""
	}
	return string(resp)
}

func (h *HttpReqFunc) GetJson(url string, params string) map[string]interface{} {
	var data map[string]interface{}
	resp, err := utils.GetRequest(url, params)
	if err != nil {
		log.Println("[Lua HTTP] Error while sending GET request: ", err)
		return nil
	}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println("[Lua HTTP] Error while parsing JSON: ", err)
		return nil
	}
	return data
}

func (h *HttpReqFunc) Post(url string, params interface{}) string {
	resp, err := utils.PostRequest(url, params)
	if err != nil {
		log.Println("[Lua HTTP] Error while sending POST request: ", err)
		return ""
	}
	return string(resp)
}

func (h *HttpReqFunc) PostJson(url string, params interface{}) map[string]interface{} {
	var data map[string]interface{}
	resp, err := utils.PostRequest(url, params)
	if err != nil {
		log.Println("[Lua HTTP] Error while sending POST request: ", err)
		return nil
	}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println("[Lua HTTP] Error while parsing JSON: ", err)
		return nil
	}
	return data
}

func HttpReq(L *lua.LState) {
	var HttpReqFunc = &HttpReqFunc{}
	L.SetGlobal("HttpReq", luar.New(L, HttpReqFunc))
}