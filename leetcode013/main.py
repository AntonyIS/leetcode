from typing import List


def countBits ( n:int) -> List[int] :
    dp = [0] * (n+1)
    offset = 1

    for i in range(1, n+ 1):
        n= 2
        if offset * 2 == i  :
            offset = i
        dp[i] = 1 + dp[i-offset]

    return dp

if __name__ == "__main__":
    res = countBits(10)
    print(res)

