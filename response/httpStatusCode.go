package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvalid = 20002 // Email
)

var msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParamInvalid: "Email is invalid",
}
