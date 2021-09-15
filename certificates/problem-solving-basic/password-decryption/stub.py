#!/bin/python3

import math
import os
import random
import re
import sys



#
# Complete the 'decryptPassword' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#

def decryptPassword(s):
    # Write your code here
    pass

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = decryptPassword(s)

    fptr.write(result + '\n')

    fptr.close()
