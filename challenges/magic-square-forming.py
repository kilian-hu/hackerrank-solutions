#!/bin/python3

import math
import os
import random
import re
import sys

def is_magic(s):
    tmp = [
        s[0] + s[1] + s[2],
        s[3] + s[4] + s[5],
        s[6] + s[7] + s[8],
        s[0] + s[3] + s[6],
        s[1] + s[4] + s[7],
        s[2] + s[5] + s[8],
        s[0] + s[4] + s[8],
        s[2] + s[4] + s[6],
    ]
    return all(x == tmp[0] for x in tmp)

def cost(s1, s2):
    return sum(abs(a - b) for a, b in zip(s1, s2))

#
# Complete the 'formingMagicSquare' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY s as parameter.
#

def formingMagicSquare(s):
    s = s[0] + s[1] + s[2]
    ss = [None for _ in range(9)]
    a = [True for _ in range(10)]
    
    def rec(i, best):
        if i == 9 and is_magic(ss):
            return cost(s, ss)
        else:
            for j in range(1, 10):
                if a[j]:
                    a[j] = False
                    ss[i] = j
                    best = min(best, rec(i + 1, best))
                    a[j] = True
            return best
    
    return rec(0, sys.maxsize)
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = []

    for _ in range(3):
        s.append(list(map(int, input().rstrip().split())))

    result = formingMagicSquare(s)

    fptr.write(str(result) + '\n')

    fptr.close()
