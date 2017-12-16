package handlers

// this method will be invoked by the router
func CreditPointsHandler(request Request, response Response) {
	creditRequestHandler := creditRequestImpl{}
	creditRequestHandler.CreditPointsHandler(request, response)
}

// This method contains the handler logic.
// creditRequestImpl is used to share state with all methods
// Each method only executes its logic if previous method has not set error in the status
// Finally writeResponse will write success/error to the response based on the status
func (r *creditRequestImpl) CreditPointsHandler(request Request, response Response)  {
	r.Init(request, response).
		validateRequest(). //seprated out into credit_request_validator.go
		fetchWalletInfo().
		createCreditPointRequest().
		creditPoints().
		writeResponse()
}

type Request interface {
	GetParam(string) (string)
	GetParamAsInt(string) (int)
}

type Response interface {
	WriteResponse(string)
}

type CreditRequest struct {
	phoneNumber string
	value int
}
type WalletInfo struct {
	phoneNumber string
	walletId string
}
type CreditPointsRequest struct {
	phoneNumber string
	walletId string
	points int
}

type Status struct {
	httpStatusCode int
	message        string
}

func (r *Status) HasError() bool {
	return r.httpStatusCode >= 400
}

type creditRequestImpl struct {
	request             Request
	response 			Response
	creditPointsRequest CreditPointsRequest
	creditRequest       CreditRequest
	walletInfo          WalletInfo
	status              Status
}

func (r *creditRequestImpl) Init(request Request, response Response) (*creditRequestImpl) {

	r.request = request
	r.creditRequest = CreditRequest{
		phoneNumber: request.GetParam("phoneNumber"),
		value: request.GetParamAsInt("value"),
	}

	return r
}


func (r *creditRequestImpl) fetchWalletInfo() (*creditRequestImpl) {
	if r.status.HasError() {
		return r
	}
	// make wallet call.
	// do the wallet related validation
	// and update r.status

	// if not blacklisted
	r.walletInfo.phoneNumber = r.creditRequest.phoneNumber
	r.walletInfo.walletId = "wallet-id"
	return r
}

func (r *creditRequestImpl) createCreditPointRequest() (*creditRequestImpl) {

	if r.status.HasError() {
		return r
	}
	r.creditPointsRequest.phoneNumber = r.creditRequest.phoneNumber
	r.creditPointsRequest.walletId = r.walletInfo.walletId
	r.creditPointsRequest.points = r.creditRequest.value
	return r
}

func (r *creditRequestImpl) creditPoints() (*creditRequestImpl) {

	if r.status.HasError() {
		return r
	}

	// make credit points call.
	// and update r.status

	// if credit points was success
	r.status.httpStatusCode = 201

	return r
}

func (r *creditRequestImpl) writeResponse() {
	// write the response based on r.status
}
