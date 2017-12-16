package handlers

import "github.com/sirupsen/logrus"

func (r *creditRequestImpl) validateRequest() (*creditRequestImpl) {

	if len(r.creditRequest.phoneNumber) <= 0 ||
		r.creditRequest.value < 0 ||
		r.creditRequest.value > 1000 {
		logrus.Error(">>>>>>>", r.creditRequest.phoneNumber)
		r.status.message = "Request values are not valid"
		r.status.httpStatusCode = 402
	}

	return r
}