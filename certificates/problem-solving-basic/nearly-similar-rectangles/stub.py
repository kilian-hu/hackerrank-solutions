#!/bin/python3

import math
import os
import random
import re
import sys


#
# Complete the 'nearlySimilarRectangles' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts 2D_LONG_INTEGER_ARRAY sides as parameter.
#

def nearlySimilarRectangles(sides):
    # Write your code here
    pass

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    sides_rows = int(input().strip())
    sides_columns = int(input().strip())

    sides = []

    for _ in range(sides_rows):
        sides.append(list(map(int, input().rstrip().split())))

    result = nearlySimilarRectangles(sides)

    fptr.write(str(result) + '\n')

    fptr.close()

