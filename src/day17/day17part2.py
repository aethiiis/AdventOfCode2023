import timeit
from heapq import heappush, heappop
def part2():

    tab = [list(map(int, line.strip())) for line in open("input_day17_test", "r")]

    parcourus = set()
    frontier = [(0, 0, 0, 0, 0, 0)] #cost, x, y, dx, dy, consecutive steps

    while frontier:
        cost, x, y, dx, dy, con = heappop(frontier)

        if x == len(tab)-1 and y == len(tab[0])-1 and con >= 4:
            print(cost)
            break

        if (x, y, dx, dy, con) in parcourus:
            continue

        parcourus.add((x, y, dx, dy, con))

        if con < 10 and (dx, dy) != (0, 0):
            next_X = x + dx
            next_Y = y + dy
            if 0 <= next_X < len(tab) and 0 <= next_Y < len(tab[0]):
                heappush(frontier, (cost + tab[next_X][next_Y], next_X, next_Y, dx, dy, con+1))
        
        if con >= 4 or (dx, dy) == (0, 0):
            for next_dx, next_dy in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
                if (next_dx, next_dy) != (dx, dy) and (next_dx, next_dy) != (-dx, -dy):
                    next_X = x + next_dx
                    next_Y = y + next_dy
                    if 0 <= next_X < len(tab) and 0 <= next_Y < len(tab[0]):
                        heappush(frontier, (cost + tab[next_X][next_Y], next_X, next_Y, next_dx, next_dy, 1))

part2()