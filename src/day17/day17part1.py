from queue import PriorityQueue

with open("input_day17", "r") as fichier :
    lines = fichier.readlines()
    lines = [line.strip() for line in lines]

class Grid:
    def __init__(self, lines):
        self.lignes = len(lines)
        self.colonnes = len(lines[0])
        self.grid = dict()
        for i in range(len(lines)):
            for j in range(len(lines[i])):
                self.grid[(i, j)] = int(lines[i][j])

    def neighbors(self, coord):
        ligne, colonne = coord
        voisins = [
            (ligne-1, colonne),
            (ligne+1, colonne),
            (ligne, colonne-1),
            (ligne, colonne+1)
        ]

        voisins_valides = [(x, y) for x, y in voisins if self.is_in_grid((x, y))]

        return voisins_valides
    
    def is_in_grid(self, coord):
        return 0 <= coord[0] < self.lignes and 0 <= coord[1] < self.colonnes

def heuristic(a, b):
    return abs(a[0] - b[0]) + abs(a[1] - b[1])

def update_steps(current_pos, current_con, next_pos):
    if current_pos[0] > next_pos[0] and current_con[0] != 0: #continuation haut
        return (current_con[0]+1, 0, 0, 0)
    if current_pos[0] < next_pos[0] and current_con[1] != 0: #continuation bas
        return (0, current_con[1]+1, 0, 0)
    if current_pos[1] > next_pos[1] and current_con[2] != 0: #continuation gauche
        return (0, 0, current_con[2]+1, 0)
    if current_pos[1] < next_pos[1] and current_con[3] != 0: #continuation droite
        return (0, 0, 0, current_con[3]+1)
    if current_pos[0] > next_pos[0]: #haut
        return (1, 0, 0, 0)
    if current_pos[0] < next_pos[0]: #bas
        return (0, 1, 0, 0)
    if current_pos[1] > next_pos[1]: #gauche
        return (0, 0, 1, 0)
    if current_pos[1] < next_pos[1]: #droite
        return (0, 0, 0, 1)
    

start = (0, 0)
goal = (len(lines)-1, len(lines[0])-1)
frontier = PriorityQueue()
frontier.put((start, (0, 0, 0, 0)), 0)
came_from = dict()
cost_so_far = dict()
came_from[start] = None
cost_so_far[start] = int(lines[start[0]][start[1]])
graph = Grid(lines=lines)
consecutiveSteps = dict()

while not frontier.empty():
    current = frontier.get()

    if current[0] == goal:
        break

    for next_pos in graph.neighbors(current[0]):
        new_cost = cost_so_far[current[0]] + graph.grid[next_pos]
        next_con = update_steps(current_pos=current[0], current_con=current[1], next_pos=next_pos)
        if (next_pos not in cost_so_far or new_cost < cost_so_far[next_pos]) and not any(element >= 3 for element in next_con):
            print(next_con)
            cost_so_far[next_pos] = new_cost
            priority = new_cost# + heuristic(goal, next_pos)
            frontier.put((next_pos, next_con), priority)
            came_from[next_pos] = current[0]
            

chemin = [goal]
while goal in came_from and came_from[goal] is not None:
    goal = came_from[goal]
    chemin.insert(0, goal)

print(chemin)
chemin_grid = [['.' for j in range(len(lines[0]))] for i in range(len(lines))]
count = 0
for element in chemin:
    count += int(lines[element[0]][element[1]])
    chemin_grid[element[0]][element[1]] = '#'

for ligne in chemin_grid:
    for element in ligne:
        
        print(element, end=' ')
    print()
print(count)