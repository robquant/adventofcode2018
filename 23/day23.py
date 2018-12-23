import re
import sys
from collections import namedtuple

Particle = namedtuple("Particle", ["x", "y", "z", "r"])

def dist(p1, p2):
    return abs(p1.x - p2.x) + abs(p1.y - p2.y) + abs(p1.z - p2.z) 

particles = []
max_particle = None
for line in open(sys.argv[1]):
    x,y,z,r = tuple(map(int, re.findall("-?\d+", line)))
    particle = Particle(x, y, z, r)
    if max_particle is None or particle.r >= max_particle.r:
        max_particle = Particle(*particle)
    particles.append(particle)
print("Max particle: ", max_particle)

counter = 0
for particle in particles:
    if dist(particle, max_particle) <= max_particle.r:
        counter += 1

print(counter)
