#!/bin/python3

import math
import os
import random
import re
import sys



#
# Complete the 'maxCost' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER_ARRAY cost
#  2. STRING_ARRAY labels
#  3. INTEGER dailyCount
#

def maxCost(cost, labels, dailyCount):
    # Write your code here
    pass

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    cost_count = int(input().strip())

    cost = []

    for _ in range(cost_count):
        cost_item = int(input().strip())
        cost.append(cost_item)

    labels_count = int(input().strip())

    labels = []

    for _ in range(labels_count):
        labels_item = input()
        labels.append(labels_item)

    dailyCount = int(input().strip())

    result = maxCost(cost, labels, dailyCount)

    fptr.write(str(result) + '\n')

    fptr.close()
