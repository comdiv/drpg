package actor

// Developer - main actor class, основной класс акторов
type Developer struct {
	ActorBase
	// Experience - level of experience on current level, уровень опыта на текущем уровне навыков
	// when reach 1.0 - next level, когда достигает 1.0 - следующий уровень
	Experience float64
	// Level - level of developer
	Level Level
	// SmartRate - rate of smartness, more exp, less penalty on uplevel tasks,
	// общий уровень сообразительности - быстрее копится опыт, меньше штрафов за избыточную сложность
	// -1.0 - not smart, 2 less experiense, x2 more penalty on upper level
	// 0.0 - usual
	// 1.0 - x2 faster experience, x2 less penalty on level
	SmartRate float64
	// TestRate - rate of TDD skill, cause coverage and less bugs
	// 0.0 - no tests, all time to develop, but many bugs
	// 0.5 - writes test, but lost time, less bugs, but less speed
	// 1.0 - less bugs, with no penalty to whole time
	TestRate float64
	// SpeedRate - own execution speed rate
	// -1.0 - 2-time less speed than usual
	// 0.0 - usual speed
	// 1.0 - 2-time better speed than usual
	SpeedRate float64
	// ReviewRate - can be used only on other's task, decrease bugs, improve completenes, give exp to others (not own)
	// навыки ревьюера - позволяет увеличивать готовность, снижать баги и давать опыт в чужих задачах, не своих
	ReviewRate float64
}
