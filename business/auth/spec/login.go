package spec

// InputLogin ...
type InputLogin struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
