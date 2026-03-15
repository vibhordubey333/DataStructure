# Best Time to Buy and Sell Stock — LeetCode 121 (E)

## 1. What is the problem?

**LeetCode 121 — Best Time to Buy and Sell Stock**

You are given an array `prices` where `prices[i]` is the price of a stock on day `i`. You may complete at most one buy and one sell. Find the maximum profit you can achieve (0 if you cannot profit).

**Example 1**

- **Input:** `prices = [7, 1, 5, 3, 6, 4]`
- **Output:** `5`
- **Explanation:** Buy on day 2 (1), sell on day 5 (6); profit = 6 − 1 = 5.

**Example 2**

- **Input:** `prices = [7, 6, 4, 3, 1]`
- **Output:** `0`
- **Explanation:** No transaction yields profit.

---

## 2. Brute Force Solution

Try every pair of days `(i, j)` with `i < j` and take the maximum of `prices[j] - prices[i]`. Correct because we consider every possible buy/sell pair.

### Python (brute force)

```python
def max_profit_brute(prices: list[int]) -> int:
    n = len(prices)
    best = 0
    for i in range(n):
        for j in range(i + 1, n):
            best = max(best, prices[j] - prices[i])
    return best
```

### Golang (brute force)

```go
func maxProfitBrute(prices []int) int {
	best := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			diff := prices[j] - prices[i]
			if diff > best {
				best = diff
			}
		}
	}
	return best
}
```

- **Time:** O(n²)  
- **Space:** O(1)

---

## 3. Optimized Solution

**Idea:** For each day, the best profit if we sell **today** is: today’s price minus the **minimum price seen so far**. So one pass: track the minimum price seen so far and the maximum profit (current price − min so far).

**Steps:**

1. Initialize `min_price = infinity`, `max_profit = 0`.
2. For each `price` in `prices`:
   - Update `max_profit = max(max_profit, price - min_price)`.
   - Update `min_price = min(min_price, price)`.
3. Return `max_profit`.

### Python (optimized)

```python
def max_profit(prices: list[int]) -> int:
    min_price = float('inf')
    max_profit = 0
    for p in prices:
        max_profit = max(max_profit, p - min_price)
        min_price = min(min_price, p)
    return max_profit
```

### Golang (optimized)

```go
func maxProfit(prices []int) int {
	minPrice := 1 << 31
	maxProfit := 0
	for _, p := range prices {
		if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
		if p < minPrice {
			minPrice = p
		}
	}
	return maxProfit
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

**Example 1:** `prices = [7, 1, 5, 3, 6, 4]`

| Step | p | min_price | profit = p − min | max_profit |
|------|---|-----------|------------------|------------|
| 0    | 7 | 7         | 0                | 0          |
| 1    | 1 | 1         | 0                | 0          |
| 2    | 5 | 1         | 4                | 4          |
| 3    | 3 | 1         | 2                | 4          |
| 4    | 6 | 1         | 5                | 5          |
| 5    | 4 | 1         | 3                | 5          |

Output: `5`.

---

## 5. Edge Cases

1. **Single day:** `[1]` → 0.  
2. **Always decreasing:** `[5,4,3,2,1]` → 0.  
3. **Always increasing:** `[1,2,3,4,5]` → 4.  
4. **Two days:** `[1, 4]` → 3.  
5. **Min on last day:** e.g. `[2, 4, 1]` → 2 (buy 2, sell 4).

---

## 6. Pitfalls

1. **Selling before buying:** Always compute profit as sell price − buy price with buy index &lt; sell index.  
2. **Updating min after profit:** Update min **after** using it for profit so “buy” is in the past.  
3. **Negative profit:** Don’t return negative; max profit is at least 0.  
4. **Kadane-style:** This is “max difference with min so far,” not subarray sum.

---

## 7. Follow-ups

1. **Multiple transactions (buy/sell many times)?** LeetCode 122: sum every positive price increase (greedy).  
2. **At most 2 transactions?** LeetCode 123: DP with states (no stock, one buy, one sell, two buys, two sells).  
3. **At most k transactions?** DP: `dp[t][d]` = max profit with at most t transactions by day d.  
4. **Cooldown after sell?** LeetCode 309: state machine (hold, sold, cooldown).

---

## 8. When to Use

- **Single pass + track extremum:** “Best pair (buy, sell)” with buy before sell → track minimum so far and max profit.  
- **Pattern:** One-pass greedy; similar ideas in “max difference,” “maximum subarray” (Kadane).
