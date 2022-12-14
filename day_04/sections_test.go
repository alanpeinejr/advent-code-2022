package main

import ("testing")

func TestContainsSection(t *testing.T) {
	var section1, section2 SectionRange = SectionRange{0,100}, SectionRange{1,99}
	want := true
	overlaps := section1.contains(section2)
	if want != overlaps {
		t.Fatalf(`Input "%v, %v" should return %v but instead returned %v`, section1, section2, want, overlaps);
    }
}

func TestDoesNotContainSection(t *testing.T) {
	var section1, section2 SectionRange = SectionRange{0,100}, SectionRange{1,99}
	want := false
	overlaps := section2.contains(section1)
	if want != overlaps {
		t.Fatalf(`Input "%v, %v" should return %v but instead returned %v`, section1, section2, want, overlaps);
    }
}

func TestOverlapsSectionUpper(t *testing.T) {
	var section1, section2 SectionRange = SectionRange{1,14}, SectionRange{10,15}
	want := true
	overlaps := section1.overlaps(section2)
	if want != overlaps {
		t.Fatalf(`Input "%v, %v" should return %v but instead returned %v`, section1, section2, want, overlaps);
    }
}

func TestOverlapsSectionLower(t *testing.T) {
	var section1, section2 SectionRange = SectionRange{13,16}, SectionRange{10,15}
	want := true
	overlaps := section1.overlaps(section2)
	if want != overlaps {
		t.Fatalf(`Input "%v, %v" should return %v but instead returned %v`, section1, section2, want, overlaps);
    }
}

func TestOverlapsSectionOuterUpper(t *testing.T) {
	var section1, section2 SectionRange = SectionRange{10,15}, SectionRange{14,16}
	want := true
	overlaps := section1.overlaps(section2)
	if want != overlaps {
		t.Fatalf(`Input "%v, %v" should return %v but instead returned %v`, section1, section2, want, overlaps);
    }
}
func TestOverlapsSectionFully(t *testing.T) {
	var section1, section2 SectionRange = SectionRange{10,15}, SectionRange{11,14}
	want := true
	overlaps := section1.overlaps(section2)
	if want != overlaps {
		t.Fatalf(`Input "%v, %v" should return %v but instead returned %v`, section1, section2, want, overlaps);
    }
}

func TestOverlapsSectionOuterLower(t *testing.T) {
	var section1, section2 SectionRange = SectionRange{10,15}, SectionRange{9,10}
	want := true
	overlaps := section1.overlaps(section2)
	if want != overlaps {
		t.Fatalf(`Input "%v, %v" should return %v but instead returned %v`, section1, section2, want, overlaps);
    }
}

func TestDoesNotOverlapSection(t *testing.T) {
	var section1, section2 SectionRange = SectionRange{0,50}, SectionRange{51,99}
	want := false
	overlaps := section2.overlaps(section1)
	if want != overlaps {
		t.Fatalf(`Input "%v, %v" should return %v but instead returned %v`, section1, section2, want, overlaps);
    }
}

func TestParsingSections(t *testing.T) {
	input := "2-50,4-51"
	want1 :=  SectionRange{2,50}
	want2 := SectionRange{4,51}
	var section1, section2 SectionRange = getSectionRanges(input)
	if (want1 != section1){
		t.Fatalf(`Input "%v" should return first section: %v but instead returned %v`, input, want1, section1)
	}
	if (want2 != section2){
		t.Fatalf(`Input "%v" should return first section: %v but instead returned %v`, input, want2, section2)
	}
}

func TestCountingBorderContainsPairsOfSections(t *testing.T){
	input := "2-50,3-51"
	want := 0
	count := countFullContainingSections(input)
	if (want != count){
		t.Fatalf(`Input "%v" should result in %v but is instead %v`, input, want, count)
	}
}

func TestCountingMultipleLinesContainsPairsOfSections(t *testing.T){
	input := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
	want := 2
	count := countFullContainingSections(input)
	if (want != count){
		t.Fatalf(`Input "%v" should result in %v but is instead %v`, input, want, count)
	}
}
func TestCountingMultipleLinesOverlappingPairsOfSections(t *testing.T){
	input := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
	want := 4
	count := countOverlappingSections(input)
	if (want != count){
		t.Fatalf(`Input "%v" should result in %v but is instead %v`, input, want, count)
	}
}