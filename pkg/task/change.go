package task

import (
	"time"

	"githab.com/comdiv/drpg/pkg/actor"
)

// ChangeHistoryItem - item of task-live history of changes, элемент истории изменения задачи
type ChangeHistoryItem struct {
	// Time - time of change, время изменения задачи
	Time time.Time
	// Type  - type of change, тип изменения задачи
	Type ChangeHistoryItemType
	// Actor - subject, who applies this change, действующее лицо, применившее изменения
	Actor actor.IActor
	// Value - attached value for changing, depends on [Type], связанное значение с изменениями, тип зависит от типа изменений
	Value any
}

// ChangeHistoryItemType - string enum of task history changing type, тип изменения в истории изменений задачи
type ChangeHistoryItemType string

const (
	ChangeType     ChangeHistoryItemType = "Type"
	ChangeState    ChangeHistoryItemType = "State"
	ChangeAssignee ChangeHistoryItemType = "Assignee"
	ChangeWork     ChangeHistoryItemType = "Work"
)
