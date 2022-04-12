package response

import "clean-hexa/business/content"

type GetContents struct {
	Contents []*GetContentByIDResponse `json:"contents"`
}

func NewGetContents(contentList []content.Content) *GetContents {
	contentResponses := make([]*GetContentByIDResponse, 0)

	for _, value := range contentList {
		contentResponses = append(contentResponses, NewGetContentByIDResponse(value))
	}

	return &GetContents{
		Contents: contentResponses,
	}
}
