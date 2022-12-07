package main

import ("testing")

func TestBuilding(t *testing.T){
	var input string = "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k"
	want := 48381165
	actual := buildFileSystem(input).calculateSize()
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}
func TestSumming(t *testing.T){
	var input string = "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k"
	want := 95437
	files :=buildFileSystem(input);
	files.calculateSize()
	actual := findSizeOfDirectoriesLessThanTenHundredThousand(files)
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}

func TestFindingSmallestToDeleteForNeededSpace(t *testing.T) {
	var input string = "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k"
	want := "d"
	files :=buildFileSystem(input);
	files.calculateSize()
	actual := findSmallestDirForNeededSpace(70000000, 30000000, files).name
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}