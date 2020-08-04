package martainUunExtractor

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
)

type customClaims map[string]interface{}

func between(value string, a string, b string) string {
	firstINX := strings.Index(value, a)
	if firstINX == -1 {
		return ""
	}
	value = value[firstINX+1:]
	lastINX := strings.Index(value, b)
	if lastINX == -1 {
		return ""
	}
	value = value[:lastINX]
	return value
}

func unmarshalClaims(raw []byte) map[string]interface{} {
	var claims customClaims
	err := json.Unmarshal(raw,&claims)
	if err != nil {
		log.Fatalln(err)
	}
	return claims
}

func decodeFromBase64(encodeSTR string) []byte {
	data,err :=base64.StdEncoding.DecodeString(encodeSTR)
	if err != nil{
		log.Fatalln(err)
	}
	return data
}