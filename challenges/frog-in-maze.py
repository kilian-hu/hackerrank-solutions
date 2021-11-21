#!/bin/python3

import math
import os
import random
import re
import sys



if __name__ == '__main__':
    first_multiple_input = input().rstrip().split()
    n = int(first_multiple_input[0])
    m = int(first_multiple_input[1])
    k = int(first_multiple_input[2])

    board = [list(input()) for _ in range(n)]
    t_exit = dict()
    
    for kk in range(k):
        second_multiple_input = input().rstrip().split()
        i1 = int(second_multiple_input[0]) - 1
        j1 = int(second_multiple_input[1]) - 1
        i2 = int(second_multiple_input[2]) - 1
        j2 = int(second_multiple_input[3]) - 1
        t_exit[(i1, j1)] = (i2, j2)
        t_exit[(i2, j2)] = (i1, j1)

    states = [(-1, -1)]  # Death state.
    state2idx = dict()
    v = [0.0]  # State values, i.e. probability.
    s_init = -1
    transitions = [[(0, 1.0)]]  # Self loop when dead.
    
    for i1 in range(n):
        for j1 in range(m):
            x = board[i1][j1]
            if x in ["A", "%", "O"]:
                state2idx[(i1, j1)] = len(states)
                states.append((i1, j1))
            elif x in ["#", "*"]:
                state2idx[(i1, j1)] = 0       
            else:
                assert False, x
    
    for i1 in range(n):
        for j1 in range(m):
            x = board[i1][j1]
            state_idx = state2idx[(i1, j1)]
            if x in ["A", "%", "O"]:
                if x == "A":
                    s_init = state_idx
                    v.append(0.0)
                elif x == "%":
                    v.append(1.0)
                    transitions.append([(state_idx, 1.0)])  # Self loop when exit.
                elif x == "O":
                    v.append(0.0)
                else:
                    assert False, x
                if x in ["A", "O"]:
                    i2, j2 = t_exit[(i1, j1)] if (i1, j1) in t_exit else (i1, j1)
                    succs = []
                    deaths = 0
                    for a, b in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
                        i3, j3 = i2 + a, j2 + b
                        if (0 <= i3 < n) and (0 <= j3 < m):
                            y = board[i3][j3]
                            if y in ["A", "%", "O"]:
                                succs.append((i3, j3))
                            elif y == "*":
                                deaths += 1
                    if len(succs) == 0:
                        transitions.append([(0, 1.0)])
                    else:
                        t = [(state2idx[s], 1 / (len(succs) + deaths)) for s in succs]
                        if deaths > 0:
                            t.append((0, deaths / (len(succs) + deaths)))
                        transitions.append(t)
          
    while True:
        v_old = v.copy()
        for state in range(len(states)):
            x = 0.0
            for succ, prob in transitions[state]:
                x += v[succ] * prob
            v[state] = x
        keep_going = False
        for state in range(len(states)):
            if abs(v[state] - v_old[state]) > 1e-10:
                keep_going = True
                break
        if not keep_going:
            break

    print(v[s_init])
