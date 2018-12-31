from functools import lru_cache
from collections import namedtuple
from enum import IntEnum

depth = 4845
target_x = 6
target_y = 770

# depth = 510
# target_x = 10
# target_y = 10


def erosion_level(x, y):
    return (geologic_index(x, y) + depth) % 20183


@lru_cache(maxsize=None)
def geologic_index(x, y):
    if x == 0 and y == 0:
        return 0
    if x == target_x and y == target_y:
        return 0

    if y == 0:
        return x * 16807
    if x == 0:
        return y * 48271
    else:
        return erosion_level(x - 1, y) * erosion_level(x, y - 1)


def region_type(x, y):
    return erosion_level(x, y) % 3


class Tool(IntEnum):
    CLIMBING_GEAR = 0
    TORCH = 1
    NEITHER = 2


class RegionType(IntEnum):
    ROCKY = 0
    WET = 1
    NARROW = 2


Region = namedtuple("Region", ["x", "y", "tool"])

s = 0
for x in range(target_x + 1):
    for y in range(target_y + 1):
        s += region_type(x, y)

print(s)

from dijkstra import Graph, dijkstra, shortest

allowed_tools = {
    RegionType.ROCKY: set((Tool.CLIMBING_GEAR, Tool.TORCH)),
    RegionType.WET: set((Tool.CLIMBING_GEAR, Tool.NEITHER)),
    RegionType.NARROW: set((Tool.TORCH, Tool.NEITHER)),
}

padding = 50
graph = Graph()
for x in range(target_x + padding + 1):
    for y in range(target_y + padding + 1):
        rt = RegionType(region_type(x, y))
        for tool in allowed_tools[rt]:
            graph.add_vertex(Region(x, y, tool))

for x in range(target_x + padding):
    for y in range(target_y + padding):
        rt = region_type(x, y)
        tools = allowed_tools[rt]
        t1, t2 = tools
        graph.add_edge(Region(x, y, t1), Region(x, y, t2), 7)
        for n in ((x + 1, y), (x, y + 1)):
            rt_n = region_type(*n)
            if rt == rt_n:
                graph.add_edge(Region(x, y, t1), Region(*n, t1), 1)
                graph.add_edge(Region(x, y, t2), Region(*n, t2), 1)
            else:
                overlap = tools.intersection(allowed_tools[rt_n]).pop()
                graph.add_edge(Region(x, y, overlap), Region(*n, overlap), 1)


start = Region(0, 0, Tool.TORCH)
end = Region(target_x, target_y, Tool.TORCH)

dijkstra(graph, graph.get_vertex(start), graph.get_vertex(end))
target = graph.get_vertex(end)
path = [target.get_id()]
shortest(target, path)
# print("The shortest path : %s" % (path[::-1]))

s = 0
for v, w in zip(path[:-1], path[1:]):
    if v.tool == w.tool:
        s += 1
    else:
        s += 7
print(s)
