package val

import (
	"fmt"
	"regexp"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	regOnlyAlphabetic = regexp.MustCompile(`^[a-zA-Z]+$`)
	regOnlyNumber     = regexp.MustCompile(`^[0-9]+$`)
)

func IsOnlyAlphabetic(value string) error {
	if regOnlyAlphabetic.MatchString(value) {
		return nil
	}
	return fmt.Errorf("string contains non-alphabetic characters")
}
func IsOnlyNumber(value string) error {
	if regOnlyNumber.MatchString(value) {
		return nil
	}
	return fmt.Errorf("string contains non-number characters")
}

func FieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func InvalidArgumentsError(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{
		FieldViolations: violations,
	}
	statusInvalid := status.New(codes.InvalidArgument, "Invalid Parameter")
	statusDetail, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}
	return statusDetail.Err()
}
