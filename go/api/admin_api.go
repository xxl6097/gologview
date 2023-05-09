package api

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
)

type GeneralResponse struct {
	Code int
	Msg  string
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
