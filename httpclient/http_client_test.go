package httpclient

import (
	"crypto/md5"
	"myhttp/interfaces/mock"
	"testing"
)

func init() {
	httpClient = &mock.MockIHTTPClient{}
}

func TestValidateAndAppendSchemeWithoutScheme(t *testing.T) {

	address := "google.com"
	expectedAddress := "http://google.com"

	fixedAddress, err := validateAndAppendScheme(address)

	validateErrforNil(err, t)
	validateExpectedAndReceived(expectedAddress, fixedAddress, t)
}

func TestValidateAndAppendSchemeWithScheme(t *testing.T) {

	address := "http://google.com"
	expectedAddress := "http://google.com"

	fixedAddress, err := validateAndAppendScheme(address)

	validateErrforNil(err, t)
	validateExpectedAndReceived(expectedAddress, fixedAddress, t)
}

func TestValidateAndAppendSchemeInvalidAddress(t *testing.T) {

	address := " "
	_, err := validateAndAppendScheme(address)
	validateErrforNotNil(err, t)
}

func TestGetResponseMd5(t *testing.T) {

	address := "http://noerror.com"
	md5Checksum, err := getResponseMd5(address)
	expectedMd5Checksum := md5.Sum([]byte("body response from noerror.com"))

	validateErrforNil(err, t)

	if md5Checksum != expectedMd5Checksum {
		t.Errorf("failed expected [%x], received [%x]", expectedMd5Checksum, md5Checksum)
	}
}

func TestGetResponseMd5BadStatus(t *testing.T) {

	address := "http://error.com"
	md5Checksum, err := getResponseMd5(address)
	expectedMd5Checksum := [16]byte{}

	validateErrforNotNil(err, t)

	if md5Checksum != expectedMd5Checksum {
		t.Errorf("failed expected [%x], received [%x]", expectedMd5Checksum, md5Checksum)
	}
}

func validateExpectedAndReceived(expected string, received string, t *testing.T) {
	if expected != received {
		t.Errorf("failed expected [%s], received [%s]", expected, received)
	}
}

func validateErrforNil(err error, t *testing.T) {
	if err != nil {
		t.Errorf("failed : [error : %v]", err)
	}
}

func validateErrforNotNil(err error, t *testing.T) {
	if err == nil {
		t.Errorf("failed : err expected")
	}
}
