# An open acre will become filled with trees if three or more adjacent acres contained trees. Otherwise, nothing happens.
# An acre filled with trees will become a lumberyard if three or more adjacent acres were lumberyards. Otherwise, nothing happens.
# An acre containing a lumberyard will remain a lumberyard if it was adjacent to at least one other lumberyard and at least one acre containing trees. Otherwise, it becomes open.

# Number of wood acres * number of lumberyards after 10 minutes
# . open ground
# | tree
# # lumberyard

OPEN = "."
TREE = "|"
LUMBER = "#"


def neighbours(field, row, col):
    trees, lumber = 0, 0
    for delta in ((-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)):
        neighbour = field[row + delta[0]][col + delta[1]]
        if neighbour == OPEN:
            continue
        elif neighbour == TREE:
            trees += 1
        elif neighbour == LUMBER:
            lumber += 1
        elif neighbour is None:
            continue
        else:
            raise ValueError
    return trees, lumber


def print_field(field):
    return "".join(
        c for row in field[1 : len(field) - 1] for c in row[1 : len(row) - 1]
    )


def step(data):
    next_gen = [line[:] for line in data]
    for i_row in range(1, len(next_gen) - 1):
        for i_col in range(1, len(next_gen[0]) - 1):
            trees, lumber = neighbours(data, i_row, i_col)
            cell = data[i_row][i_col]
            if cell == OPEN:
                if trees >= 3:
                    next_gen[i_row][i_col] = TREE
            elif cell == TREE:
                if lumber >= 3:
                    next_gen[i_row][i_col] = LUMBER
            elif cell == LUMBER:
                if lumber == 0 or trees == 0:
                    next_gen[i_row][i_col] = OPEN
    return next_gen


def count(data):
    trees = sum(1 for row in data for c in row if c == TREE)
    lumber = sum(1 for row in data for c in row if c == LUMBER)
    return trees * lumber


data = [[None] + list(line.rstrip("\n")) + [None] for line in open("input.txt")]
empty_row = [None for _ in range(len(data[0]))]
data = [empty_row] + data + [empty_row]

cycle_start, cycle_end = None, None
seen = {}
i = 0

while True:
    i += 1
    data = step(data)
    as_str = print_field(data)
    if seen.get(as_str, None) is not None:
        cycle_start = seen[as_str]
        cycle_end = i
        break
    seen[as_str] = i
    if i == 10:
        print(count(data))


target = int(1e9)
remaining = (target - cycle_start) % (cycle_end - cycle_start)
for _ in range(remaining):
    data = step(data)
print(count(data))
