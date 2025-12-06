input_test = "11-22,95-115,098-1012,1188511880-1188511890,222220-222224, 1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

input = open("input.txt").read()

def part_one(input: str):
    res = []
    for line in input.splitlines():
        for r in line.split(","):
            r = r.strip()
            if not r: 
                continue
            left, right = map(int, r.split("-"))
            res.append((left, right))

    total = 0
    for i in range(1, 100000):
        x = int(str(i)*2)
        for l, r in res:
            if l <= x <= r:
                total += x
    print("part one:", total)

def part_two(input: str):
    res = []
    for line in input.splitlines():
        for r in line.split(","):
            r = r.strip()
            if not r: 
                continue
            left, right = map(int, r.split("-"))
            res.append((left, right))

    s = set()
    for i in range(1, 100000):
        for j in range(2, 10):
            x = int(str(i)*j)
            for l, r in res:
                if l <= x <= r:
                    s.add(x)
                    break

    print("part two: ", sum(s))
                
part_one(input)
part_two(input)
