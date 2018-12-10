def parse_node(input, index):
    nchildren = input[index]
    index += 1
    nmeta = input[index]
    index += 1
    children = []
    for _ in range(nchildren):
        child, index = parse_node(input, index)
        children.append(child)
    return (nchildren, nmeta, children, input[index:index + nmeta]), index+nmeta 

def meta(node):
    return node[3]

def children(node):
    return node[2]

def walk(node, op):
    yield op(node)
    for child in children(node):
        yield from walk(child, op)

def walk_indexed(node):
    if len(children(node)) == 0:
        return sum(meta(node))
    s = 0
    for i in meta(node):
        if i <= node[0]:
            s += walk_indexed(children(node)[i-1])
    return s


    

input = list(map(int, open("input.txt").read().split()))
root = parse_node(input, 0)[0]

print(sum(walk(root, lambda node: sum(meta(node)))))
print(walk_indexed(root))
