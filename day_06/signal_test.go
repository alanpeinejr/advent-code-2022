package main

import ("testing")

func TestHasDuplicates(t *testing.T){
	var input string = "abca"
	want := true
	actual := hasDuplicates(input)
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}
func TestNoHasDuplicates(t *testing.T){
	var input string = "abcd"
	want := false
	actual := hasDuplicates(input)
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}

func TestFirstUnique4(t *testing.T) {
	var input string = "abcd" 
	want := 4
	index := findPacketIdentifierEndIndex(input)
	if want != index {
		t.Fatalf(`Input "%v" should return %v but instead returned %v`, input, want, index);
    }
}
func TestLaterUnique4(t *testing.T) {
	var input string = "aabcdeeeeeeee" 
	want := 5
	index := findPacketIdentifierEndIndex(input)
	if want != index {
		t.Fatalf(`Input "%v" should return %v but instead returned %v`, input, want, index);
    }
}
func TestLastUnique4(t *testing.T) {
	var input string = "aabcd" 
	want := 5
	index := findPacketIdentifierEndIndex(input)
	if want != index {
		t.Fatalf(`Input "%v" should return %v but instead returned %v`, input, want, index);
    }
}
func TestNoUnique4(t *testing.T) {
	var input string = "abccccccccccccc" 
	want := -1
	index := findPacketIdentifierEndIndex(input)
	if want != index {
		t.Fatalf(`Input "%v" should return %v but instead returned %v`, input, want, index);
    }
}
func TestGivenSampleUnique4(t *testing.T) {
	var input string = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw" 
	want := 11
	index := findPacketIdentifierEndIndex(input)
	if want != index {
		t.Fatalf(`Input "%v" should return %v but instead returned %v`, input, want, index);
    }
}
func TestGivenSampleUnique14(t *testing.T) {
	var input string = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg" 
	want := 29
	index := findMessageIdentifierEndIndex(input)
	if want != index {
		t.Fatalf(`Input "%v" should return %v but instead returned %v`, input, want, index);
    }
}