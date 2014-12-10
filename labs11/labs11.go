package main

import "fmt"
import "encoding/json"

type M struct {
	AA *A
	BB []*A
	CC []*A
	DD []*A
	EE map[string]*A
}

type A struct {
	X    int
	Y    int
	Z    int
	Next *A
}

func main() {
	var m = &M{
		AA: &A{
			X: 1, Y: 2, Z: 3, Next: &A{
				X: 10, Y: 20, Z: 30, Next: &A{
					X: 100, Y: 200, Z: 300,
				},
			},
		},
		BB: []*A{
			&A{X: 1, Y: 2, Z: 3},
			&A{X: 10, Y: 20, Z: 30},
			&A{X: 100, Y: 200, Z: 300},
		},
		CC: []*A{
			&A{
				X: 1, Y: 2, Z: 3, Next: &A{
					X: 10, Y: 20, Z: 30, Next: &A{
						X: 100, Y: 200, Z: 300,
					},
				},
			},
			&A{
				X: 4, Y: 5, Z: 6, Next: &A{
					X: 40, Y: 50, Z: 60, Next: &A{
						X: 400, Y: 500, Z: 600,
					},
				},
			},
			&A{
				X: 7, Y: 8, Z: 9, Next: &A{
					X: 70, Y: 80, Z: 90, Next: &A{
						X: 700, Y: 800, Z: 900,
					},
				},
			},
		},
		DD: make([]*A, 3),
		EE: map[string]*A{
			"A": &A{X: 1, Y: 2, Z: 3},
			"B": &A{X: 10, Y: 20, Z: 30},
			"C": &A{X: 100, Y: 200, Z: 300},
		},
	}

	var j, err = json.Marshal(m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", j)

	var b = new(M)

	var err2 = json.Unmarshal(j, b)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("\n[M.AA]")

	for a := b.AA; a != nil; a = a.Next {
		fmt.Printf("%p -> %+v\n", a, *a)
	}

	fmt.Println("\n[M.BB]")

	for i, a := range b.BB {
		fmt.Printf("%d -> %+v\n", i, *a)
	}

	fmt.Println("\n[M.CC]")

	for i, a := range b.CC {
		fmt.Printf("%d ->\n", i)
		for aa := a; aa != nil; aa = aa.Next {
			fmt.Printf("  %p -> %+v\n", aa, *aa)
		}
	}

	fmt.Println("\n[M.DD]")

	for i, a := range b.DD {
		fmt.Printf("%d -> %v\n", i, a)
	}

	fmt.Println("\n[M.EE]")

	for i, a := range b.EE {
		fmt.Printf("%s -> %+v\n", i, *a)
	}
}
