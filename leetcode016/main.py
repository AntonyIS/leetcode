from typing import List
def missingNumber(nums: List[int]) -> int:
    res = len(nums)
    
    for i in range(len(nums)):
       
        res += (i - nums[i])
    return res


res = missingNumber([9,6,4,2,3,5,7,0,1])
print(res)