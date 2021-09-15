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
    # Gets timeouts.
    even = [0]
    odd = [0]
    for i in range(len(arr)):
        if i % 2 == 0:
            even.append(even[-1] + arr[i])
            odd.append(odd[-1])
        else:
            even.append(even[-1])
            odd.append(odd[-1] + arr[i])
    ans = 0
    for i in range(len(arr)):
        for j in range(i + 1, len(arr) + 1):
            a = even[j] - even[i]
            b = odd[j] - odd[i]
            ans = max(ans, (a - b)**2)
    return ans
    
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
