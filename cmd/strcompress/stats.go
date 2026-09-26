package main

import (
	"fmt"
	"sort"
)

// runStats prints a census of string literals under the given roots to
// help pick a compression threshold.
func runStats(roots []string, threshold int) error {
	dirs, err := discover(roots)
	if err != nil {
		return err
	}

	type pkgStats struct {
		name        string
		count       int
		bytes       int
		uniqueCount int
		uniqueBytes int
		totalBytes  int
	}

	var total pkgStats
	results := make([]pkgStats, 0, len(dirs))

	for _, dir := range dirs {
		files, err := parsePackage(dir)
		if err != nil {
			return err
		}
		stats := pkgStats{name: dir}
		values := make(map[string]int)
		for _, f := range files {
			c := &collector{threshold: threshold}
			c.collect(f.ast)
			for _, cand := range c.cands {
				stats.count++
				stats.bytes += len(cand.value)
				values[cand.value]++
			}
			stats.totalBytes += c.skippedSmall
		}
		stats.uniqueCount = len(values)
		for value := range values {
			stats.uniqueBytes += len(value)
		}
		results = append(results, stats)
		total.count += stats.count
		total.bytes += stats.bytes
		total.uniqueCount += stats.uniqueCount
		total.uniqueBytes += stats.uniqueBytes
		total.totalBytes += stats.totalBytes
	}

	fmt.Printf("threshold %d bytes\n", threshold)
	fmt.Printf("packages: %d\n", len(results))
	fmt.Printf("literals >= threshold: %d (%d bytes, %d unique / %d bytes)\n",
		total.count, total.bytes, total.uniqueCount, total.uniqueBytes)
	fmt.Printf("bytes in sub-threshold literals: %d\n", total.totalBytes)

	sort.Slice(results, func(i, j int) bool { return results[i].bytes > results[j].bytes })
	if len(results) > 15 {
		results = results[:15]
	}
	fmt.Println("\ntop packages by compressible bytes:")
	for _, s := range results {
		fmt.Printf("%-60s %8d bytes, %5d literals (%d unique)\n", s.name, s.bytes, s.count, s.uniqueCount)
	}
	return nil
}
