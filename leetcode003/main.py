from typing import Dict,List




def containsDuplicates(numbers:List[int]) -> bool :
    d:Dict = {} 
    for i in numbers:
        if i in d:
            return True
        else:
            d[i] = i
    return False

if __name__ == "__main__":
    prices:List[int] = [1,2,3,1]
    results = containsDuplicates(prices)
    print(results)
