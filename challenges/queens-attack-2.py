#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'queensAttack' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER n
#  2. INTEGER k
#  3. INTEGER r_q
#  4. INTEGER c_q
#  5. 2D_INTEGER_ARRAY obstacles
#

def queensAttack(n, k, r_q, c_q, obstacles):
    # Keep track of closest obstacles.
    row_left = 0
    row_right = n + 1
    col_up = n + 1
    col_down = 0
    f = lambda x: n - x + 1
    d1_up = ((r_q + c_q + 2 * (n - max(r_q, c_q))) // 2) + 1
    d1_down = (r_q + c_q - 2 * min(r_q, c_q)) // 2
    d1_q = (r_q + c_q) // 2
    d2_up = (f(r_q) + c_q - 2 * min(f(r_q), c_q)) // 2
    d2_down = ((f(r_q) + c_q + 2 * (n - max(f(r_q), c_q))) // 2) + 1
    d2_q = (f(r_q) + c_q) // 2
    d_q = r_q - c_q
    s_q = r_q + c_q
    for r_o, c_o in obstacles:
        if r_o == r_q:
            if c_o > c_q:
                row_right = min(row_right, c_o)
            else:
                row_left = max(row_left, c_o)
        elif c_o == c_q:
            if r_o > r_q:
                col_up = min(col_up, r_o)
            else:
                col_down = max(col_down, r_o)
        elif r_o - c_o == d_q:
            d1_o = (r_o + c_o) // 2
            if r_o > r_q:
                d1_up = min(d1_up, d1_o)
            else:
                d1_down = max(d1_down, d1_o)
        elif r_o + c_o == s_q:
            d2_o = (f(r_o) + c_o) // 2
            if c_o > c_q:
                d2_down = min(d2_down, d2_o)
            else:
                d2_up = max(d2_up, d2_o)
    return (row_right - row_left + col_up - col_down +
            d1_up - d1_down + d2_down - d2_up - 8)
            
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    k = int(first_multiple_input[1])

    second_multiple_input = input().rstrip().split()

    r_q = int(second_multiple_input[0])

    c_q = int(second_multiple_input[1])

    obstacles = []

    for _ in range(k):
        obstacles.append(list(map(int, input().rstrip().split())))

    result = queensAttack(n, k, r_q, c_q, obstacles)

    fptr.write(str(result) + '\n')

    fptr.close()
