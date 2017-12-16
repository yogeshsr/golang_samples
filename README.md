
## About
	Samples written in Golang for hands on.

# handlers
[credit_request_handler.go](handlers/credit_request_handler.go) demonstrates clutter free handler.
[credit_request_validator.go](handlers/credit_request_validator.go) & [credit_request_validator_test.go](handlers/credit_request_validator_test.go) which has seprated out the validator function.
Similarly other functions can be moved to individual files.

	Note: everthing is under the same pakage and CreditPointsHandler is the public entry point.

# Run tests
```
go test $(glide novendor)
```
