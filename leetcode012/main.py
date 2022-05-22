def hammingWeight( n: int) -> int:
    res = 0
    while n : # while n != 0
        if n % 2 :
            res += 1
        n = n >> 1
    return res


if __name__ == "__main__":
    num = 123
    res = hammingWeight(num)
    print(res)

