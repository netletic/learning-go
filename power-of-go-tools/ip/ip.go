package ip

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sort"
	"strings"
)

func ExtractIP(line string) (string, error) {
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return "", fmt.Errorf("failed to parse line %s", line)
	}
	ipAddr := net.ParseIP(fields[0])
	if ipAddr == nil {
		return "", fmt.Errorf("failed to extract IP address from line %s", line)
	}
	return ipAddr.String(), nil
}

func Main() map[string]int {
	f, err := os.Open("testdata/clf.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	frequency := map[string]int{}
	for input.Scan() {
		line := input.Text()
		ip, err := ExtractIP(line)
		if err != nil {
			continue
		}
		frequency[ip]++
	}

	type freq struct {
		addr  string
		count int
	}
	freqs := make([]freq, 0, len(frequency))
	for addr, count := range frequency {
		freqs = append(freqs, freq{addr: addr, count: count})
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})

	fmt.Println(freqs)

	fmt.Printf("%-16s%s\n", "Address", "Requests")
	for i, freq := range freqs {
		if i > 9 {
			break
		}
		fmt.Printf("%-16s%d\n", freq.addr, freq.count)
	}

	return frequency
}
