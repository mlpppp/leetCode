#!/bin/python3

import math
import os
import random
import re
import sys

# When someone calls 911, it is crucial to give them an ETA of when an officer will arrive.

# In order for our dispatchers not to have to guess, we want to write a function that given a map containing one officer and one incident location will give us the distance they have to traverse to get there, thus allowing us to give a good estimate.

# A map might look something like this:

# O _ _
# _ X T
# _ _ _

# T denotes the target incident

# O denotes the officer

# _ denotes an available route for the officer

# X denotes a square the officer can not move through

# -- In this case the officer will have to move two squares to the right and then one square down to get to the incident, making the total distance "3"



# A colleague has already started implementing a solution to this problem.

# They defined several test cases that we want to check against.

# They also implemented a base function that will read a map row by row and send it as a string (semicolon-separated) to another function.

# E.g. the map above would look like this:

# O__;_XT;___

# The output of this function should be the number of squares an officer has to move through to get to the incident (including the incident, so e.g. "3" in the example above)

# Can you complete their work and implement the missing function so that all test cases pass?

# Make sure to check all test cases to fully understand the requirements your solution should be able to cover.


#
# Complete the 'calculateDistance' function below.
#
# The function is expected to return an INTEGER.
# The function accepts STRING cityMap as parameter.
#

# -1 cannot reach
# -2 no O/T
# None no input

## ! input might not be squre

import math

def rearrangeCitymap(cityMap):
    rows = cityMap.split(";")
    max_len = 0
    for string in rows:
        if len(string) > max_len:
            max_len = len(string)
    for idx,string in enumerate(rows):
        lenDiff = max_len-len(string)
        rows[idx] = string + lenDiff*"X"

    return rows
    
        

dir_mat = [(1,0), (0,-1), (-1,0), (0,1)] 

def mapToAdj(cityMap):
    cityMap = rearrangeCitymap(cityMap)
    print(cityMap)
    # number every char as node
    numToChar_map = {}
    rows = len(cityMap) # 3
    cols = len(cityMap[0])  # 4
    elems = rows*cols
    for i in range(elems):
        numToChar_map[i] = cityMap[math.floor(i/cols)][i%cols]

    # use dir_mat to create adjacent table
    OandT = {}
    adjTable = {i:[] for i in range(len(numToChar_map))}
    for i in range(rows):
        for j in range(cols):
            curNode = i*cols + j
            if numToChar_map[curNode] == "X":
                continue

            if numToChar_map[curNode] == "O":
                OandT['O'] = curNode
            if numToChar_map[curNode] == "T":
                OandT['T'] = curNode
            for dir in dir_mat:
                # print(curNode)
                newJ = j+dir[1]
                newI = i+dir[0]
                if newJ >= cols or newI >= rows or newJ < 0 or newI < 0:
                    continue
                nextNode = (newI)*cols + newJ
                if (numToChar_map[nextNode] != "X"):
                    adjTable[curNode].append(nextNode)
    return adjTable, OandT
  
                    



       
            
                
            
    
def GraphBFS(adjTable, OandT):
    queue = []
    visited = set()
    queue.append(OandT['O'])
    visited.add(OandT['O'])
    step = 0
    while(len(queue)>0):
        queueSize = len(queue)
        for i in range(queueSize):
            curNode = queue.pop(0)
            if curNode == OandT['T']:
                return step
            for nextNode in adjTable[curNode]:
                if nextNode not in visited:
                    queue.append(nextNode)
                    visited.add(nextNode)
        step += 1           
    return -1
def calculateDistance(cityMap):
    if (not cityMap):
        return None
    adjTable, OandT = mapToAdj(cityMap)
    # print(adjTable)
    # print(OandT)
    if (('O' not in OandT) or ('T' not in OandT)):
        return -2

    steps = GraphBFS(adjTable, OandT)

    return steps

if __name__ == '__main__':
    cityMap  = "OX_;_X_T;___"
    steps = calculateDistance(cityMap)
    print (steps)

