#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'acmTeam' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts STRING_ARRAY topic as parameter.
#

def acmTeam(topic):
    most = 0
    most_cnt = 0
    for i in range(len(topic) - 1):
        for j in range(i + 1, len(topic)):
            x = sum(t1 == "1" or t2 == "1" for t1, t2 in zip(topic[i], topic[j]))
            if x > most:
                most = x
                most_cnt = 1
            elif x == most:
                most_cnt += 1
    return most, most_cnt

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    m = int(first_multiple_input[1])

    topic = []

    for _ in range(n):
        topic_item = input()
        topic.append(topic_item)

    result = acmTeam(topic)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
