#!/bin/python3

import math
import os
import random
import re
import sys


#
# Complete the 'sortedSum' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY a as parameter.
#


def sortedSum(a):
    s = 0
    for i in range(len(a)):
        x = list(reversed(sorted(a[:(i+1)],reverse=True)))
        g = 0
        for i in range(len(x)):
            g = g+(x[i]*(i+1))
        s=s+g
    return s

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    a_count = int(input().strip())

    a = []

    for _ in range(a_count):
        a_item = int(input().strip())
        a.append(a_item)

    result = sortedSum(a)

    fptr.write(str(result) + '\n')

    fptr.close()
