package model

import "github.com/google/uuid"

type GetTopicParam struct {
	TopicID     uuid.UUID `json:"-"`
	TopicName   string    `json:"-"`
	ClassID     uuid.UUID `json:"-"`
	ClassName   string    `json:"-"`
	ClassTypeID uuid.UUID `json:"-"`
}
