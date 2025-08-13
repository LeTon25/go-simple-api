package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvalid = 20002 // Email
	ErrorCodeTokenInvalid = 20003 // Token
)

var msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParamInvalid: "Email is invalid",
	ErrorCodeTokenInvalid: "Token is invalid",
}
