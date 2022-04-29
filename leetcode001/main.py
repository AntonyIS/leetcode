from typing import List


def twoSum(numbers: List[int], target:int) -> List[int]:
    # Store values visted in a dictionary : key --> values, value --> index
    compliments = {} # stores the values that have been visited
    # Loop all the items in the list
    for i in range(len(numbers)):
        diff = target - numbers[i]
        if diff in compliments:
            return [compliments[diff], i]
        else:
            compliments[numbers[i]] = i


    return []


if __name__ == "__main__":
    numbers:List[int] = [2,7,11,15]
    target:int = 9 
    results = twoSum(numbers, target)
    print(results)