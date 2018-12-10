from operator import itemgetter
from collections import defaultdict

test = [(1, 1), (1, 6), (8, 3), (3, 4), (5, 5), (8, 9)]


def dist(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])


def part1(coords):
    max_x = max(coords, key=itemgetter(0))[0]
    max_y = max(coords, key=itemgetter(1))[1]
    areas = defaultdict(int)
    infinite = set()
    for x in range(max_x + 1):
        for y in range(max_y + 1):
            dists = sorted(
                ((dist((x, y), point), point) for point in coords), key=itemgetter(0)
            )
            if dists[0][0] < dists[1][0]:
                areas[dists[0][1]] += 1
                if x == 0 or y == 0 or x == max_x or y == max_y:
                    infinite.add(dists[0][1])

    return sorted(
        ((point, area) for point, area in areas.items() if point not in infinite),
        key=itemgetter(1),
    )[-1][1]


def part2(coords, maxdistsum):
    max_x = max(coords, key=itemgetter(0))[0]
    max_y = max(coords, key=itemgetter(1))[1]
    npoints = 0
    for x in range(max_x + 1):
        for y in range(max_y + 1):
            distsum = sum(dist((x, y), point) for point in coords)
            if distsum < maxdistsum:
                npoints += 1
    return npoints


input = [tuple(map(int, i.split(", "))) for i in open("input.txt").readlines()]

print(part1(input))
print(part2(input, 10000))
