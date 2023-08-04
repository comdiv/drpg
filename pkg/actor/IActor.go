package actor

// IActor - facade of any actor in game, фасад любого действующего лица в игре
type IActor interface {
	// GetId - read-accessor - id of actor, идентификатор актора
	GetId() string
	// GetName - read-accessor - readable name of actor, читаемое имя актора
	GetName() string
}

// ActorBase - minimal struct to express actor, can be mixin a-ka "base class"
type ActorBase struct {
	Id   string
	Name string
}

func (ab *ActorBase) GetId() string   { return ab.Id }
func (ab *ActorBase) GetName() string { return ab.Name }
