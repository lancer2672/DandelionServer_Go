package apicalls

type ExternalServiceApier interface {
	AuthApier
}
type ExternalService struct {
}

var ExternalServiceIns ExternalServiceApier

func init() {
	ExternalServiceIns = NewExternalService()
}
func NewExternalService() ExternalServiceApier {
	return &ExternalService{}
}
