#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'plusMinus' function below.
#
# The function accepts INTEGER_ARRAY arr as parameter.
#

def plusMinus(arr):
    n = len(arr)
    p, m, z = 0, 0, 0
    for x in arr:
        if x > 0:
            p += 1
        elif x < 0:
            m += 1
        else:
            z += 1
    print(f"{p/n:.6f}\n{m/n:.6f}\n{z/n:.6f}")
        

if __name__ == '__main__':
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)
