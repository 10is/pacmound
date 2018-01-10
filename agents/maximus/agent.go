package maximus

import (
	"fmt"
	"math"

	"github.com/10is/pacmound"
)

type Agent struct {
	scope                    pacmound.ScopeGetter
	reward                   pacmound.RewardGetter
	directionToPreviousBlock pacmound.Direction
	stepIndex                int
}

func (agent *Agent) SetRewardGetter(f pacmound.RewardGetter) { agent.reward = f }
func (agent *Agent) SetScopeGetter(f pacmound.ScopeGetter)   { agent.scope = f }

func (agent *Agent) CalculateIntent() pacmound.Direction {
	intent := agent.addHawk()
	agent.directionToPreviousBlock = intent.TurnLeft().TurnLeft()
	return intent
}

func (agent *Agent) addHawk() pacmound.Direction {
	fmt.Print(agent.stepIndex)
	agent.stepIndex++
	agent.scope.DisplayRegion(3)

	directions := pacmound.Directions()
	rewards := make([]float64, len(directions))

	for i, direction := range directions {
		if agent.wouldcrash(direction) {
			rewards[i] = math.Inf(-1)
			continue
		}
		if agent.TryOpertunisticShortPath(direction) {
			return direction
		}
		if agent.wouldeat(direction) {
			return direction
		}
		if agent.directionToPreviousBlock == direction {
			rewards[i] -= 500
		}
	}
	fmt.Println(rewards)
	fmt.Println(directions)

	chosenDir := maxDirection(rewards, directions)
	fmt.Printf("chose: %s\n", chosenDir)
	return chosenDir
}

func maxDirection(rewards []float64, actions []pacmound.Direction) pacmound.Direction {
	maxDir, maxReward := 0, rewards[0]
	for i := 1; i < len(rewards); i++ {
		if rewards[i] > maxReward && !math.IsNaN(rewards[i]) {
			maxDir, maxReward = i, rewards[i]
		}
	}
	return actions[maxDir]
}

func (agent *Agent) Kill()                    { agent.stepIndex = 0 }
func (agent *Agent) Damage(d pacmound.Damage) { /*log.Printf("Simple took damage: %s", d)*/ }
func (agent *Agent) Warning(err error)        { /*agent.warning = err*/ }

func (agent *Agent) TryOpertunisticShortPath(direction pacmound.Direction) bool {
	xObsticle, yObsticle := direction.Transform()
	xObsticle, yObsticle = yObsticle*2, yObsticle*2
	obStructedBlock := agent.scope(xObsticle, yObsticle)
	if obStructedBlock == nil || !obStructedBlock.IsObstructed() {
		return false
	}

	xReward, yReward := direction.Transform()
	rewardBlock := agent.scope(xReward, yReward)
	return rewardBlock != nil && rewardBlock.Reward() > 0
}

func (agent *Agent) wouldcrash(direction pacmound.Direction) bool {
	xt, yt := direction.Transform()
	block := agent.scope(xt, yt)
	return block == nil || block.IsObstructed()
}

func (agent *Agent) wouldeat(direction pacmound.Direction) bool {
	xt, yt := direction.Transform()
	block := agent.scope(xt, yt)
	return block != nil && block.Reward() > 0
}

func (agent *Agent) IsOccupiedWithPython(direction pacmound.Direction) bool {
	xt, yt := direction.Transform()
	block := agent.scope(xt, yt)
	return block == nil || block.IsObstructed()
}
