package day16

import (
	"container/heap"
	"math"

	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/geom"
)

type path struct {
	pos   geom.Point
	dir   geom.Direction
	score int
}

type mazeTile struct {
	char       rune
	bestScores [4]int
	parents    [4][]state
}

type state struct {
	pos geom.Point
	dir geom.Direction
}

var dirs [4][3]geom.Direction = [4][3]geom.Direction{
	{geom.Up, geom.Left, geom.Right},
	{geom.Right, geom.Up, geom.Down},
	{geom.Down, geom.Right, geom.Left},
	{geom.Left, geom.Down, geom.Up},
}

func PartOne(lines []string, extras ...any) any {
	bestScore, _ := traverseMaze(lines)
	return bestScore
}

func PartTwo(lines []string, extras ...any) any {
	_, bestSpots := traverseMaze(lines)
	return bestSpots
}

func createMaze(lines []string) (maze *ds.Grid[mazeTile], start geom.Point) {
	maze = ds.NewGrid[mazeTile](len(lines[0]), len(lines))

	for y, line := range lines {
		for x, char := range line {
			tile := maze.Get(x, y)
			tile.char = char
			for i := range tile.bestScores {
				tile.bestScores[i] = math.MaxInt
			}
			if char == 'S' {
				start = geom.Point{X: x, Y: y}
			}
		}
	}
	return maze, start
}

func traverseMaze(lines []string) (int, int) {
	maze, start := createMaze(lines)
	priorityQueue := ds.NewPriorityQueue(func(a, b path) bool {
		return a.score < b.score
	})

	heap.Push(priorityQueue, path{pos: start, dir: geom.Right, score: 0})

	var endStates []state
	bestScore := math.MaxInt

	for priorityQueue.Len() > 0 {
		item := heap.Pop(priorityQueue).(path)
		if item.score > bestScore {
			break
		}

		tile := maze.Get(item.pos.XY())
		cardinalIndex := item.dir.GetCardinalIndex()

		if tile.bestScores[cardinalIndex] < item.score {
			continue
		}

		if tile.char == 'E' {
			if item.score < bestScore {
				bestScore = item.score
				endStates = append(endStates[:0], state{pos: item.pos, dir: item.dir})
			} else if item.score == bestScore {
				endStates = append(endStates, state{pos: item.pos, dir: item.dir})
			}
			continue
		}

		for _, nextDir := range dirs[cardinalIndex] {
			next := item.pos.Add(nextDir)
			nextScore := item.score + 1
			if nextDir != item.dir {
				nextScore += 1000
			}

			nextTile := maze.Get(next.XY())
			nextCardinalIndex := nextDir.GetCardinalIndex()
			nextBestScore := nextTile.bestScores[nextCardinalIndex]

			if nextTile.char == '#' || nextScore > nextBestScore {
				continue
			}

			if nextScore < nextBestScore {
				nextTile.parents[nextCardinalIndex] = nextTile.parents[nextCardinalIndex][:0]
				nextTile.bestScores[nextCardinalIndex] = nextScore
				heap.Push(priorityQueue, path{next, nextDir, nextScore})
			}
			nextTile.parents[nextCardinalIndex] = append(
				nextTile.parents[nextCardinalIndex],
				state{pos: item.pos, dir: item.dir},
			)
		}
	}

	bestPathsPoints := ds.NewSet[geom.Point]()
	if len(endStates) > 0 {
		visited := ds.NewSet[state]()
		queue := ds.Queue[state]{}

		for _, endState := range endStates {
			if visited.Add(endState) {
				queue = queue.Push(endState)
				bestPathsPoints.Add(endState.pos)
			}
		}

		var current state
		for !queue.IsEmpty() {
			current, queue = queue.Pop()

			tile := maze.Get(current.pos.XY())
			cardinalIndex := current.dir.GetCardinalIndex()

			for _, parent := range tile.parents[cardinalIndex] {
				if visited.Add(parent) {
					queue = queue.Push(parent)
					bestPathsPoints.Add(parent.pos)
				}
			}
		}
	}

	return bestScore, len(bestPathsPoints)
}
