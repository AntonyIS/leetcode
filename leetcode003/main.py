from typing import Dict,List




def containsDuplicates(numbers:List[int]) :
    visitedValues:Dict = {} # Store visited values here

    # Loop through values in numbers
    for i in numbers:
        if i in visitedValues:
            return True
        else:
            visitedValues[i] = 0

    return False

if __name__ == "__main__":
    prices:List[int] = [1,2,3,1]
    results = containsDuplicates(prices)
    print(results)
