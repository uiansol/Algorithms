#!/bin/python3

import math
import os
import random
import re
import sys


def climbingLeaderboard(ranked, player):
    result = []
    leaderboard = sorted(set(ranked), reverse = True)
    
    rank = len(leaderboard)
    for play in player:
        while rank > 0 and play >= leaderboard[rank - 1]:
            rank -= 1
        result.append(rank + 1)
            
    return result

    
                
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    ranked_count = int(input().strip())

    ranked = list(map(int, input().rstrip().split()))

    player_count = int(input().strip())

    player = list(map(int, input().rstrip().split()))

    result = climbingLeaderboard(ranked, player)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
