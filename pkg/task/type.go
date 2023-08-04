package task

import "fmt"

// Type - enum of task types, перечисление типов задач
type Type int

const (
	// TypeTask - usual task, обычная задача
	TypeTask Type = 0
	// TypeBug - bug fix task, задача на исправление ошибки
	TypeBug Type = 1
)

func (tt Type) String() string { return taskTypeToString[tt] }

// ParseType - converts string to TaskType, преобразует строку в TaskType
func ParseType(s string) (Type, error) {
	if tt, ok := taskTypeFromString[s]; ok {
		return tt, nil
	}
	return Type(0), fmt.Errorf("Type with name [%s] not found", s)
}

var taskTypeToString = map[Type]string{
	TypeTask: "task",
	TypeBug:  "bug",
}
var taskTypeFromString = func() map[string]Type {
	var res map[string]Type = make(map[string]Type, len(taskTypeToString))
	for k, v := range taskTypeToString {
		res[v] = k
	}
	return res
}()
