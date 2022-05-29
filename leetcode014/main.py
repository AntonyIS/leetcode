from typing import List


def missingNumber1(nums:List[int]) -> int:
    n = len(nums) + 1
    sum_all = int((n/2) * (n-1))
    return sum_all - sum(nums)

def missingNumber2(nums:List[int]) -> int:
    XOR_1 = 0
    XOR_2 = 0
    for i in nums:
        XOR_1 ^= i
    
    for i in range(len(nums)+ 1):
        XOR_2 ^= i
    
    return XOR_1^XOR_2

if __name__ == "__main__":
    nums = [3,0,1]
    res = missingNumber2(nums)
    print(res)
