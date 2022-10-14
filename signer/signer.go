package main

import (
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// SingleHash считает значение crc32(data)+"~"+crc32(md5(data)) ( конкатенация двух строк через ~),
// где data - то что пришло на вход (по сути - числа из первой функции)
func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	muMD := &sync.Mutex{}

	for data := range in {
		input, ok := data.(int)
		if !ok {
			logrus.Error("SingleHash: Convert failed: interface{] -> int")
		}

		value := strconv.Itoa(input)

		wg.Add(1)
		go func(data string) {
			defer wg.Done()
			hashLeft := make(chan string, 1)
			hashRight := make(chan string, 1)

			go func() {
				muMD.Lock()
				md5hash := DataSignerMd5(data)
				muMD.Unlock()

				hashRight <- DataSignerCrc32(md5hash)
			}()

			go func() {
				hashLeft <- DataSignerCrc32(data)
			}()

			out <- <-hashLeft + "~" + <-hashRight
		}(value)
	}

	wg.Wait()
}

// MultiHash считает значение crc32(th+data)) (конкатенация цифры, приведённой к строке и строки),
// где th=0..5 ( т.е. 6 хешей на каждое входящее значение ), потом берёт конкатенацию результатов в
// порядке расчета (0..5), где data - то что пришло на вход (и ушло на выход из SingleHash)
func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for data := range in {
		value, ok := data.(string)
		if !ok {
			logrus.Error("MultiHash: Convert failed: interface{] -> string")
		}

		wg.Add(1)
		go func(data string) {
			defer wg.Done()
			workers := &sync.WaitGroup{}

			dataHashes := make([]string, 6)

			for i := 0; i < 6; i++ {
				workers.Add(1)

				go func(i int) {
					defer workers.Done()

					dataHashes[i] = DataSignerCrc32(strconv.Itoa(i) + data)
				}(i)
			}
			workers.Wait()
			out <- strings.Join(dataHashes, "")

		}(value)
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var res []string
	for data := range in {
		res = append(res, data.(string))
	}
	sort.Strings(res)

	out <- strings.Join(res, "_")
}

func jobWorker(wg *sync.WaitGroup, j job, in, out chan interface{}) {
	defer wg.Done()
	j(in, out)
	close(out)
}

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}

	in := make(chan interface{})
	out := make(chan interface{})

	for _, j := range jobs {
		wg.Add(1)
		go jobWorker(wg, j, in, out)
		in = out
		out = make(chan interface{})
	}

	wg.Wait()

	close(out)
}
