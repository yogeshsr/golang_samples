package handlers

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestValidateUpdatesStatusWithErrorWhenPhoneNumberIsEmptyInCreditRequest(t *testing.T) {

	creditRequest := CreditRequest{
		value:       55,
		phoneNumber: "",
	}
	r := creditRequestImpl{
		creditRequest: creditRequest,
	}
	r.validateRequest()
	assert.Equal(t, 402, r.status.httpStatusCode)
}

func TestValidateUpdatesStatusWithSuccessWhenPhoneNumberIsEmptyInCreditRequest(t *testing.T)  {

	creditRequest := CreditRequest{
		value:       55,
		phoneNumber: "phone",
	}
	r := creditRequestImpl{
		creditRequest: creditRequest,
	}
	r.validateRequest()

	assert.False(t, r.status.HasError())
}
