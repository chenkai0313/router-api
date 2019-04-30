package site

import "router-api/response"

type QueryResponse struct {
	response.BaseResponse
	Data Contents `json:"data"`
}

//type ResponseData struct {
//	Contents Contents `json:"data"`
//}

type Contents  []Content

type Content struct {
	Id          int64 `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedTime string `json:"created_time"`
}
