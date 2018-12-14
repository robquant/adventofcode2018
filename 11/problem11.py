
    # Find the fuel cell's rack ID, which is its X coordinate plus 10.
    # Begin with a power level of the rack ID times the Y coordinate.
    # Increase the power level by the value of the grid serial number (your puzzle input).
    # Set the power level to itself multiplied by the rack ID.
    # Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
    # Subtract 5 from the power level.


# (3, 5) 8 -> 4

def power(coords, serial):
    x, y = coords
    cell_rack_id = x + 10
    power_level = cell_rack_id * y
    power_level += serial
    power_level *= cell_rack_id
    power_level = power_level // 100
    power_level = power_level % 10
    return power_level - 5

def test_power():
    assert power((122,79), 57) == -5
    assert power((217,196), 39) == 0
    assert power((101,153), 71) == 4


def cumsum(field):
    res = []
    for row in field:
        s = 0
        rowsum = []
        for el in row:
            s += el
            rowsum.append(s)
        res.append(rowsum)
    return res

def max_rectangle(cumsums, size):
    size_x, size_y = size
    max_sum = -9999
    max_coords = None
    for x in range(0, 301 - size_x):
        for y in range(0, 301 - size_y):
            s = 0
            for j in range(y, y + size_y):
                if x == 0:
                    s += cumsums[j][x + size_x - 1]
                else:
                    s += cumsums[j][x + size_x - 1] - cumsums[j][x - 1]
            if s > max_sum:
                max_sum = s
                max_coords = (x, y)
    return max_coords, max_sum


def fill_field(serial):
    field = [[0] * 300 for _ in range(300)]
    for x in range(1, 301):
        for y in range(1, 301):
            field[y-1][x-1] = power((x,y), serial)
    return field

if __name__ == "__main__":
    serial = 2187
    field = fill_field(serial)
    cumsum_field = cumsum(field)
    coords, max_sum = max_rectangle(cumsum_field, (3,3))
    x,y = coords
    print(x+1, y+1)

    max_coords = None
    max_sum = -9999
    rect_size = None
    for size in range(1, 301):
        print(size)
        coords, sum = max_rectangle(cumsum_field, (size, size))
        if sum > max_sum:
            max_sum = sum
            max_coords = coords
            rect_size = size
    x,y = max_coords
    print(x+1, y+1, rect_size)
