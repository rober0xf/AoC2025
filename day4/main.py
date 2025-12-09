input = open("input.txt").read().splitlines()

matrix = [list(line) for line in input]


"""
    (-1,-1)    (-1,0)  (-1,1)

    (0,-1)     (0,0)   (0,1)

    (1,-1)     (1,0)   (1,1)
"""


def calculate_neighbors(matrix, row, col):
    n_adjacents = 0
    neighbors = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]  # moore

    for x, y in neighbors:
        current_row = row + x
        current_col = col + y
        if current_row >= len(matrix) or current_row < 0:
            continue
        if current_col >= len(matrix[0]) or current_col < 0:
            continue
        if matrix[current_row][current_col] == "@":
            n_adjacents += 1

    return n_adjacents


def part_one():
    result = 0
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if matrix[i][j] == "@" and calculate_neighbors(matrix, i, j) < 4:
                result += 1
    return result


def part_two():
    result = 0
    loop = True

    while loop:
        loop = False
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                if matrix[i][j] == "@" and calculate_neighbors(matrix, i, j) < 4:
                    result += 1
                    matrix[i][j] = "."
                    loop = True
    return result


print(part_one())
print(part_two())
