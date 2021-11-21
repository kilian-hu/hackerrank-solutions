#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'dayOfProgrammer' function below.
#
# The function is expected to return a STRING.
# The function accepts INTEGER year as parameter.
#

def dayOfProgrammer(year):
    if year <= 1917:
        return f"{12 if year % 4 == 0 else 13}.09.{year}"
    elif year == 1918:
        return "26.09.1918"
    else:
        d = 12 if (year % 4 == 0 and year % 100 != 0) or year % 400 == 0 else 13
        return f"{d}.09.{year}"

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    year = int(input().strip())

    result = dayOfProgrammer(year)

    fptr.write(result + '\n')

    fptr.close()
