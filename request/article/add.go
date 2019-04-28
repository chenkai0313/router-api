package article

//手机号
type ArticleAddRequest struct {
	Title   string `json:"title"  form:"title"`
	Content string `json:"content"  form:"content"`
}
