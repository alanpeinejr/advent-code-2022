package main

import ("testing")

func TestFindDuplicateType(t *testing.T) {
	var input = "vJrwpWtwJgWrhcsFMMfFFhFp"
	want := "p"
	duplicate := findDuplicateType(input);
	if want != duplicate {
		t.Fatalf(`Input "%v" should return duplicate type %q but instead returned %q`, input, want, duplicate);
    }
}

func TestPriorityScore(t *testing.T) {
	var input = "A"
	want := 27
	score := getPriority()(input);
	if want != score {
		t.Fatalf(`Input "%v" should return priority score %v but instead returned %v`, input, want, score);
    }
}

func TestRummage(t *testing.T) {
	var input = "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"
	want := 157
	score := rummageRucksack(input)
	if want != score {
		t.Fatalf(`Input "%v" should return priority score %v but instead returned %v`, input, want, score);
    }
}

func TestFindDuplicateAmongstThree(t *testing.T) {
	var input = []string{"vJrwpWtwJgWrhcsFMMfFFhFp","jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL","PmmdzqPrVvPwwTWBwg"}
	want := "r"
	badge := findBadgeType(input);
	if want != badge {
		t.Fatalf(`Input "%v" should return badge %q but instead returned %q`, input, want, badge);
    }
}

func TestFindBadgeScore(t *testing.T) {
	var input = "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"
	want := 70
	badgeScore := findBadges(input);
	if want != badgeScore {
		t.Fatalf(`Input "%v" should return badge %v but instead returned %v`, input, want, badgeScore);
    }
}