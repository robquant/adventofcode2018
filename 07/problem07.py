import re
from collections import defaultdict
from functools import reduce
from operator import itemgetter

test = """Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
"""

pattern = re.compile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")
depends = defaultdict(set)

for line in open("input.txt").readlines():
    # for line in test.split("\n"):
    if line:
        dep = pattern.match(line).groups()
        depends[dep[1]].add(dep[0])

all_nodes = set(reduce(lambda s1, s2: s1.union(s2), depends.values()))
all_nodes = all_nodes.union(set(depends.keys()))
for node in all_nodes:
    if not node in depends:
        depends[node] = set()


def part1(all_nodes, depends):
    completed = set()
    order = ""

    while len(completed) < len(all_nodes):
        node = min(
            node
            for node, deps in depends.items()
            if deps.issubset(completed) and not node in completed
        )
        order += node
        completed.add(node)
    return order


def availabe_workers(workers):
    return (i for i, w in enumerate(workers) if idle(w))


def time_for_letter(letter):
    return 61 + ord(letter) - ord("A")


def next_finished(workers):
    next_dt = 9999
    index = 0
    for i, w in enumerate(workers):
        if idle(w):
            continue
        if w[1] < next_dt:
            index = i
            next_dt = w[1]
    return index


def idle(worker):
    return worker[1] == 0


IDLE = ("_", 0)


def part2(all_nodes, depends, nworkers=5):
    workers = [IDLE] * nworkers
    total_time = 0
    completed = set()
    in_progress = set()
    while len(completed) < len(all_nodes):
        available_letters = sorted(
            node
            for node, deps in depends.items()
            if deps.issubset(completed)
            and not node in completed
            and not node in in_progress
        )
        for worker_index, letter in zip(availabe_workers(workers), available_letters):
            workers[worker_index] = (letter, time_for_letter(letter))
        min_index = next_finished(workers)
        letter, dt = workers[min_index]
        total_time += dt
        completed.add(letter)
        # print("Finished:", workers[min_index])
        current_workers = []
        for w in workers:
            if idle(w):
                current_workers.append(w)
            else:
                current_workers.append((w[0], w[1] - dt))
        workers = current_workers
        # print(workers)
        in_progress = {l for l, _ in workers}
    return total_time


print(part1(all_nodes, depends))
print(part2(all_nodes, depends))
