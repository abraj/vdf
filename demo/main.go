// VDF
//
// Copyright 2019 by KeyFuse
//
// GPLv3 License

package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/abraj/vdf/sqrt"
)

var (
	flagX string
	flagT string

	// p256 = "60464814417085833675395020742168312237934553084050601624605007846337253615407"

	// p256: 2**256 - 189 (p ≡ 3 mod 4)
	p256 = "115792089237316195423570985008687907853269984665640564039457584007913129639747"
)

func init() {
	flag.StringVar(&flagX, "x", "", "x value")
	flag.StringVar(&flagT, "t", "1000", "t value")
}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " -x=[x value] -t=[y value]")
	fmt.Println("For example: " + os.Args[0] + " -x=7 -t=1000")
}

func sqrtVDF(t int64, x *big.Int, p *big.Int) {
	vdf := sqrt.NewVDFSqrt(p)
	// Delay.
	start := time.Now()
	y := vdf.Delay(t, x)
	cur := time.Now()
	delay := cur.Sub(start).Seconds()
	fmt.Println("-->")
	fmt.Println("Type:\tSloth Quadratic Residue")
	fmt.Printf("Value:\t%v\n", y.String())
	fmt.Printf("Delay:\t%.5f secs\n", delay)

	// Verify.
	start = time.Now()
	v := vdf.Verify(t, x, y)
	cur = time.Now()
	verify := cur.Sub(start).Seconds()
	fmt.Printf("Verify:\t%.5f secs\n", verify)
	fmt.Printf("Verify Result:\t%v\n", v)
	fmt.Printf("Delay/Verify:\t%.5f\n", delay/verify)
}

func main() {
	flag.Parse()
	if flagX == "" {
		usage()
		os.Exit(0)
	}
	t, _ := strconv.ParseInt(flagT, 10, 64)
	x, _ := new(big.Int).SetString(flagX, 10)
	p, _ := new(big.Int).SetString(p256, 10)
	fmt.Printf("T:%v, x:%v, P:%v\n", t, x.Text(10), p.Text(10))

	sqrtVDF(t, x, p)
}
