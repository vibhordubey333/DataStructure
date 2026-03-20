# Coin Change — LeetCode 322 (M)

## 1. What is the problem?

**LeetCode 322 — Coin Change**

You are given an integer array `coins` and an integer `amount`. Return the fewest number of coins needed to make `amount` (or -1 if impossible). Each coin may be used unlimited times.

**Example 1**

- **Input:** coins = [1, 2, 5], amount = 11
- **Output:** 3 (5+5+1)

**Example 2**

- **Input:** coins = [2], amount = 3
- **Output:** -1

---

## 2. Brute Force Solution

Recurse: try each coin, subtract from amount, recurse; take min of 1 + recurse(amount - c). Base: amount 0 → 0; amount < 0 → infinity. Correct but exponential.

### Python (brute force)

```python
def coin_change_brute(coins: list[int], amount: int) -> int:
    def f(a):
        if a == 0:
            return 0
        if a < 0:
            return float('inf')
        best = float('inf')
        for c in coins:
            best = min(best, 1 + f(a - c))
        return best
    out = f(amount)
    return -1 if out == float('inf') else out
```

### Golang (brute force)

```go
func coinChangeBrute(coins []int, amount int) int {
	var f func(int) int
	f = func(a int) int {
		if a == 0 {
			return 0
		}
		if a < 0 {
			return 1 << 31
		}
		best := 1 << 31
		for _, c := range coins {
			next := 1 + f(a-c)
			if next < best {
				best = next
			}
		}
		return best
	}
	out := f(amount)
	if out >= 1<<31 {
		return -1
	}
	return out
}
```

- **Time:** O(amount * |coins|) with memo; without memo exponential.  
- **Space:** O(amount) with memo

---

## 3. Optimized Solution (DP)

**Idea:** dp[a] = minimum number of coins to make amount a. dp[0]=0; dp[a] = 1 + min over coins c of dp[a-c] (if a>=c). Fill from 1 to amount.

**Steps:**

1. dp = [inf] * (amount+1); dp[0] = 0.  
2. For a from 1 to amount: for each coin c <= a, dp[a] = min(dp[a], 1+dp[a-c]).  
3. Return dp[amount] if finite else -1.

### Python (optimized)

```python
def coin_change(coins: list[int], amount: int) -> int:
    dp = [float('inf')] * (amount + 1)
    dp[0] = 0
    for a in range(1, amount + 1):
        for c in coins:
            if c <= a:
                dp[a] = min(dp[a], 1 + dp[a - c])
    return dp[amount] if dp[amount] != float('inf') else -1
```

### Golang (optimized)

```go
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if c <= a && 1+dp[a-c] < dp[a] {
				dp[a] = 1 + dp[a-c]
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
```

- **Time:** O(amount * len(coins))  
- **Space:** O(amount)

---

## 4. Walkthrough

coins = [1,2,5], amount = 11. dp[0]=0; dp[1]=1; dp[2]=1; … dp[5]=1; … dp[11]=3. Return 3.

---

## 5. Edge Cases

1. **amount = 0:** return 0.  
2. **Impossible:** e.g. coins=[2], amount=3 → -1.  
3. **One coin type:** amount divisible by coin.  
4. **Large amount:** DP is still feasible.

---

## 6. Pitfalls

1. **Initializing dp[0]=0:** Essential.  
2. **Unbounded coins:** This is unbounded knapsack; same coin can be used many times.  
3. **Return -1:** When no combination reaches amount.

---

## 7. Follow-ups

1. **Number of combinations (not min coins)?** DP where dp[a] += dp[a-c].  
2. **Print one solution?** Store parent pointer.  
3. **Limited supply per coin?** 0/1 or bounded knapsack.

---

## 8. When to Use

- **Unbounded knapsack / minimum coins:** 1D DP over amount.  
- **Pattern:** Classic DP; similar to “minimum number of squares” etc.
