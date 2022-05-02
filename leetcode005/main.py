def getSum(a:int, b:int) -> int :
    # Bit shifting manupilation
    XOR = a ^ b # Returns 1 if one of the bits is 1 and the other is 0 else returns false
    AND = a & b # Returns 1 if all bits are 1 else 0

    if AND == 0:
        # Nothing to carry | shift left
        return XOR
    else:
        
        return getSum(XOR, AND << 1)

if __name__ == "__main__":
    a:int = 2
    b:int = 10
    sum = getSum(a, b)
    print(sum)