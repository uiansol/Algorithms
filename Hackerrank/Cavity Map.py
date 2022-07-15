#!/bin/python3

import math
import os
import random
import re
import sys


def cavityMap(grid):
    for i in range(len(grid)):
        if i == 0 or i == len(grid) - 1:
            continue
        for j in range(len(grid[i])):
            if j == 0 or j == len(grid[i]) - 1:
                continue
            if isCavity(grid, i, j):
                grid[i] = grid[i][:j] + 'X' + grid[i][j + 1:]

    return grid
                
def isCavity(grid, i, j):
    if grid[i - 1][j] < grid[i][j] and\
       grid[i + 1][j] < grid[i][j] and\
       grid[i][j - 1] < grid[i][j] and\
       grid[i][j + 1] < grid[i][j]:
        return True
    return False

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    grid = []

    for _ in range(n):
        grid_item = input()
        grid.append(grid_item)

    result = cavityMap(grid)

    fptr.write('\n'.join(result))
    fptr.write('\n')

    fptr.close()
