package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"sync"
	"time"
)

type EnshuranceMap struct {
	sync.RWMutex
	EMap map[int]string
}

type PointerOrNo struct {
	Num int
}

func (p PointerOrNo) Value() {
	p.Num = 111
}

func (p *PointerOrNo) Pointer() {
	p.Num = 111
}

func main() {
	p := PointerOrNo{Num: 777}
	fmt.Println(p.Num)
	p.Value()
	fmt.Println(p.Num)
	p.Pointer()
	fmt.Println(p.Num)
	m5 := make(map[string]map[string]int)
	m5["count"] = map[string]int{"ru": 1}
	m5["count"]["ru"]++
	m5["count"]["ru"]++
	fmt.Println(m5["count"]["ru"])

	for j, v := range []string{"ru", "en", "jp"} {
		for i := 0; i < 5; i++ {
			m5["count"][v] += i * (j + 1)
		}
	}

	for k, v := range m5["count"] {
		fmt.Println("Country:", k, "Number:", v)
	}

	mSync := sync.Map{}
	mSync.Store(1, "_one")
	mSync.Store(2, "_two")
	mSync.Store(3, "_three")
	mSync.Store(4, "_four")

	go func() {
		for i := 1; i < 5; i++ {
			v, ok := mSync.Load(i)
			if !ok {
				log.Fatal("bad load")
			}
			fmt.Println("Key_1:", i, "Value_1:", v)
		}
	}()

	go func() {
		for i := 1; i < 5; i++ {
			v, ok := mSync.Load(i)
			if !ok {
				log.Fatal("no value")
			}
			mSync.Store(i, v.(string)+"_test")
			fmt.Println("Key:", i, "Value:", v)
		}
	}()
	m := make(map[int]string, 2)
	m2 := make(map[int][]string, 2)
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	m[4] = "four"

	// map is unsafe
	go func() {
		for k, v := range m {
			fmt.Println("Key:", k, "Value:", v)
		}
	}()

	go func() {
		for k, v := range m {
			m[k] = v + " :" + strconv.Itoa(k)
		}
	}()

	m2[0] = append(m2[0], "num1", "num2", "num3")
	m3 := EnshuranceMap{EMap: make(map[int]string)}
	m3.Lock()
	m3.EMap[1] = "One"
	m3.Unlock()

	m3.RLock()
	fmt.Println(m3.EMap)
	m3.RUnlock()

	fmt.Println(m)
	fmt.Println(m2)

	//  disordered
	fmt.Println("Unordered")
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}
	fmt.Println()

	// ordered
	var key []int
	for k := range m {
		key = append(key, k)
	}

	sort.Ints(key)
	fmt.Println("Ordered")
	for _, k := range key {
		fmt.Println("Key:", k, "Value:", m[k])
	}
	time.Sleep(time.Second * 1)
}
