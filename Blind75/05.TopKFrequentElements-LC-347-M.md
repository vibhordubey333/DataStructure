# Top K Frequent Elements — LeetCode 347 (M)

## 1. What is the problem?

**LeetCode 347 — Top K Frequent Elements**

Given an integer array nums and an integer k, return the k most frequent elements. Answer may be in any order. Better than O(n log n) is required.

**Example 1**

- **Input:** nums = [1,1,1,2,2,3], k = 2
- **Output:** [1,2]

**Example 2**

- **Input:** nums = [1], k = 1
- **Output:** [1]

---

## 2. Brute Force Solution

Count frequency of each number (map). Sort (number, count) by count descending, take first k numbers. O(n log n). Correct.

### Python (brute — sort)

```python
def top_k_frequent_brute(nums: list[int], k: int) -> list[int]:
    from collections import Counter
    count = Counter(nums)
    pairs = [(c, x) for x, c in count.items()]
    pairs.sort(key=lambda p: -p[0])
    return [x for _, x in pairs[:k]]
```

### Golang (brute)

```go
func topKFrequentBrute(nums []int, k int) []int {
	count := make(map[int]int)
	for _, x := range nums {
		count[x]++
	}
	type pair struct{ num, c int }
	var pairs []pair
	for num, c := range count {
		pairs = append(pairs, pair{num, c})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].c > pairs[j].c })
	out := make([]int, 0, k)
	for i := 0; i < k && i < len(pairs); i++ {
		out = append(out, pairs[i].num)
	}
	return out
}
```

- **Time:** O(n log n)  
- **Space:** O(n)

---

## 3. Optimized Solution (Bucket Sort / Min-Heap)

**Idea (Bucket):** Max frequency is at most n. Create buckets: bucket[i] = list of numbers with frequency i. Iterate from highest frequency down, collect numbers until we have k. O(n).

**Idea (Min-Heap):** Keep a min-heap of size k by frequency. For each (num, count), if heap size < k push (count, num); else if count > min of heap, pop min and push (count, num). Top k in heap at end. O(n log k).

**Steps (Bucket):**

1. Count frequencies.  
2. buckets = [[] for _ in range(len(nums)+1)]; for num, c in count.items(): buckets[c].append(num).  
3. out = []; for i from n down to 1: out.extend(buckets[i]); if len(out) >= k: return out[:k].  
4. Return out[:k].

### Python (bucket — O(n))

```python
def top_k_frequent(nums: list[int], k: int) -> list[int]:
    from collections import Counter
    count = Counter(nums)
    n = len(nums)
    buckets = [[] for _ in range(n + 1)]
    for num, c in count.items():
        buckets[c].append(num)
    out = []
    for i in range(n, 0, -1):
        out.extend(buckets[i])
        if len(out) >= k:
            return out[:k]
    return out[:k]
```

### Golang (bucket)

```go
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, x := range nums {
		count[x]++
	}
	buckets := make([][]int, len(nums)+1)
	for num, c := range count {
		buckets[c] = append(buckets[c], num)
	}
	var out []int
	for i := len(nums); i >= 1 && len(out) < k; i-- {
		for _, num := range buckets[i] {
			out = append(out, num)
			if len(out) == k {
				return out
			}
		}
	}
	return out
}
```

- **Time:** O(n)  
- **Space:** O(n)

---

## 4. Walkthrough

nums = [1,1,1,2,2,3], k = 2. count = {1:3, 2:2, 3:1}. buckets[1]=[3], buckets[2]=[2], buckets[3]=[1]. From i=3: add 1. From i=2: add 2. out = [1,2]. Return [1,2].

---

## 5. Edge Cases

1. **k = len(unique):** Return all.  
2. **k = 1:** Return single most frequent.  
3. **Ties:** Any order; bucket order is fine.  
4. **Single element:** Return that element.

---

## 6. Pitfalls

1. **Bucket index:** Frequency can be 1 to n; buckets[0] unused.  
2. **Order:** Problem allows any order; bucket order is valid.  
3. **Better than O(n log n):** Bucket gives O(n); heap gives O(n log k).

---

## 7. Follow-ups

1. **Kth largest element in stream?** Min-heap of size k.  
2. **Top k frequent words (with tie-break by lex)?** Heap or sort with custom comparator.  
3. **Top k in stream (data stream)?** Approximate: Count-Min Sketch or heap.

---

## 8. When to Use

- **Top k by frequency:** Bucket sort (freq 1..n) or min-heap of size k.  
- **Pattern:** Frequency count + bucket/heap; same in “sort characters by frequency”, “top k words”.
