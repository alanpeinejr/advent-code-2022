package main

import ("testing")

func TestFindMaxCalorie(t *testing.T) {
	data := "1000"
	want := 1000
	calories := findMaxCalories(data);
	if want != calories {
		t.Fatalf(`Input %s only should return "%v" but instead returned "%v"`, data, want, calories);
    }
}

func TestFindMaxCaloriesOfTwoItems(t *testing.T) {
	data := 
	`1000
	 1000`
	want := 2000
	calories := findMaxCalories(data);
	if want != calories {
		t.Fatalf(`Input %s only should return "%v" but instead returned "%v"`, data, want, calories);
    }
}

func TestFindMaxCaloriesOfTwoElves(t *testing.T) {
	data := 
	`1000

	 2000`
	want := 3000
	calories := findMaxCalories(data);
	if want != calories {
		t.Fatalf(`Input %s only should return "%v" but instead returned "%v"`, data, want, calories);
    }
}

func TestFindMaxCaloriesOfTwoElvesTwoItems(t *testing.T) {
	data := 
	`1000
	 16000

	 2000`
	want := 19000
	calories := findMaxCalories(data);
	if want != calories {
		t.Fatalf(`Input %s only should return "%v" but instead returned "%v"`, data, want, calories);
    }
}

func TestFindMaxCaloriesTop3(t *testing.T) {
	data := 
	`1000
	 2000

	 1000

	 1000

	 2000

	 2000

	 4000

	 3500`
	want := 10500
	calories := findMaxCalories(data);
	if want != calories {
		t.Fatalf(`Input %s only should return "%v" but instead returned "%v"`, data, want, calories);
    }
}