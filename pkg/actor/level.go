package actor

import "fmt"

// Level - actor skill level  June -> June+ -> Middle -> Middle+ -> Senior
type Level int

const (
	// StateBacklog - task is in backlog, задача просто в бэклоге проекта
	LevelJunior Level = 0
	// StateOpen - task is open for progress, scheduled to sprint, задача открыта, готова к работе над ней
	LevelJuniorPlus Level = 1
	// StateProgress - task is in progress now, задача сейчас в работе
	LevelMiddle Level = 2
	// StateReturn - task, returned to continue work from [StateTest], задача, возвращенная на доработку со статуса [StateTest]
	LevelMiddlePlus Level = 4
	// StateTest - task is in somehow testing process, задача в каком-то состоянии приемки, внешних испытаний
	LevelSenior Level = 8
)

func (lvl Level) String() string { return levelToString[lvl] }

// ParseState - converts string to TaskType, преобразует строку в TaskType
func ParseState(s string) (Level, error) {
	if tt, ok := levelFromString[s]; ok {
		return tt, nil
	}
	return Level(0), fmt.Errorf("Level with name [%s] not found", s)
}

var levelToString = map[Level]string{
	LevelJunior:     "junior",
	LevelJuniorPlus: "junior+",
	LevelMiddle:     "middle",
	LevelMiddlePlus: "middle+",
	LevelSenior:     "senior",
}
var levelFromString = func() map[string]Level {
	var res map[string]Level = make(map[string]Level, len(levelToString))
	for k, v := range levelToString {
		res[v] = k
	}
	return res
}()
