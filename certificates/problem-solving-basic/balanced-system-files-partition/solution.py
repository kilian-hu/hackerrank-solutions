#!/bin/python3

import math
import os
import random
import re
import sys


#
# Complete the 'mostBalancedPartition' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER_ARRAY parent
#  2. INTEGER_ARRAY files_size
#

def mostBalancedPartition(parent, files_size):
    n = len(parent)
    children = [[] for _ in range(n)]
    for i in range(1, n):
        children[parent[i]].append(i)
    size_sums = [None for _ in range(n)]
    
    def size_sums_rec(i):
        size_sums[i] = files_size[i] + sum(size_sums_rec(c) for c in children[i])
        return size_sums[i]
        
    size_sums_rec(0)
    return min(abs(size_sums[0] - 2 * ss) for ss in size_sums[1:])

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    parent_count = int(input().strip())

    parent = []

    for _ in range(parent_count):
        parent_item = int(input().strip())
        parent.append(parent_item)

    files_size_count = int(input().strip())

    files_size = []

    for _ in range(files_size_count):
        files_size_item = int(input().strip())
        files_size.append(files_size_item)

    result = mostBalancedPartition(parent, files_size)

    fptr.write(str(result) + '\n')

    fptr.close()
