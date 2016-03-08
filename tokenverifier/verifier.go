package tokenverifier

type Token interface {
	GetId() string
	GetName() string
	GetPicture() string
}

type TokenVerifier interface {
	Verify(tokenString string) (Token, error)
}

type VerifierKey int

const (
	GOOGLE VerifierKey = iota
	FACEBOOK
)

var (
	verifiers = map[VerifierKey]TokenVerifier{
		GOOGLE:   &googleTokenVerifier{},
		FACEBOOK: &facebookTokenVerifier{},
	}
)

func GetVerifier(key VerifierKey) TokenVerifier {
	return verifiers[key]
}
