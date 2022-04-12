package response

import "clean-hexa/business/content"

type GetContentByIDResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func NewGetContentByIDResponse(content content.Content) *GetContentByIDResponse {
	var contentResponse GetContentByIDResponse
	contentResponse.ID = content.ID
	contentResponse.Title = content.Name
	contentResponse.Desc = content.Description

	return &contentResponse
}
