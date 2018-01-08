package maximus

import (
	"github.com/10is/pacmound"
)

type Agent struct {
	scope  pacmound.ScopeGetter
	reward pacmound.RewardGetter
}

func (agent *Agent) SetRewardGetter(f pacmound.RewardGetter) { agent.reward = f }
func (agent *Agent) SetScopeGetter(f pacmound.ScopeGetter)   { agent.scope = f }

func (agent *Agent) CalculateIntent() pacmound.Direction { return pacmound.Direction(0) }

func (agent *Agent) Kill()                    {}
func (agent *Agent) Damage(d pacmound.Damage) { /*log.Printf("Simple took damage: %s", d)*/ }
func (agent *Agent) Warning(err error)        { /*agent.warning = err*/ }
