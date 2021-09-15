#!/bin/python3

import math
import os
import random
import re
import sys



#
# Complete the 'findSubstring' function below.
#
# The function is expected to return a STRING.
# The function accepts following parameters:
#  1. STRING s
#  2. INTEGER k
#

def findSubstring(s, k):
    vowels = ["a", "e", "i", "o", "u"]
    cur = best = sum([c in vowels for c in s[:k]])
    ans = 0
    for i in range(k, len(s)):
        cur += s[i] in vowels
        cur -= s[i - k] in vowels
        if cur > best:
            best = cur
            ans = i - k + 1
    if best > 0:
        return s[ans:(ans+k)]
    else:
        return "Not found!"

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    k = int(input().strip())

    result = findSubstring(s, k)

    fptr.write(result + '\n')

    fptr.close()
