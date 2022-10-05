from typing import List

# Problem Link: https://www.codingninjas.com/codestudio/problems/number-of-subsets_3952532
# Solution works but TLE Encountered

def findWays(num: List[int], tar: int) -> int:
    count = 0
    return countSubSets(num, tar, len(num), count)
    pass


def countSubSets(num, tar, n, count) -> int:
    if tar == 0:
        count += 1
        return count

    if n == 0:
        return count

    if num[n - 1] > tar:  # Ignoring element if it's gt target sum
        return countSubSets(num, tar, n - 1, count)

    count = countSubSets(num, tar, n - 1, count)  # Ignoring element
    count = countSubSets(num, tar - num[n - 1], n - 1, count)

    return count


def main():
    nums = [1, 2, 2, 3] # Output: CountSubSetSum_Recursive:  3
    tar = 3
    print("CountSubSetSum_Recursive: ", findWays(nums, tar))


if __name__ == '__main__':
    main()
