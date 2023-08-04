package task

import (
	"time"

	"githab.com/comdiv/drpg/pkg/actor"
)

// Work - work element, элемент работы
type Work struct {
	// Duration - duration of work, продолжительность работы
	Duration time.Duration
	// WorkType - type of work, тип работы
	Type WorkType
	// Actor - actor of work, лицо осуществлявшее работы
	Actor actor.IActor
}

// WorkType - enum of work activity type, элемент работы
type WorkType string

const (
	// WorkDevelop - can increase completens and bugs, разработка, увеличивает уровень выполнения, а также уровень багов
	WorkDevelop WorkType = "dev"
	// WorkTest - can find new bugs - only increase task bug level, проверка, приемка - может выявить новые баги
	WorkTest WorkType = "test"
	// WorkFix - can decrease bug level - работы по устранению багов, могут уменьшать уровень багов, но не увеличивать completeness
	WorkFix WorkType = "fix"
)
