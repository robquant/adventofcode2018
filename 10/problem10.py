
import sys

def parse_line(line):
    first_open_bracket = line.find("<")
    first_comma = line.find(",")
    first_close_bracket = line.find(">")
    second_open_bracket = line.find("<", first_close_bracket)
    second_comma = line.find(",", second_open_bracket)
    second_close_bracked = line.find(">", second_comma)
    pos = (int(line[first_open_bracket+1:first_comma]), int(line[first_comma+1: first_close_bracket]))
    vel = (int(line[second_open_bracket+1:second_comma]), int(line[second_comma+1: second_close_bracked]))
    return pos, vel

def move(particle_pos, particle_vel):
    return [(p[0]+v[0], p[1]+v[1]) for p, v in zip(particle_pos, particle_vel)]

def print_pos(particle_pos):
    min_x = min(p[0] for p in particle_pos)
    max_x = max(p[0] for p in particle_pos)
    min_y = min(p[1] for p in particle_pos)
    max_y = max(p[1] for p in particle_pos)
    particles = set(particle_pos)
    for y in range(min_y, max_y + 1):
        for x in range(min_x, max_x):
            if (x, y) in particles:
                sys.stdout.write("#")
            else:
                sys.stdout.write(".")
        if (max_x, y) in particles:
            sys.stdout.write("#\n")
        else:
            sys.stdout.write(".\n")

def extend(particle_pos):
    min_x, min_y = particle_pos[0]
    max_x, max_y = particle_pos[0]
    for p in particle_pos:
        if p[0] < min_x:
            min_x = p[0]
        if p[0] > max_x:
            max_x = p[0]
        if p[1] < min_y:
            min_y = p[1]
        if p[1] > max_y:
            max_y = p[1]
    return abs(max_x - min_x) + abs(max_y - min_y)

particle_pos, particle_vel = [], []
for line in open("input.txt"):
    pos, vel = parse_line(line)
    particle_pos.append(pos)
    particle_vel.append(vel)

last_extend = 1e9
last_pos = []
nsteps = 0
while last_extend > extend(particle_pos):
    last_extend = extend(particle_pos)
    last_pos = particle_pos[:]
    particle_pos = move(particle_pos, particle_vel)
    nsteps += 1
print_pos(last_pos)
print(nsteps-1)