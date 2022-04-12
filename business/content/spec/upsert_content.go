package spec

type UpsertContentSpec struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
}
