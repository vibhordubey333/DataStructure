from typing import List
from numpy import np
from array import *


# Problem Link: https://www.codingninjas.com/codestudio/problems/number-of-subsets_3952532

# TODO: Fix Error while building.

def findWays(num: List[int], tar: int) -> int:
    n = len(num)
    rows, cols = (n + 1, tar + 1)
    cache = [[0] * cols] * rows

    # Initializing cache with -1
    for i in cache:
        for j in range(len(cache[0])):
            cache[i][j] = -1

    return countSubSets(num, tar, n, cache)
    pass


def countSubSets(num, tar, n, cache) -> int:
    if tar == 0:
        return 1

    if n == 0:
        return 0

    if num[n - 1] > tar:  # Ignoring element if it's gt target sum
        cache[n][tar] = countSubSets(num, tar, n - 1, )
        return cache[n][tar]

    excluding = countSubSets(num, tar, n - 1, cache)  # Ignoring element
    including = countSubSets(num, tar - num[n - 1], n - 1, cache)

    cache[n][tar] = excluding + including
    return cache[n][tar]


def main():
    nums = [1, 2, 2, 3]  # Output: CountSubSetSum_Recursive:  3
    tar = 3
    print("CountSubSetSum_Recursive: ", findWays(nums, tar))


if __name__ == '__main__':
    main()
