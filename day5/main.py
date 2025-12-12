import collections

input = open("./input.txt").read()
lines = input.splitlines()


def part_one(input):
    seen = set()
    fresh_ids = []
    ranges = []

    saw_middle = False
    for i in range(0, len(input)):
        if lines[i] == "":
            saw_middle = True
        elif saw_middle and lines[i] != "":
            fresh_ids.append(lines[i])
        else:
            ranges.append(lines[i])

    for i in range(len(ranges)):
        for j in range(len(fresh_ids)):
            lower, upper = map(int, lines[i].split("-"))
            if lower <= int(fresh_ids[j]) <= upper:
                seen.add(fresh_ids[j])
    return len(seen), ranges


def part_two(ranges):
    count = 0
    s = collections.Counter()

    for line in ranges:
        lower, upper = map(int, line.split("-"))
        s[lower] += 1
        s[upper + 1] -= 1

    prev = None
    current = 0

    for key in sorted(s.keys()):
        p = current
        current += s[key]

        if p == 0:
            prev = key
            continue
        count += key - prev
        prev = key

    return count


res, ranges = part_one(lines)
print(part_two(ranges))
