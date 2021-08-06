package middleware

import "net/http"

type AuthMiddleWare struct {
	Next http.Handler
}

func (am *AuthMiddleWare)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	if am.Next==nil{
		am.Next=http.DefaultServeMux
	}
	auth:=r.Header.Get("Authorization")
	if auth!=""{
		am.Next.ServeHTTP(w,r)
	}else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}