#!/bin/python3

import math
import os
import random
import re
import string
import sys



#
# Complete the 'authEvents' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts 2D_STRING_ARRAY events as parameter.
#

P = 131
M = 10**9 + 7
PP = [P**i for i in range(11)]
APPENDS = [""] + list(string.ascii_letters) + [str(d) for d in range(10)]

def calc_hash(pw):
    cur_h = 0
    for i in range(len(pw)):
        cur_h += ord(pw[-(i+1)]) * PP[i]
    return cur_h % M

def authEvents(events):
    cur_h = None
    good_hashs = None
    ans = []
    for event, value in events:
        if event == "setPassword":
            good_hashs = set(calc_hash(value + x) for x in APPENDS)
        else:
            assert event == "authorize"
            ans.append(1 if int(value) in good_hashs else 0)
    return ans

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    events_rows = int(input().strip())
    events_columns = int(input().strip())

    events = []

    for _ in range(events_rows):
        events.append(input().rstrip().split())

    result = authEvents(events)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
