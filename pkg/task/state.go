package task

import "fmt"

// State - enum of task states, статусы задач
// Backlog <-> Open <-> Progress -> Test -> Complete
//
//	^         |
//	|------- Return
type State int

const (
	// StateBacklog - task is in backlog, задача просто в бэклоге проекта
	StateBacklog State = 0
	// StateOpen - task is open for progress, scheduled to sprint, задача открыта, готова к работе над ней
	StateOpen State = 1
	// StateProgress - task is in progress now, задача сейчас в работе
	StateProgress State = 2
	// StateReturn - task, returned to continue work from [StateTest], задача, возвращенная на доработку со статуса [StateTest]
	StateReturn State = 4
	// StateTest - task is in somehow testing process, задача в каком-то состоянии приемки, внешних испытаний
	StateTest State = 8
	// StateComplete - task is successfully finished, задача успешно завершена
	StateComplete State = 16
)

func (tt State) String() string { return stateToString[tt] }

// ParseState - converts string to TaskType, преобразует строку в TaskType
func ParseState(s string) (State, error) {
	if tt, ok := stateFromString[s]; ok {
		return tt, nil
	}
	return State(0), fmt.Errorf("State with name [%s] not found", s)
}

var stateToString = map[State]string{
	StateBacklog:  "backlog",
	StateOpen:     "open",
	StateProgress: "progress",
	StateReturn:   "return",
	StateTest:     "test",
	StateComplete: "complete",
}
var stateFromString = func() map[string]State {
	var res map[string]State = make(map[string]State, len(stateToString))
	for k, v := range stateToString {
		res[v] = k
	}
	return res
}()
