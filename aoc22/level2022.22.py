import numpy as np
import re

from dataclasses import dataclass
from typing import Tuple, Dict, List, NamedTuple

from lib import level_ab  # annotation for loading test data and submitting solutions


# based on https://github.com/blubber-rubber/AdventOfCode2022/blob/main/Day22/part2.py
ROTATIONS = ['^', '>', 'v', '<']
DIR = [(1, 0), (0, 1), (-1, 0), (0, -1)]


class Symbol(NamedTuple):  # allows indexing and regular property accesss
    name: str
    orientation: str

    @property
    def next_rotation(self) -> 'Symbol':
        r_index = ROTATIONS.index(self.orientation)
        return Symbol(self.name, ROTATIONS[(r_index + 1) % 4])


@dataclass
class Stamp:
    _name_orientation: Symbol | Tuple[str, str]
    _neighbors: List[Symbol | Tuple[str, str]]

    def __post_init__(self):  # map to Symbol
        self._name_orientation = Symbol(*self._name_orientation)
        self._neighbors = [Symbol(*n) for n in self._neighbors]

    def rotate_until(self, name_orientation: Symbol):
        while self._name_orientation != name_orientation:
            self._name_orientation = self._name_orientation.next_rotation
            self._neighbors = [n.next_rotation for n in self._neighbors]
            self._neighbors = self._neighbors[-1:] + self._neighbors[:-1]
        return self._neighbors


stamps = {'U': Stamp(('U', '^'), [('L', '>'), ('B', 'v'), ('R', '<'), ('F', '^')]), # noqa
          'D': Stamp(('D', '^'), [('R', '<'), ('B', '^'), ('L', '>'), ('F', 'v')]), # noqa
          'F': Stamp(('F', '^'), [('L', '^'), ('U', '^'), ('R', '^'), ('D', 'v')]), # noqa
          'B': Stamp(('B', '^'), [('R', '^'), ('U', 'v'), ('L', '^'), ('D', '^')]), # noqa
          'L': Stamp(('L', '^'), [('B', '^'), ('U', '<'), ('F', '^'), ('D', '<')]), # noqa
          'R': Stamp(('R', '^'), [('F', '^'), ('U', '>'), ('B', '^'), ('D', '>')])} # noqa


@dataclass
class Face:
    pos_sym: Symbol
    world: np.ndarray
    pos_abs: Tuple[int, int]

    def __post_init__(self):
        self.pos_sym = Symbol(*self.pos_sym)
        self.orientation_original = self.pos_sym.orientation

    def rotate(self):
        self.pos_sym = self.pos_sym.next_rotation
        self.world = np.rot90(self.world, k=-1)

    def rotate_to_original(self, x: int, y: int, d: int, side_length: int) -> Tuple[int, int, int]:
        while self.pos_sym.orientation != self.orientation_original:
            self.rotate()
            x, y = side_length - 1 - y, x
            d = (d + 1) % len(DIR)
        return x + self.pos_abs[0] + 1, y + self.pos_abs[1] + 1, d


def create_faces(grid: np.ndarray, side: int) -> Dict[str, Face]:
    """ Generates a dictionary that maps a face to the mini-square it corresponds to in the grid
    and its orientation in the grid. It then """
    w, h = grid.shape
    pos_square = {(y, x): grid[x:x + side, y:y + side]
                  for x in range(0, w, side)
                  for y in range(0, h, side)
                  if (grid[x:x + side, y:y + side] != " ").all()}

    # positions of squares in the grid
    p0 = min(pos_square, key=lambda x: (x[1], x[0]))
    position_queue = [p0]  # positions to be stamped
    mappings = {p0: Symbol("U", "^")}  # Mini-square to face and orientation
    faces = {"U": Face(("U", "^"), pos_square[p0], p0)}  # noqa E203

    seen = set()  # positions already stamped
    while position_queue:  # Still positions to be stamped
        pos = x, y = position_queue.pop(0)
        seen.add(pos)
        face = mappings[pos][0]
        neighbour_faces = stamps[face].rotate_until(mappings[pos])
        for i, name_orientation in enumerate(neighbour_faces):
            pos = (x + DIR[i][0] * side, y + DIR[i][1] * side)
            if pos in pos_square and pos not in seen:  # Check if we need to stamp the neighbours
                position_queue.append(pos)
                mappings[pos] = name_orientation
                faces[name_orientation[0]] = Face(name_orientation, pos_square[pos], pos)  # noqa E203
    return faces


@level_ab(22, test=('        ...#\n        .#..\n        #...\n        ....\n...#.......#\n........#...\n..#....#....\n..........#.\n        ...#....\n        .....#..\n        .#......\n        ......#.\n\n10R5L5R10L4R5L5', 6032, 5031), sep="\n")
def solve(lines, level):
    # Pad lines to make them all the same length (required for numpy)
    line_len = max(len(line) for line in lines[:-2])
    world = np.array([list(line + " " * (line_len - len(line))) for line in lines[:-2]])
    side_length = int(((world != " ").sum() / 6) ** 0.5)
    d, face, faces = 0, None, []  # to avoid warnings
    if level:
        faces = create_faces(world, side_length)
        face = faces["U"]
        y = x = 0
    else:
        world = np.pad(world, 1, constant_values=" ")
        y, x = list(zip(*np.where("." == world)))[0]  # noqa  # numpy typing error

    for token in re.findall(r"(\d+|[RL])", lines[-1]):
        if not token.isnumeric():  # Rotation direction
            d += [-1, 1][token == "R"]
            d %= 4
            continue
        for _ in range(int(token)):  # stepping instruction
            xn, yn = x + DIR[d][0], y + DIR[d][1]
            if level:  # move on face
                if 0 <= xn < side_length and 0 <= yn < side_length:  # Check if we are crossing an edge
                    face_next = face  # no edge crossed, face stays same
                else:  # An edge was crossed
                    face_next, direction = stamps[face.pos_sym.name].rotate_until(face.pos_sym)[d]
                    face_next = faces[face_next]
                    while face_next.pos_sym.orientation != direction:  # Align the neighbouring face with stamp
                        face_next.rotate()
                    xn %= side_length
                    yn %= side_length
                if face_next.world[yn][xn] == "#":
                    break
                face = face_next
            else:
                if world[yn, xn] == " ":
                    while world[yn - DIR[d][1], xn - DIR[d][0]] != " ":
                        xn -= DIR[d][0]
                        yn -= DIR[d][1]
                if world[yn, xn] == "#":
                    break
            x, y = xn, yn

    if level:
        x, y, d = face.rotate_to_original(x, y, d, side_length)

    return y * 1000 + x * 4 + d