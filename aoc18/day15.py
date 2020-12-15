import sys
from collections import deque
from dataclasses import dataclass

gob_attack = 3
try:    elf_attack = int(sys.argv[1])
except: elf_attack = gob_attack

@dataclass(order=True)
class Unit:
    y : int
    x : int
    hp : int
    typ : str
    @property
    def attack(self):
        return elf_attack if self.typ == 'E' else gob_attack
    def dist(a,b):
        return abs(a.x-b.x) + abs(a.y-b.y)
    def __str__(self):
        return f'[{self.typ}: {self.hp}]'

G = []
E = []
M = []
for i,line in enumerate(sys.stdin):
    line = line.strip()
    for j,c in enumerate(line):
        if c == 'G':
            G.append(Unit(i,j,200,'G'))
        if c == 'E':
            E.append(Unit(i,j,200,'E'))
    M.append([int(c == '#') for c in line])

elfded = False

def viz():
    H = len(M)
    W = len(M[0])
    Mc = [m[:] for m in M]

    units = sorted(G + E)

    for e in units:
        if e.typ == 'E': Mc[e.y][e.x] = 3
        else:            Mc[e.y][e.x] = 2

    by_row = [[] for _ in range(H)]
    for f in sorted(G + E):
        by_row[f.y].append(f'[{f.typ}: {f.hp}]')

    for i in range(H):
        s = ''.join('.#GE'[Mc[i][j]] for j in range(W))
        print(s, ' '.join(by_row[i]))
    print()

def neigh(x,y):
    return [(x,y-1),(x-1,y),(x+1,y),(x,y+1)]

turns = 0
while E and G:
    units = sorted(G+E)

    for unit in units:
        if unit.typ == 'X': continue

        enemies = sorted(e for e in units if e.typ != unit.typ)
        targets = set().union(*(set(neigh(e.x,e.y)) for e in enemies if e.typ != 'X'))

        cands = []
        goald = 1e9
        blocked = {(e.x,e.y) for e in units if e.typ != 'X'}
        Q = deque([(unit.x,unit.y,0,None)])
        while Q:
            x,y,d,step1 = Q.popleft()
            if d > goald: break
            if M[y][x] or ((x,y) in blocked and d > 0): continue
            blocked.add((x,y))

            if d == 1:
                step1 = (x,y)
            if (x,y) in targets:
                goald = d
                cands.append([y,x,d,step1])
            for s,t in neigh(x,y):
                Q.append((s,t,d+1,step1))

        if len(cands) == 0: continue

        x,y,d,step1 = min(cands)
        if d > 0:
            unit.x,unit.y = step1
        if d <= 1:
            target = min((e for e in enemies if e.typ != 'X' and e.dist(unit) == 1),
                         key=lambda x: x.hp)
            target.hp -= unit.attack
            if target.hp <= 0:
                if target.typ == 'E':
                    elfded = True
                target.typ = 'X'

    # Purge the dead
    E = [e for e in E if e.typ != 'X']
    G = [e for e in G if e.typ != 'X']

    turns += 1
    print(f'=== round {turns} ===')
    viz()

hp = sum(e.hp for e in G) + sum(e.hp for e in E)
print(f'Result: {(turns-1)*hp}')
print(f'Did an elf die? {"YNeos"[1-elfded::2]}')