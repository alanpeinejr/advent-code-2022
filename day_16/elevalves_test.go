package main

import (
	"testing"
)

func TestFindImpossilbeSpots(t *testing.T){
	input := "Valve AA has flow rate=0; tunnels lead to valves DD, II, BB\nValve BB has flow rate=13; tunnels lead to valves CC, AA\nValve CC has flow rate=2; tunnels lead to valves DD, BB\nValve DD has flow rate=20; tunnels lead to valves CC, AA, EE\nValve EE has flow rate=3; tunnels lead to valves FF, DD\nValve FF has flow rate=0; tunnels lead to valves EE, GG\nValve GG has flow rate=0; tunnels lead to valves FF, HH\nValve HH has flow rate=22; tunnel leads to valve GG\nValve II has flow rate=0; tunnels lead to valves AA, JJ\nValve JJ has flow rate=21; tunnel leads to valve II"
	rooms := parseLines(input)
	paths:= getAllShortestPaths(rooms)
	//because we'll never open the 0 rooms, we dont have to consider them
	roomsRemaining := make([]string, 0)
	for _, room := range rooms {
		if room.FlowRate != 0 {
			roomsRemaining = append(roomsRemaining, room.Name)
		}
	}
	actual := goToNextBest(rooms, paths, "AA", 30, 0, 0, 0, roomsRemaining)
	want:=1651
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted pressure to be %v`, input, actual, want )
	}
}
func TestFindFirstChoice(t *testing.T){
	input := "Valve AA has flow rate=0; tunnels lead to valves DD, II, BB\nValve BB has flow rate=13; tunnels lead to valves CC, AA\nValve CC has flow rate=2; tunnels lead to valves DD, BB\nValve DD has flow rate=20; tunnels lead to valves CC, AA, EE\nValve EE has flow rate=3; tunnels lead to valves FF, DD\nValve FF has flow rate=0; tunnels lead to valves EE, GG\nValve GG has flow rate=0; tunnels lead to valves FF, HH\nValve HH has flow rate=22; tunnel leads to valve GG\nValve II has flow rate=0; tunnels lead to valves AA, JJ\nValve JJ has flow rate=21; tunnel leads to valve II"
	rooms := parseLines(input)
	paths:= getAllShortestPaths(rooms)
	//because we'll never open the 0 rooms, we dont have to consider them
	roomsRemaining := make([]string, 0)
	for _, room := range rooms {
		if room.FlowRate != 0 {
			roomsRemaining = append(roomsRemaining, room.Name)
		}
	}
	actual := goToNextBest(rooms, paths, "AA", 3, 0, 0, 0, roomsRemaining)
	want:=20
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted pressure to be %v`, input, actual, want )
	}
}