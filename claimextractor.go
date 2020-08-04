package martainUunExtractor

import (
	"encoding/json"
	"github.com/google/martian"
	"github.com/google/martian/parse"
	"net/http"
)

func init() {
	parse.Register("header.uunToAuthModifier", extractorFromJSON)
}

type JwtExtractorModifier struct {
	Claim string
}

type JwtExtractorJSON struct {
	Claim string               `json:"claim"`
	Scope []parse.ModifierType `json:"scope"`
}

func (m *JwtExtractorModifier) ModifyRequest(req *http.Request) error {
	token := req.Header.Get("Authorization")
	if token == "" {
		return nil
	}
	rawToken := between(token, ".", ".")
	tokenClaims := decodeFromBase64(rawToken)
	claims := unmarshalClaims(tokenClaims)
	value := claims[m.Claim].(string)
	req.Header.Set("Authorization", value)
	return nil
}

func ExtractorNewModifier(claim string) martian.RequestModifier {
	return &JwtExtractorModifier{
		Claim: claim,
	}
}

func extractorFromJSON(b []byte) (*parse.Result, error) {
	msg := &JwtExtractorJSON{}

	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}
	return parse.NewResult(ExtractorNewModifier(msg.Claim), msg.Scope)
}
