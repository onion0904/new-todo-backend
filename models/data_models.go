package models

import (
	"gorm.io/gorm"
	_"fmt"
)

// type Status string

// const (
//     StatusNotStarted Status = "未着手"
//     StatusInProgress Status = "進行中"
//     StatusCompleted  Status = "完了"
// )

type Todo struct {
    gorm.Model
    Title string
	Status string
	Priority int
}

// func (t *Todo) BeforeSave(tx *gorm.DB) (err error) {
//     switch t.Status {
//     case StatusNotStarted, StatusInProgress, StatusCompleted:
//     default:
//         err = fmt.Errorf("invalid status: %s", t.Status)
//     }
//     return
// }
