#!/bin/python3

import math
import os
import random
import re
import sys

def acmTeam(topic):
    pairs = {}
    for t1 in range(len(topic) - 1):
        for t2 in range(t1 + 1, len(topic)):
            nTopics = 0
            for i in range(len(topic[t1])):
                if topic[t1][i] == '1' or topic[t2][i] == '1':
                    nTopics += 1
    
            if nTopics in pairs.keys():
                pairs[nTopics] += 1
            else:
                pairs[nTopics] = 1
                
    maxTopics = max(pairs.keys())
    return [maxTopics, pairs[maxTopics]]
    
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
