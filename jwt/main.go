package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// 解析 JWT 的函数
func ParseJWT(tokenString string) (*CustomClaims, error) {
	// 使用一个示例密钥（此处直接使用了原始字符串，若需要可以 Base64 解码）
	rawKey := []byte("allcoreisapowerfulmicroservicearchitectureupgradedandoptimizedfromacommercialproject") // 密钥示例

	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否是 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名算法: %v", token.Header["alg"])
		}
		return rawKey, nil
	})

	// 检查是否解析成功
	if err != nil {
		return nil, fmt.Errorf("解析失败: %v", err)
	}

	// 类型断言并检查 token 是否有效
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("无效的 Token 或 Claims 类型错误")
	}
}

func main() {
	// 示例 JWT
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXB0X2NvZGUiOiIwMDEiLCJ0ZW5hbnRfaWQiOiIwMDAwMDAiLCJhcmVhX25hbWUiOiIiLCJhY2NvdW50X3R5cGUiOiJsb25nIiwicm9sZV9uYW1lX3JlYWwiOiLmgLvpg6jnrqHnkIblkZgiLCJ1c2VyX25hbWUiOiJhZG1pbiIsImxhdGl0dWRlIjoiMzEuODUwOTc5IiwicmVhbF9uYW1lIjoi566h55CG5ZGYIiwiY2xpZW50X2lkIjoiYWxsY29yZSIsIm1ham9yIjoiIiwicm9sZV9pZCI6IjE1Njk5MzczOTYyMjYyNjkxODYiLCJzY29wZSI6WyJhbGwiXSwib2F1dGhfaWQiOiIiLCJleHAiOjE3NDQyODQ0MzIsImp0aSI6ImJhYjlkMDg5LTliODktNDdkNS04Nzg4LTBjZDBmMjAwYTUxZiIsImFwcF9jb2RlIjoiIiwibG9uZ2l0dWRlIjoiMTEzLjQ3ODY5OSIsImZyb21fdHlwZSI6IndlYiIsImVsZXZhdGlvbiI6bnVsbCwiYXJlYSI6bnVsbCwiY3JlYXRlX3RpbWUiOiIyMDIzLTAyLTE2IDEwOjAwOjUwIiwicm9sZV9ob21lX2NvbmZpZyI6IiIsImF2YXRhciI6IiIsImF1dGhvcml0aWVzIjpbInN5c19tYW5hZ2VyIl0sInJvbGVfbmFtZSI6InN5c19tYW5hZ2VyIiwibGljZW5zZSI6InBvd2VyZWQgYnkgYWxsY29yZSIsInBvc3RfaWQiOiIzNjNlZDMwMDk4M2UzMjQ3MDgyNzA0ODdkNTk3OTM5OSIsInVzZXJfaWQiOiI3NWViY2EyM2M5ZDQ3NzE3OGExNTY1N2IwYmQzYzgxZSIsIm5pY2tfbmFtZSI6IueuoeeQhuWRmCIsImRlcHRfY2F0ZWdvcnkiOiIxIiwiZGV0YWlsIjp7InR5cGUiOiJ3ZWIifSwiZGVwdF9pZCI6IjEiLCJyZWZfdHlwZSI6ImRlZmF1bHQiLCJhY2NvdW50IjoiYWRtaW4iLCJpc19mbHllciI6InllcyJ9.boCgDE-w1vi7Xzy237DETaeuopSb8YVfBUQ0MTcZtAw"

	// 调用解析函数
	c, err := ParseJWT(tokenString)
	if err != nil {
		log.Fatalf("解析失败: %v", err)
	}

	// 输出解析结果
	fmt.Printf("解析成功，用户ID: %s\n", c.UserID)
}
