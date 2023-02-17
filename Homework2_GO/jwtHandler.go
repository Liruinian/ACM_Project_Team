package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

// 自定义加密结构体数据需要实现jwt.StandardClaims 方法
type Test struct {
	Name    string `json:"name"`
	Context string `json:"context"`
	jwt.StandardClaims
}

var SignKey = []byte("1123654gfdhgfd")

func jwttester() {
	a := Test{
		Name:    "v_6556u",
		Context: "哈哈哈",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),      // 开始生效时间
			ExpiresAt: time.Now().Unix() + 20, // 过期时间 （5秒后过期）
			Issuer:    "xuzhi",                //签发人
		},
	}
	//jwt.NewWithClaims 生成token
	//jwt.SigningMethodHS256  加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a) //使用自定义体a
	tokenString, err := token.SignedString(SignKey)
	if err != nil {
		fmt.Println("error")
		log.Println(err.Error())
		return
	}
	fmt.Println("加密后的token字符串:", tokenString)

	//解密
	s, err := ParseToken(tokenString)
	if err != nil {
		log.Panicln(err.Error())
	}
	fmt.Println("解密：", s)

	fmt.Println("=====================================")

	tokenA := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //自带的jwt.MapClaims
		"name":    "司大帅",
		"content": "asdfg",
		"exp":     time.Now().Unix(),
		"iss":     "lisa",
	})

	//token.SignedString 将token 转换为string

	tokenStringA, err := tokenA.SignedString(SignKey)
	if err != nil {
		fmt.Println("error")
		log.Println(err.Error())
		return
	}

	fmt.Println("加密后的token字符串", tokenStringA)

	//在这里如果也使用jwt.ParseWithClaims的话，第二个参数就写jwt.MapClaims{}
	//例如jwt.ParseWithClaims(tokenString, jwt.MapClaims{},func(t *jwt.Token) (interface{}, error){}

	fmt.Println(jwtDecode(tokenStringA, SignKey))

}

func jwtDecode(token string, SignKey []byte) map[string]interface{} {
	token_decode, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return token_decode.Claims.(jwt.MapClaims)
}

func ParseToken(tokenString string) (*Test, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Test{}, func(token *jwt.Token) (i interface{}, err error) {
		return SignKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 令牌有效
	if claims, ok := token.Claims.(*Test); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
