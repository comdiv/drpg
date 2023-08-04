package task

import "githab.com/comdiv/drpg/pkg/actor"

/*
	Task

Everything in Software developer can be expressed as task.
Core thing. Main one.

В мире программирования по идее все выражается задачами.
Центральная вещь. Возможно главная
*/
type Task struct {
	// Id - unique identifier of task, уникальный идентификатор задачи
	Id string
	// Title - title or summary of task, заголовок, или краткое описание задачи
	Title string
	// Type - type of the task, тип задачи
	Type Type
	// Priority - priority of task, приоритет задач
	Priority Priority
	// Estimation - base time estimation, базовая оценка задачи в часах
	BaseEstimation int
	// Assignee - assignee person, executor of task, назначенный ответственный, исполнитель задачи
	Assignee actor.IActor
	// Author - author person, one who required a task, лицо, создавшее задачу и нуждающееся в ее выполнении
	Author actor.IActor
	// Level - hardness level of task, best fit actor's level, уровень сложности задачи, наилучшее соответствие уровню исполнителя
	Level actor.Level
	// CompleteRate - current rate of completeness, when 1.0 - can be moved to Test, текущий уровень выполнения, при достижении 1.0 может сместиться на Test
	CompleteRate float64
	// CoverageRate - current rate of test coverage for task, текущий уровень покрытия задачи тестами
	CoverageRate float64
	// TaskBugRate - task itself buggy rate, cause returns from tests, уровень забагованности самой задачи - влияет на вероятность не пройти тесты
	TaskBugRate float64
	// HiddenBugRate - rate of late-found bugs, caused by this task, уровень наличия скрытых багов, которые задача вносит в проект, потом прилетают как инциденты
	HiddenBugRate float64
}
