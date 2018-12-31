from collections import defaultdict

def dist(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1]) + abs(p1[2] - p2[2]) + abs(p1[3] - p2[3])

def form_cluster(p, points, close):
    cluster = set([p])
    to_add = close[p]
    while to_add:
        c = to_add.pop()
        if c not in cluster:
            cluster.add(c)
            points.remove(c)
            to_add += close[c]
    return cluster

points = [tuple(map(int, line.split(','))) for line in open("input.txt")]

close = defaultdict(list)
for i, p1 in enumerate(points):
    for j in range(i+1, len(points)):
        d = dist(p1, points[j])
        if d <= 3:
            close[p1].append(points[j])
            close[points[j]].append(p1)

clusters = []

while points:
    p = points.pop()
    cluster = form_cluster(p, points, close)
    clusters.append(cluster)

print(len(clusters))
