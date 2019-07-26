#coding=utf-8

import random
import csv
from functools import reduce
import os
import sys

# 无权无向图
def unweight_graph(nums, path):
    nodes = [i for i in range(0, int(nums))]
    l = [[i] for i in nodes]
    for item in l:
        max_num = 5 if len(nodes) > 5 else len(nodes) - 1
        node = random.sample(nodes, random.randint(0, max_num))
        node_clean = [i for i in node if i > item[0]] 
        for j in node_clean:
            if item[0] not in l[j]:
                l[j].append(item[0]) 
        item += node_clean
    if os.path.exists(path):
        path += '_1' 

    with open(path, 'w') as f:
            writer =  csv.writer(f)
            for i in l:
                i = [str(j) for j in i]
                writer.writerow([' '.join(i)])

    return path

if __name__ == '__main__':
    try:
        print(unweight_graph(sys.argv[1], sys.argv[2]))
    except Exception as e:
        print(e)