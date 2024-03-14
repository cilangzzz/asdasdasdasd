/**
  @creator: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/10/16
  @desc: //TODO
**/

package http

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"goDnParse/dao"
	"goDnParse/model"
	"log"
	"net/http"
	"strconv"
)

//var routerLock sync.RWMutex

//// GenerateDnLink 解析下载
//func GenerateDnLink(w http.ResponseWriter, r *http.Request) {
//	// 获取原始文件URL参数
//	url := r.URL.Query().Get("url")
//	ext := filepath.Ext(url)
//	timer := time.Now()
//	currentTime := timer.Format("0601020505.000")
//	dnDeal := func(w http.ResponseWriter, r *http.Request) {
//		// 请求文件
//
//		resp, err := http.Get(url)
//
//		if err != nil {
//			w.WriteHeader(500)
//			return
//		}
//		defer func(Body io.ReadCloser) {
//			err := Body.Close()
//			if err != nil {
//				return
//			}
//		}(resp.Body)
//
//		// 设置响应头
//		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
//		w.Header().Set("Content-Disposition", "attachment; filename="+currentTime+ext)
//
//		// 复制body到响应
//		_, err = io.Copy(w, resp.Body)
//		if err != nil {
//			return
//		}
//	}
//	tId, err := aes.GenerateTId(ext)
//	if err != nil {
//		return
//	}
//	routerLock.Lock()
//	http.HandleFunc("/dn/"+currentTime+tId+ext, dnDeal)
//	routerLock.Unlock()
//	StatusOk(w, currentTime+tId+ext)
//}
//
//// GenerateToken  生成密钥
//func GenerateToken(w http.ResponseWriter, r *http.Request) {
//	token, err := aes.GenerateToken()
//	if err != nil {
//		return
//	}
//	cookie := http.Cookie{
//		Name:  "token",
//		Value: token,
//	}
//	http.SetCookie(w, &cookie)
//	if err != nil {
//		return
//	}
//	_, err = w.Write([]byte(token))
//}
//
//// FileUploadHandle 文件上传
//func FileUploadHandle(w http.ResponseWriter, r *http.Request) {
//	token := r.URL.Query().Get("token")
//	if token != "www.cilang.buzz" {
//		StatusFatal(w, "身份校验失败")
//		return
//	}
//	file, handle, err := r.FormFile("file")
//	if err != nil {
//		StatusFatal(w, "获取失败")
//		return
//	}
//	defer func(file multipart.File) {
//		err := file.Close()
//		if err != nil {
//			return
//		}
//	}(file)
//	targetFile, err := os.Create("static/file/" + handle.Filename)
//	if err != nil {
//		StatusFatal(w, "创建失败")
//		return
//	}
//	defer func(targetFile *os.File) {
//		err := targetFile.Close()
//		if err != nil {
//			return
//		}
//	}(targetFile)
//	_, err = io.Copy(targetFile, file)
//	if err != nil {
//		StatusFatal(w, "写入失败")
//		return
//	}
//	StatusOk(w, "上传成功")
//}

// AddUser 添加用户
func AddUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token != "fake.com.user" {
		StatusOk(w, "身份校验失败")
		return
	}
	if err := r.ParseForm(); err != nil {
		StatusOk(w, "榜单解析失败")
		return
	}
	// 创建一个解析器
	decoder := schema.NewDecoder()
	// 解析URL中的参数到结构体
	//log.Println(r.URL.Query())
	//log.Println(r.PostForm)
	var user model.User
	if err := decoder.Decode(&user, r.PostForm); err != nil {
		StatusOk(w, "参数错误")
		return
	}
	//log.Println(user)
	err := dao.Insert(&user)
	if err != nil {
		log.Println(err)
		StatusOk(w, "添加错误")
		return
	}
	StatusOk(w, "添加成功,用户id"+strconv.FormatInt(user.Id, 10))
}

// GetUser 获取用户
func GetUser(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")
	intUuid, err := strconv.Atoi(uuid)
	if err != nil {
		StatusOk(w, "获取失败")
		return
	}
	user := new(model.User)
	user.Id = int64(intUuid)
	err = dao.GetUserByUid(user)
	if err != nil {
		StatusOk(w, "获取失败")
		return
	}
	res, err := json.Marshal(user)
	if err != nil {
		StatusOk(w, "获取失败")
		return
	}
	StatusOk(w, string(res))
}
