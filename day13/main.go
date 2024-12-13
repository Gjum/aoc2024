package main

import (
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type xy struct {
	x, y int
}

type Game struct {
	a, b, p xy
}

func readFile(inpath string) ([]Game, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return nil, err
	}
	texts := strings.Split(strings.Trim(string(bytes[:]), "\n"), "\n\n")
	var input []Game
	re := regexp.MustCompile(`[0-9]+`)
	for _, text := range texts {
		matches := re.FindAllString(text, -1)
		ax, _ := strconv.Atoi(matches[0])
		ay, _ := strconv.Atoi(matches[1])
		bx, _ := strconv.Atoi(matches[2])
		by, _ := strconv.Atoi(matches[3])
		px, _ := strconv.Atoi(matches[4])
		py, _ := strconv.Atoi(matches[5])
		input = append(input, Game{xy{ax, ay}, xy{bx, by}, xy{px, py}})
	}
	return input, nil
}

// px = a * ax + b * bx
// py = a * ay + b * by

// a * ax = px - b * bx
// a = (px - b * bx) / ax
// (px - b * bx) / ax = (py - b * by) / ay
// px*ay - b *bx*ay = py*ax - b *by*ax
// b *by*ax - b *bx*ay = py*ax - px*ay
// b * (by*ax - bx*ay) = py*ax - px*ay
// b = (py*ax - px*ay) / (by*ax - bx*ay)

func part1(input []Game) int {
	output := 0
	for _, game := range input {
		ax := game.a.x
		ay := game.a.y
		bx := game.b.x
		by := game.b.y
		px := game.p.x
		py := game.p.y
		b := (py*ax - px*ay) / (by*ax - bx*ay)
		a := (px - b*bx) / ax
		if a*ax+b*bx == px && a*ay+b*by == py {
			output += 3*a + b
		}
	}
	return output
}

func part2(input []Game, off *big.Int) int64 {
	output := int64(0)
	for _, game := range input {
		ax := big.NewInt(int64(game.a.x))
		ay := big.NewInt(int64(game.a.y))
		bx := big.NewInt(int64(game.b.x))
		by := big.NewInt(int64(game.b.y))
		px := big.NewInt(int64(game.p.x))
		px.Add(px, off)
		py := big.NewInt(int64(game.p.y))
		py.Add(py, off)
		b := big.NewInt(0).Div(big.NewInt(0).Sub(big.NewInt(0).Mul(py, ax), big.NewInt(0).Mul(px, ay)),
			big.NewInt(0).Sub(big.NewInt(0).Mul(by, ax), big.NewInt(0).Mul(bx, ay)))
		a := big.NewInt(0).Div(big.NewInt(0).Sub(px, big.NewInt(0).Mul(b, bx)), ax)
		x := big.NewInt(0).Add(big.NewInt(0).Mul(a, ax), big.NewInt(0).Mul(b, bx))
		y := big.NewInt(0).Add(big.NewInt(0).Mul(a, ay), big.NewInt(0).Mul(b, by))
		if px.Cmp(x) == 0 && py.Cmp(y) == 0 {
			output += 3*a.Int64() + b.Int64()
		}
	}
	return output
}

func expect(expected, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("got %d but expected %d\n", actual, expected))
	}
	fmt.Printf("OK %d\n", actual)
}

func main() {
	inEx, err := readFile("day13/example.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	expect(480, part1(inEx))
	expect(480, int(part2(inEx, big.NewInt(0))))

	inCh, err := readFile("day13/challenge.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p1:", part1(inCh))
	fmt.Println("p1 big:", part2(inCh, big.NewInt(0)))
	fmt.Println("p2:", part2(inCh, big.NewInt(10000000000000)))
}
