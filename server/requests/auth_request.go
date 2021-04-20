package requests

// LoginRequest ....
type LoginRequest struct {
	User      string `json:"user" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Recaptcha string `json:"recaptcha"`
	RemoteIP  string `json:"remote_ip" validate:"required"`
}
