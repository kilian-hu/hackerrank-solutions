#!/bin/python3

from collections import defaultdict

import math
import os
import random
import re
import sys

#
# Complete the 'organizingContainers' function below.
#
# The function is expected to return a STRING.
# The function accepts 2D_INTEGER_ARRAY container as parameter.
#

def organizingContainers(container):
    n = len(container)
    ball_cnt = [0 for _ in range(n)]
    cc = defaultdict(int)
    for c in container:
        cc[sum(c)] += 1
        for i, bc in enumerate(c):
            ball_cnt[i] += bc
    for bc in ball_cnt:
        if bc in cc and cc[bc] > 0:
            cc[bc] -= 1
        else:
            return "Impossible"
    return "Possible"

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input().strip())

    for q_itr in range(q):
        n = int(input().strip())

        container = []

        for _ in range(n):
            container.append(list(map(int, input().rstrip().split())))

        result = organizingContainers(container)

        fptr.write(result + '\n')

    fptr.close()
