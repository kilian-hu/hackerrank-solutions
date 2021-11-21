#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'extraLongFactorials' function below.
#
# The function accepts INTEGER n as parameter.
#

def extraLongFactorials(n):
    f = lambda x: x * f(x - 1) if x > 1 else 1
    print(f(n))

if __name__ == '__main__':
    n = int(input().strip())

    extraLongFactorials(n)
