package action

import (
	"bytes"
	"encoding/json"
	"net/http"
	"signaling/src/comerrors"
	"signaling/src/framework"
	"strconv"
)

type comHttpResp struct {
	Errno  int         `json:"errNo"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

func writeJsonErrorResponse(cerr *comerrors.Errors, w http.ResponseWriter, cr *framework.ComRequest) {
	cr.Logger.AddNotice("errNo", strconv.Itoa(cerr.Errno()))
	cr.Logger.AddNotice("errMsg", cerr.Error())
	cr.Logger.Warningf("request process failed")

	resp := comHttpResp{
		Errno:  cerr.Errno(),
		ErrMsg: "process error",
	}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(resp)
	w.Write(buffer.Bytes())
}
