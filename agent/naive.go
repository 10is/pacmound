package agent

import "github.com/crhntr/comp469/pacman"

type Naive struct {
	dead  bool
	score pacman.ScoreGetter
	scope pacman.ScopeGetter

	searchDistance, unchangedScoreCount int
	previousScore                       float64

	warning error
}

func (p *Naive) Kill()                               { p.dead = true }
func (p *Naive) SetScoreGetter(f pacman.ScoreGetter) { p.score = f }
func (p *Naive) SetScopeGetter(f pacman.ScopeGetter) { p.scope = f }
func (p *Naive) Warning(err error)                   { p.warning = err }

func (p *Naive) CalculateIntent() pacman.Direction {
	if p.dead {
		return pacman.DirectionNone
	}
	// time.Sleep(time.Second / 10)

	d, maxReward := 0, 0.0
	for i, dir := range directions {
		x, y := dir.Transform()

		dirReward := 0.0
		for out := 0; out <= p.searchDistance; out++ {
			b := p.scope(x*out, y*out)
			if b == nil || b.IsOccupied() || b.IsObstructed() {
				continue
			}
			dirReward += b.Reward()
			if dirReward > maxReward {
				d, maxReward = i, dirReward
			}
		}
		maxReward = dirReward
	}

	if p.score() == p.previousScore {
		p.unchangedScoreCount++
		if p.unchangedScoreCount > 0 {
			p.searchDistance++
		}
	} else {
		p.previousScore = p.score()
		p.unchangedScoreCount = 0
	}
	return directions[d]
}
