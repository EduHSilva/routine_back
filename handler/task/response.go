package task

import (
	"time"
)

type ResponseData struct {
	ID           uint      `json:"id"`
	CreateAt     time.Time `json:"createAt"`
	UpdateAt     time.Time `json:"updateAt"`
	DeletedAt    time.Time `json:"deletedAt,omitempty"`
	Title        string    `json:"title"`
	Frequency    string    `json:"frequency"`
	Priority     string    `json:"priority"`
	DateStart    time.Time `json:"date_start,omitempty"`
	DateEnd      time.Time `json:"date_end,omitempty"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	CategoryName string    `json:"category"`
}

type ResponseTask struct {
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseTasks struct {
	Message string         `json:"message"`
	Data    []ResponseData `json:"data"`
}
