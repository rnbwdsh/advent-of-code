import functools
import itertools
import json
import math
import re
import sys
from collections import Counter, defaultdict
from copy import deepcopy, copy
from dataclasses import dataclass
from functools import partial
from multiprocessing import Pool
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
from tqdm.auto import tqdm

# to remove unused import warnings
_ = sys, itertools, re, json, functools, math, numba, np, pd, nx, plt, Tuple, List, Dict, Set, Callable, Optional, Union, Any, Iterable, Iterator, TypeVar, Generic, Type, Generator, deepcopy, Counter, defaultdict, z3, frozendict, windowed, convolve2d, dataclass, tqdm, copy, Pool  # noqa


def level_ab(day: int, test: Tuple = None, levels=(0, 1), quiet=False, sep="\n", apply=lambda a: a):
    def parse(dd):
        return array([apply(d) for d in dd.split(sep)])

    def inner(solver):
        for level in levels:
            # solve with testdata
            if test and (actual := test[1 + level]) is not None:
                result = solver(parse(test[0]), level)
                assert result == actual, f"Test failed {result} != {actual} with data\n{repr(test[0])}"
                print("✓", end=" ")

            # solve real
            sol = solver(parse(get_data(day=day)), level)
            submit(sol, 'ab'[level], day=day, quiet=quiet)
        return solver  # return the original function, so you can put another annotation on it

    return inner


level_a = partial(level_ab, levels=(0,))
level_b = partial(level_ab, levels=(1,))
