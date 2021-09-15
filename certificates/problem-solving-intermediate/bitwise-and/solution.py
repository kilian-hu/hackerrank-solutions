#!/bin/python3

import math
import os
import random
import re
import sys

from collections import defaultdict

#
# Complete the 'countPairs' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts INTEGER_ARRAY arr as parameter.
#

def countPairs(arr):
    po2 = lambda x: x > 0 and not (x & (x - 1))
    d = defaultdict(int)
    for x in arr:
        d[x] += 1
    d = list(d.items())
    ans = 0
    for i in range(len(d)):
        a, a_cnt = d[i]
        for j in range(i, len(d)):
            b, b_cnt = d[j]
            if po2(a & b):
                if a == b:
                    ans += (a_cnt * (a_cnt - 1)) // 2
                else:
                    ans += a_cnt * b_cnt
    return ans          

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    arr_count = int(input().strip())

    arr = []

    for _ in range(arr_count):
        arr_item = int(input().strip())
        arr.append(arr_item)

    result = countPairs(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
