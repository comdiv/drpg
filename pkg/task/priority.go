package task

import "fmt"

// Priority - enum of task priority level, перечисление приоритетов задач
type Priority int

const (
	// PriorityNormal - usual priority, обычный приоритет
	PriorityNormal Priority = 0
	// PriorityCritical - critical priority, критический приоритет
	PriorityCritical Priority = 16
	// PriorityMinor - низкий приоритет
	PriorityMinor Priority = -1
)

func (tt Priority) String() string { return priorityToString[tt] }

// ParsePriority - converts string to TaskType, преобразует строку в TaskType
func ParsePriority(s string) (Priority, error) {
	if tt, ok := priorityFromString[s]; ok {
		return tt, nil
	}
	return Priority(0), fmt.Errorf("TaskType with name [%s] not found", s)
}

var priorityToString = map[Priority]string{
	PriorityNormal:   "normal",
	PriorityCritical: "critical",
	PriorityMinor:    "minor",
}
var priorityFromString = func() map[string]Priority {
	var res map[string]Priority = make(map[string]Priority, len(priorityToString))
	for k, v := range priorityToString {
		res[v] = k
	}
	return res
}()
