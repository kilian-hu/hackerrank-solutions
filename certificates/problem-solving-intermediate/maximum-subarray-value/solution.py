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
    n = len(arr)
    
    if n == 1:
        return arr[0]**2

    # Initialize running totals for even and odd sums
    even_sum = 0
    odd_sum = 0
    
    # Initialize max and min difference seen so far
    max_diff = float('-inf')
    min_diff = float('inf')
    
    # Track the maximum possible value
    max_value = float('-inf')
    
    for i in range(n):
        if i % 2 == 0:
            even_sum += arr[i]
        else:
            odd_sum += arr[i]
        
        current_diff = even_sum - odd_sum
        
        # Update the max value considering current_diff
        max_value = max(max_value, current_diff**2)
        
        # Update max_diff and min_diff
        max_diff = max(max_diff, current_diff)
        min_diff = min(min_diff, current_diff)
    
    # Calculate the maximum possible value
    max_value = max(max_value, (max_diff - min_diff)**2)
    return max_value
    
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
