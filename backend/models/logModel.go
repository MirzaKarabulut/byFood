package models

type RequestLog struct {
	ID        uint   `gorm:"primaryKey"`
	UserAgent string `json:"user_agent" validate:"required"`
	SessionID string `json:"sessionID" validate:"required"`
	DateTime  string `json:"datetime" validate:"required,datetime=2006-01-02 15:04:05"`
	Request   string `json:"request" validate:"required"`
	Response  string `json:"response" validate:"required"`
}

type EventLog struct {
	ID          uint   `gorm:"primaryKey"`
	EventName   string `json:"event_name" validate:"required"`
	Source      string `json:"source" validate:"required"`
	Tags        string `json:"tags" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserAgent   string `json:"user_agent" validate:"required"`
	SessionID   string `json:"sessionID" validate:"required"`
	DateTime    string `json:"datetime" validate:"required,datetime=2006-01-02 15:04:05"`
}
