package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JWT(t *testing.T) {
	isValid, err := ValidateTokenJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTIwOTU5NzYsImlhdCI6MTY5MjA5MjM3NiwiaWQiOjQsImlzcyI6ImVGaWxtIEFQSXMiLCJuYmYiOjE2OTIwOTIzNzYsInJvbGUiOiJNZW1iZXIiLCJ1c2VybmFtZSI6ImRpbWFzc2ZlYiJ9.QugG0q-GzC_piHiLqZwtee_KWWokqxLomlp3tnlcEmM")
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.True(t, isValid)
}
