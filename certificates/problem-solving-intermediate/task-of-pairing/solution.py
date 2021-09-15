#!/bin/python3

import math
import os
import random
import re
import sys



#
# Complete the 'taskOfPairing' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts LONG_INTEGER_ARRAY freq as parameter.
#

def taskOfPairing(freq):
# Could break for n > sys.maxsize, but that's a lot.
    n = len(freq)
    inf = sys.maxsize
    if n == 1:
        return freq[0] // 2
    ans = 0
    # Normalize.
    for i in range(n):
        x = 0 if freq[i] <= 2 else ((freq[i] + 1) // 2) - 1
        ans += x
        freq[i] -= 2 * x
    # (#movable, #immovable)
    # dp := [(0, 0), (1, 0), (2, 0), (0, 1), (1, 1), (2, 1)]
    dp = [-inf] * 6
    dp[freq[0]] = 0
    if freq[0] == 2:
        dp[0] = 1
    for i in range(1, n):
        dpp = [-inf] * 6
        f = freq[i]
        # Case: Do nothing.
        dpp[f] = max(dp)
        # Case: Just get one from prev.
        dpp[f + 3] = max(dp[1], dp[2], dp[4], dp[5])
        if f >= 1:
            # Case: Prev has at least one movable or immovable. Combine once.
            dpp[f - 1] = max(dp[1:]) + 1
            # Case: Prev has movable and immovable. Combine with immovable and get movable.
            dpp[f - 1 + 3] = max(dpp[f - 1 + 3], max(dp[4], dp[5]) + 1)
        if f >= 2:
            # Case: Combine own two.
            dpp[f - 2] = max(dp) + 1
            # Case: Combine own two and try to get movable from prev.
            dpp[f - 2 + 3] = max(dp[1], dp[2], dp[4], dp[5]) + 1
        dp = dpp
    return ans + max(dp)
            
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    freq_count = int(input().strip())

    freq = []

    for _ in range(freq_count):
        freq_item = int(input().strip())
        freq.append(freq_item)

    result = taskOfPairing(freq)

    fptr.write(str(result) + '\n')

    fptr.close()
