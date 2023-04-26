package request

type RequestPayload struct {
	Action         string                `json:"action"`
	Login          LoginPayload          `json:"login,omitempty"`
	Register       RegisterPayload       `json:"register,omitempty"`
	AccountConfirm AccountConfirmPayload `json:"accountConfirm,omitempty"`
	Mail           MailPayload           `json:"mail,omitempty"`
}
