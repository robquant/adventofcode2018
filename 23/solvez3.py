import re
import sys
import z3

intmat = []
for line in open(sys.argv[1]):
    x, y, z, r = tuple(map(int, re.findall("-?\d+", line)))
    intmat.append([x, y, z, r])
bots = []

for l in intmat:
    x, y, z, r = l
    bots.append((r, x, y, z))


best = max(bots)
br, bx, by, bz = best
tot = 0

for b in bots:
    r, x, y, z = b
    if abs(x - bx) + abs(y - by) + abs(z - bz) <= br:
        tot += 1


def dist1d(a, b):
    d = a - b
    return z3.If(d >= 0, d, -d)


def manhattan(ax, ay, az, bx, by, bz):
    return dist1d(ax, bx) + dist1d(ay, by) + dist1d(az, bz)


solver = z3.Optimize()

bestx = z3.Int("bestx")
besty = z3.Int("besty")
bestz = z3.Int("bestz")
distance = z3.Int("distance")

inside = []
for i, b in enumerate(bots):
    br, *bxyz = b
    bot = z3.Int("b{:4d}".format(i))
    ok = z3.If(manhattan(bestx, besty, bestz, *bxyz) <= br, 1, 0)
    solver.add(bot == ok)
    inside.append(bot)

solver.add(distance == manhattan(bestx, besty, bestz, 0, 0, 0))

solver.maximize(z3.Sum(*inside))
solver.minimize(distance)
solver.check()

m = solver.model()
min_distance = m.eval(distance)

print(min_distance)
