#!/bin/python3

import math
import os
import random
import re
import sys


#
# Complete the 'maxSubarrayValue' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts INTEGER_ARRAY arr as parameter.
#

def maxSubarrayValue(arr):
    max_p, max_n, sum_p, sum_n,size =0,0,0,0,len(arr)
    for i in range(size):
        val = arr[i] if i % 2 == 0 else -arr[i]
        # for postive
        if val < sum_p+val:
            sum_p += val
        else:
            sum_p=val
        if sum_p > max_p:
            max_p = sum_p
        # for negative
        if val > sum_n+val:
            sum_n += val
        else:
            sum_n=val
        if sum_n < max_n:
            max_n = sum_n
        
    return max_p**2 if max_p**2 > max_n**2 else max_n**2

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    arr_count = int(input().strip())

    arr = []

    for _ in range(arr_count):
        arr_item = int(input().strip())
        arr.append(arr_item)

    result = maxSubarrayValue(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
