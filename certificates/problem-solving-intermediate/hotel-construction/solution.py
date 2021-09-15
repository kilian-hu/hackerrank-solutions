#!/bin/python3

import math
import os
import random
import re
import sys

from itertools import product

#
# Complete the 'numberOfWays' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY roads as parameter.
#

def numberOfWays(roads):
    n = len(roads) + 1
    adj = [[] for _ in range(n)]
    for i, j in roads:
        adj[i - 1].append(j - 1)
        adj[j - 1].append(i - 1)
    ans = 0
    
    def dfs(x, d):
        dist[x] = d
        for y in adj[x]:
            if dist[y] == -1:
                dfs(y, d + 1)
    
    # Brute force.
    for i in range(n - 2):
        for j in range(i + 1, n - 1):
            for k in range(j + 1, n):
                dist = [-1 for _ in range(n)]
                dfs(i, 0)
                if dist[j] != dist[k]:
                    continue
                dist = [-1 for _ in range(n)]
                dfs(j, 0)
                if dist[i] == dist[k]:
                    ans += 1
    return ans

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    roads_rows = int(input().strip())
    roads_columns = int(input().strip())

    roads = []

    for _ in range(roads_rows):
        roads.append(list(map(int, input().rstrip().split())))

    result = numberOfWays(roads)

    fptr.write(str(result) + '\n')

    fptr.close()
