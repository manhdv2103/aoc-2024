package main

import (
	"container/heap"
	"maps"
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

type point = utils.Point

type state struct {
	pos point
	dir point
}
type pqItem = utils.PQItem[state]

func solveP1(input string) string {
	maze, start, end := parseInput(input)

	visited := make(utils.Set[point])
	pq := make(utils.PriorityQueue[state], 0)
	heap.Init(&pq)

	var endStateItem *pqItem = nil
	heap.Push(&pq, &pqItem{
		Value:    state{start, point{X: 1, Y: 0}},
		Priority: 0,
	})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItem)
		s := item.Value

		if visited[s.pos] {
			continue
		}
		visited[s.pos] = true

		if s.pos == end {
			endStateItem = item
			break
		}

		for _, direction := range utils.Directions {
			nextPos := utils.AddPoint(s.pos, direction)
			if !maze[nextPos] || visited[nextPos] {
				continue
			}

			heap.Push(&pq, &pqItem{
				Value: state{
					nextPos,
					direction,
				},
				Priority: item.Priority - (1 + 1000*utils.Ternary(
					s.dir == direction,
					0,
					utils.Ternary(s.dir == utils.InversePoint(direction), 2, 1),
				)),
			})
		}
	}

	return strconv.Itoa(-endStateItem.Priority)
}

type stateP2 struct {
	pos     point
	dir     point
	visited utils.Set[point]
}
type pqItemP2 = utils.PQItem[stateP2]

type firstVisitItemKey struct {
	pos point
	dir point
}

func solveP2(input string) string {
	maze, start, end := parseInput(input)

	globalVisited := make(map[firstVisitItemKey]*pqItemP2)
	pq := make(utils.PriorityQueue[stateP2], 0)
	heap.Init(&pq)

	var endStateItem *pqItemP2 = nil
	heap.Push(&pq, &pqItemP2{
		Value:    stateP2{start, point{X: 1, Y: 0}, make(utils.Set[point])},
		Priority: 0,
	})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItemP2)
		s := item.Value

		if s.visited[s.pos] {
			continue
		}
		s.visited[s.pos] = true

		if s.pos == end {
			endStateItem = item
			break
		}

		for _, direction := range utils.Directions {
			nextPos := utils.AddPoint(s.pos, direction)
			if !maze[nextPos] || s.visited[nextPos] {
				continue
			}

			nextPqItem := pqItemP2{
				Value: stateP2{
					nextPos,
					direction,
					maps.Clone(s.visited),
				},
				Priority: item.Priority - (1 + 1000*utils.Ternary(
					s.dir == direction,
					0,
					utils.Ternary(s.dir == utils.InversePoint(direction), 2, 1),
				)),
			}

			fviKey := firstVisitItemKey{s.pos, direction}
			fvi, isVisited := globalVisited[fviKey]
			if isVisited {
				if nextPqItem.Priority == fvi.Priority {
					for k, v := range s.visited {
						fvi.Value.visited[k] = v
					}
				}
				continue
			}
			globalVisited[fviKey] = &nextPqItem

			heap.Push(&pq, &nextPqItem)
		}
	}

	return strconv.Itoa(len(endStateItem.Value.visited))
}

func parseInput(input string) (maze utils.Set[point], start point, end point) {
	maze = make(utils.Set[point])
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, tile := range line {
			switch tile {
			case '.':
				maze[point{X: x, Y: y}] = true
			case 'S':
				start = point{X: x, Y: y}
				maze[point{X: x, Y: y}] = true
			case 'E':
				end = point{X: x, Y: y}
				maze[point{X: x, Y: y}] = true
			}
		}
	}

	return maze, start, end
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
