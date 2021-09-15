#!/bin/python3

import math
import os
import random
import re
import sys



#
# Complete the 'possibleChanges' function below.
#
# The function is expected to return a STRING_ARRAY.
# The function accepts STRING_ARRAY usernames as parameter.
#

def possibleChanges(usernames):
    ans = []
    for u in usernames:
        if len(u) <= 1:
            ans.append("NO")
        for i in range(len(u) - 1):
            if u[i] > u[i + 1]:
                ans.append("YES")
                break
        else:
            ans.append("NO")
    return ans
        
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    usernames_count = int(input().strip())

    usernames = []

    for _ in range(usernames_count):
        usernames_item = input()
        usernames.append(usernames_item)

    result = possibleChanges(usernames)

    fptr.write('\n'.join(result))
    fptr.write('\n')

    fptr.close()
