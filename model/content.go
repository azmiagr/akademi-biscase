package model

type CreateContentRequest struct {
	TopicName   string `json:"topic_name" binding:"required"`
	Title       string `json:"title" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
	Description string `json:"description" binding:"required"`
	URL         string `json:"url" binding:"required"`
}
