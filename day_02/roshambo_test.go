package main

import ("testing")

func TestMatchScoringTie(t *testing.T) {
	var opp, me = 1, 1
	want := 4
	scorer := matchWinScore();
	score := scorer(opp, me);
	if want != score {
		t.Fatalf(`Input %v vs %v should return score "%v" but instead returned "%v"`, opp, me, want, score);
    }
}
func TestMatchTieIfClosureWorks(t *testing.T) {
	want := 8
	scorer := matchWinScore();
	score := scorer(1, 1);
	score = scorer(1, 1);
	if want != score {
		t.Fatalf(`Input "1,1" vs "1,1"should return score "%v" but instead returned "%v"`, want, score);
    }
}
func TestMatchWinWithRock(t *testing.T) {
	want := 7
	scorer := matchWinScore();
	score := scorer(3, 1);
	if want != score {
		t.Fatalf(`Input "3" vs "1"should return score "%v" but instead returned "%v"`, want, score);
    }
}
func TestMatchWinWithPaper(t *testing.T) {
	want := 8
	scorer := matchWinScore();
	score := scorer(1, 2);
	if want != score {
		t.Fatalf(`Input "1" vs "2" should return score "%v" but instead returned "%v"`, want, score);
    }
}
func TestMatchWinWithScissors(t *testing.T) {
	want := 9
	scorer := matchWinScore();
	score := scorer(2, 3);
	if want != score {
		t.Fatalf(`Input "2" vs "3"should return score "%v" but instead returned "%v"`, want, score);
    }
}
func TestLosing(t *testing.T) {
	want := 6
	scorer := matchWinScore();
	score := scorer(1, 3);
	score = scorer(2, 1);
	score = scorer(3, 2);
	if want != score {
		t.Fatalf(`Input should return score "%v" but instead returned "%v"`, want, score);
    }
}

func TestSimulatingOneMatch(t *testing.T) {
	input := "A X" 
	want := 3
	output := simulateMatches(input)
	if want != output {
		t.Fatalf(`Input %s should return score "%v" but instead returned "%v"`, input, want, output);
    }
}
func TestSimulatingMoreMatch(t *testing.T) {
	input := "A Y\nB X\nC Z" 
	want := 12
	output := simulateMatches(input)
	if want != output {
		t.Fatalf(`Input %s should return score "%v" but instead returned "%v"`, input, want, output);
    }
}