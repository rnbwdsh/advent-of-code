import functools
import itertools
import json
import math
import re
from collections import Counter, defaultdict
from copy import deepcopy
from dataclasses import dataclass
from functools import partial
from typing import Tuple, List, Dict, Set, Callable, Optional, Union, Any, Iterable, Iterator, TypeVar, Generic, Type, \
    Generator

import matplotlib.pyplot as plt
import networkx as nx
import numba
import numpy as np
import pandas as pd
import z3
from aocd import submit, get_data
from frozendict import frozendict
from more_itertools import windowed
from numpy import array
from scipy.signal import convolve2d

# to remove unused import warnings
_ = [itertools, re, json, functools, math, numba, np, pd, nx, plt, Tuple, List, Dict, Set, Callable, Optional, Union, Any, Iterable, Iterator, TypeVar, Generic, Type, Generator, deepcopy, Counter, defaultdict, z3, frozendict, windowed, convolve2d, dataclass]


def level_ab(day: int, test: Tuple = None, levels=(0, 1), quiet=False, sep="\n", apply=lambda a: a):
    def solver(solve):
        for level in levels:
            parse = lambda dd: array([apply(d) for d in dd.split(sep)])

            # solve with testdata
            if test and (actual := test[1 + level]) is not None:
                result = solve(parse(test[0]), level)
                assert result == actual, f"Test failed {result} != {actual} with data\n{repr(test[0])}"
                print("✓", end=" ")

            # solve real
            sol = solve(parse(get_data(day=day)), level)
            submit(sol, 'ab'[level], day=day, quiet=quiet)
        return solve  # return the original function, so you can put another annotation on it

    return solver


level_a = partial(level_ab, levels=(0,))
level_b = partial(level_ab, levels=(1,))


@level_ab(19, test=("""
""", 1, 2), sep="\n")
def solve(lines: List[str], level=0):
    return None