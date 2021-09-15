#!/bin/python3

import math
import os
import random
import re
import sys


#
# Complete the 'longestSubarray' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY arr as parameter.
#

def longestSubarray(arr):
    n = len(arr)
    ans = 0
    # O(n^2) is okay because of constraints.
    for i in range(n):
        w = []
        cnt = 0
        for j in range(i, n):
            if arr[j] in w:
                cnt += 1
                continue
            if len(w) == 0:
                w.append(arr[j])
            elif len(w) == 1:
                if abs(w[0] - arr[j]) > 1:
                    break
                else:
                    w.append(arr[j])
            else:
                break
            cnt += 1
        ans = max(ans, cnt)
    return ans
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    arr_count = int(input().strip())

    arr = []

    for _ in range(arr_count):
        arr_item = int(input().strip())
        arr.append(arr_item)

    result = longestSubarray(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
