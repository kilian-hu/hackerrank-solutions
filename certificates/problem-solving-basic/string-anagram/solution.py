#!/bin/python3

import math
import os
import random
import re
import sys

from collections import defaultdict

#
# Complete the 'stringAnagram' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts following parameters:
#  1. STRING_ARRAY dictionary
#  2. STRING_ARRAY query
#

def stringAnagram(dictionary, query):
    d = defaultdict(int)
    for w in dictionary:
        d["".join(sorted(w))] += 1
    ans = []
    for w in query:
        w = "".join(sorted(w))
        ans.append(d[w] if w in d else 0)
    return ans
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    dictionary_count = int(input().strip())

    dictionary = []

    for _ in range(dictionary_count):
        dictionary_item = input()
        dictionary.append(dictionary_item)

    query_count = int(input().strip())

    query = []

    for _ in range(query_count):
        query_item = input()
        query.append(query_item)

    result = stringAnagram(dictionary, query)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
