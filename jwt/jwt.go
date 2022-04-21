package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const (
	// 定义token的有效时间
	TokenExpireDuration = time.Hour * 1
)

var SecretKey = []byte("123456")

// 请求的账户信息
type UserAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// 响应给客户端的数据
type ResponseToClient struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JWT中的payload中不要放重要数据,因为这部分数据通过Base64URL算法能反解出来
type Claims struct {
	// 自定义的`私有`数据,在payload中
	UserAccount
	// jwt的标准的claims
	jwt.StandardClaims
}

// 生成token
func GetToken(name, password, role string) (string, error) {
	c := Claims{
		UserAccount: UserAccount{
			Username: name,
			Password: password,
			Role:     role,
		},
		StandardClaims: jwt.StandardClaims{
			// token过期时间
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			// 签发人
			Issuer:  "captain",
			Subject: "jwt test",
		},
	}
	// 生成token,默认采用HMAC SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 加上签名(需要用到秘钥),生成完整的token
	return token.SignedString(SecretKey)
}

// 解析token
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func WriteToResponse(w http.ResponseWriter, code, message string, data interface{}) {
	var resp ResponseToClient
	resp.Code = code
	resp.Message = message
	resp.Data = data
	respJson, _ := json.Marshal(resp)
	// 设置响应为json格式
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	fmt.Fprintf(w, "%v\n", string(respJson))
}

// 默认处理函数
func defaultFunc(w http.ResponseWriter, r *http.Request) {

}

// 模拟用户登陆接口,响应json数据
// 目的是让客户端从服务器端获取token
// 处理请求逻辑之后,响应给客户的数据中包含新生成的token值
func AuthFunc(w http.ResponseWriter, r *http.Request) {
	var user UserAccount
	// 读取客户端请求的类容
	buf := make([]byte, 2048)
	n, _ := r.Body.Read(buf)
	// debug 调试请求的内容
	log.Println("json :", string(buf[:n]))
	// 将请求json数据解析出来
	err := json.Unmarshal(buf[:n], &user)
	// 如果解析错误,给客户端提示
	if err != nil {
		WriteToResponse(w, "400", err.Error(), "")
		return
	}
	// debug 调试解析之后的内容
	log.Println(user)
	// 模拟验证账户登录的逻辑(账户,密码都正确)
	if user.Username == "admin" && user.Password == "123456" {
		tokenString, _ := GetToken(user.Username, user.Password, user.Role)
		WriteToResponse(w, "200", "success", map[string]string{"token": tokenString})
		return
	} else {
		WriteToResponse(w, "400", "Account error", "")
		return
	}
	// 通过命令行的 CURL 测试
	// curl -X POST -H -H "Content-type:application/json" -d '{"username":"admin","password":"123456","role":"admin"}' http://127.0.0.1:8080/auth
}

// 模拟用户登陆(获取token)之后再请求某个接口,响应json数据
// 请求时,在请求的数据中包含token
// 包含token的载体可以是请求的url,也可以是请求头,也可以是请求体
func homeFunc(w http.ResponseWriter, r *http.Request) {
	// 此处我们模拟的token包含在请求头信息中
	authorH := r.Header.Get("Authorization")
	if authorH == "" {
		WriteToResponse(w, "401", "request header Authorization is null", "")
		return
	}
	// 将获取的Authorization 内容通过分割出来
	authorArr := strings.SplitN(authorH, " ", 2)
	// debug
	log.Println(authorArr)
	// Authorization的字符串通常是 "Bearer" 开头(可以理解为固定格式,标识使用承载模式),然后一个空格 再加上token的内容
	// Tips:  请求头中Authorization的内容直接是token也是可以的
	if len(authorArr) != 2 || authorArr[0] != "Bearer" {
		WriteToResponse(w, "402", "request header Authorization formal error", "")
		return
	}
	// 解析token这个字符串
	mc, err := ParseToken(authorArr[1])
	if err != nil {
		WriteToResponse(w, "403", err.Error(), "")
		return
	}
	// debug
	log.Println(mc)
	// 请求成功响应给客户端
	WriteToResponse(w, "200", "welcome to home", "")
}
func listFunc(w http.ResponseWriter, r *http.Request) {
	authorH := r.Header.Get("Authorization")
	if authorH == "" {
		WriteToResponse(w, "401", "request header Authorization is null", "")
		return
	}
	authorArr := strings.SplitN(authorH, " ", 2)
	// debug
	log.Println(authorArr)
	if len(authorArr) != 2 || authorArr[0] != "Bearer" {
		WriteToResponse(w, "402", "request header Authorization formal error", "")
		return
	}
	mc, err := ParseToken(authorArr[1])
	if err != nil {
		WriteToResponse(w, "403", err.Error(), "")
		return
	}
	// 模拟通过jwt确定权限的逻辑
	// 如果 用户角色是admin,那么就能访问该接口,否则就不允许
	if mc.Role == "admin" {
		WriteToResponse(w, "200", "列表数据", "")
		return
	}
	WriteToResponse(w, "404", "无权限访问", "")
}
func main() {
	http.HandleFunc("/", defaultFunc)
	http.HandleFunc("/auth", AuthFunc)
	http.HandleFunc("/home", homeFunc)
	http.HandleFunc("/list", listFunc)
	fmt.Println("start http server and listen 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer err : ", err)
	}
}
