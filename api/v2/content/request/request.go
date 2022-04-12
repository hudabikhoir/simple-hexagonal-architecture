package request

import "clean-hexa/business/content/spec"

type CreateContentRequestV2 struct {
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
}

func (req *CreateContentRequestV2) ToSpec() *spec.UpsertContentSpec {
	return &spec.UpsertContentSpec{
		Name:        req.Judul,
		Description: req.Deskripsi,
	}
}
