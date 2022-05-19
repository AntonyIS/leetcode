from typing import List


def twoSum_1(nums: List[int], target:int) -> List[int]:
    # Store values visited in a dictionary : key --> values, value --> index
    dict = {} # stores the values that have been visited
    # Loop all the items in the list
    for i in range(len(nums)):
        diff = target - nums[i]
        if diff in dict:
            return [dict[diff], i]
        else:
            dict[nums[i]] = i

    return []


def twoSum_2(nums: List[int], target:int) -> List[int]:
    dict = {}
    for i, v in enumerate(nums):
        if (target - v) in dict:
            return [dict[target - v], i]
        else:
            dict[v] = i
    return []
if __name__ == "__main__":
    numbers:List[int] = [2,7,11,15]
    target:int = 9 
    results = twoSum_2(numbers, target)
    print(results)