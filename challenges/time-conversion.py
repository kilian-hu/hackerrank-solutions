#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'timeConversion' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#

def timeConversion(s):
    is_twelve = (s[:2] == "12")
    is_pm = (s[-2:] == "PM")
    if is_pm:
        if is_twelve:
            return s[:-2]
        else:
            return str(int(s[:2]) + 12) + s[2:-2]
    else:
        if is_twelve:
            return "00" + s[2:-2]
        else:
            return s[:-2]
            
            
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = timeConversion(s)

    fptr.write(result + '\n')

    fptr.close()
