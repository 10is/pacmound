package pacmound

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Level02(gopher, python Agent) {
	loopCount, maxLoops := 0.0, 8.0*8.0

	level02(gopher, python, func(m *Maze, agentData *AgentData) bool {
		if !m.loop() || agentData.score >= (63-(loopCount*LivingCost))-0.001 || loopCount > maxLoops {
			return false
		}
		loopCount++
		return true
	})
}

func level02(gopher, python Agent, loop func(m *Maze, agentData *AgentData) bool) {
	const height, width = 9, 11
	maze := NewEmptyMaze(height, width)
	for x := 0; x < height; x++ {
		maze.setObsticle(x, 0)
		maze.setObsticle(x, width-1)
		for y := 0; y < width; y++ {
			maze.setObsticle(0, y)
			maze.setObsticle(height-1, y)
		}
	}

	for x := 1; x < height-1; x++ {
		for y := 1; y < width-1; y++ {
			if (y+2)%2 == 0 && (x+2)%2 == 0 {
				maze.setObsticle(x, y)
			} else {
				maze.setReward(x, y, 1)
			}
		}
	}

	maze[2][2].reward = 0
	maze[2][2].obsticle = false
	gopherData, err := maze.setAgent(2, 2, gopher)
	must(err)
	gopherData.t = 1
	gopher.SetScopeGetter(newScopeGetter(maze, gopherData))
	gopher.SetScoreGetter(gopherData.Score)

	for loop(&maze, gopherData) {
	}
}

func Level02Handler(getGopher, getPython AgentGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		maxLoops := 500
		loopLimit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil || loopLimit > maxLoops {
			loopLimit = maxLoops
		}
		loopCount := 0

		data := struct {
			MaxSteps int                `json:"maxSteps"`
			Scores   []float64          `json:"scores"`
			States   [][][]EncodedBlock `json:"states"`
		}{}
		data.MaxSteps = loopLimit

		gopher, python := getGopher(), getPython()
		level02(gopher, python, func(m *Maze, agentData *AgentData) bool {
			data.States = append(data.States, m.encodable())
			data.Scores = append(data.Scores, agentData.score)

			remReward := m.RemainingReward()

			if !m.loop() || remReward <= 0 || loopCount > loopLimit || agentData.dead {
				return false
			}
			loopCount++
			return true
		})

		json.NewEncoder(w).Encode(data)
	}
}
