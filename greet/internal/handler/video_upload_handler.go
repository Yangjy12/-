package handler

import (
	"Bytecode_Project/greet/Models"
	"Bytecode_Project/greet/helper"
	"Bytecode_Project/greet/internal/logic"
	"Bytecode_Project/greet/internal/svc"
	"Bytecode_Project/greet/internal/types"
	"crypto/md5"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

func VideoUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		file, fileHeader, err2 := r.FormFile("file")
		if err2!=nil {
			return
		}
		//判断文件是否存在
		bytes := make([]byte, fileHeader.Size)
		_, err2 = file.Read(bytes)
		if err2!=nil {
			return
		}
		hash:=fmt.Sprintf("%x",md5.Sum(bytes))
		rp:=new(Models.Video)
		get, err2 := svcCtx.Engine.Where("hash=?", hash).Get(rp)
		if err2!=nil {
			return
		}
		if get {
			httpx.OkJson(w,&types.VideoUploadReply{StatusCode: 0,StatusMsg: "文件已经存在"})
			return
		}

		//往cos中存储文件
		cosPath, err2 := helper.CosUpload(r)
		if err2!=nil {
			return
		}
		id := r.Header.Get("id")
		userid, err2 := strconv.ParseInt(id, 10, 64)
		username := r.Header.Get("name")
		//往 logic 中传递req
		req.Hash=hash
		req.Path=cosPath
		req.UserId=userid
		req.UserName=username
		l := logic.NewVideoUploadLogic(r.Context(), svcCtx)
		resp, err := l.VideoUpload(&req,r.Header.Get("name"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
