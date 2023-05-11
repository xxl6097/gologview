package logview

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
)

type GeneralResponse struct {
	Code int
	Msg  string
}

type TreeData struct {
	Id       string `json:"id"`
	Label    string `json:"label"`
	Children []struct {
		Id    string `json:"id"`
		Label string `json:"label"`
	} `json:"children,omitempty"`
}

func NewService() (svr *GeneralResponse) {
	svr = &GeneralResponse{}
	return svr
}

type ProxyStatusResp struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Err        string `json:"err"`
	LocalAddr  string `json:"local_addr"`
	Plugin     string `json:"plugin"`
	RemoteAddr string `json:"remote_addr"`
}
type StatusResp map[string][]ProxyStatusResp

// GET api/status
func (svr *GeneralResponse) ApiStatus(w http.ResponseWriter, r *http.Request) {
	var (
		buf []byte
		res StatusResp = make(map[string][]ProxyStatusResp)
	)

	defer func() {
		buf, _ = json.Marshal(&res)
		_, _ = w.Write(buf)
	}()

	//ps := svr.ctl.pm.GetAllProxyStatus()
	//for _, status := range ps {
	//	res[status.Type] = append(res[status.Type], NewProxyStatusResp(status, svr.cfg.ServerAddr))
	//}

	a := ProxyStatusResp{
		Name:       "clife.pi.M2-5900:6200",
		Type:       "tcp",
		Status:     "running",
		Err:        "",
		LocalAddr:  "192.168.1.2:5900",
		Plugin:     "",
		RemoteAddr: "uuxia.cn:6200",
	}

	res["tcp"] = append(res["tcp"], a)

	for _, arrs := range res {
		if len(arrs) <= 1 {
			continue
		}
		sort.Slice(arrs, func(i, j int) bool {
			return strings.Compare(arrs[i].Name, arrs[j].Name) < 0
		})
	}
}

// GET api/status
func (svr *GeneralResponse) ApiFiles(w http.ResponseWriter, r *http.Request) {
	var (
		buf []byte
	)

	fmt.Println(w, "URL中的user请求参数的值是：", r.FormValue("path"))

	fmt.Println(w, "你发送的请求地址是：", r.URL.Path)
	fmt.Println(w, "你发送的请求地址后查询字符串是：", r.URL.RawQuery)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	// 必须，设置服务器支持的所有跨域请求的方法
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	w.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	dirs, err := os.ReadDir(r.FormValue("path"))
	if err != nil {
		return
	}

	files := []TreeData{}
	for _, dir := range dirs {
		f := TreeData{
			Id:    dir.Name(),
			Label: dir.Name(),
		}
		files = append(files, f)
	}

	defer func() {
		buf, _ = json.Marshal(&files)
		_, _ = w.Write(buf)
	}()
}

/**
[
    {
        "id": "node-1",
        "label": "node-1",
        "children": [
            {
                "id": "node-1-1",
                "label": "node-1-1"
            }
        ]
    },
    {
        "id": "node-2",
        "label": "node-2",
        "children": [
            {
                "id": "node-2-1",
                "label": "node-2-1"
            }
        ]
    },
    {
        "id": "node-3",
        "label": "node-3",
        "children": [
            {
                "id": "node-3-1",
                "label": "node-3-1"
            },
            {
                "id": "node-3-2",
                "label": "node-3-2"
            }
        ]
    }
]
*/
