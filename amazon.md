1. Search Directly 
```python 
import os
import re
from abc import ABC, abstractmethod
from typing import List

# --- STRATEGY ---
class Filter(ABC):
    @abstractmethod
    def matches(self, file_path: str) -> bool:
        pass

class NameFilter(Filter):
    def __init__(self, pattern: str):
        self.regex = re.compile(pattern)

    def matches(self, file_path: str) -> bool:
        return bool(self.regex.search(os.path.basename(file_path)))

class SizeFilter(Filter):
    def __init__(self, min_size: int = 0, max_size: int = float('inf')):
        self.min = min_size
        self.max = max_size

    def matches(self, file_path: str) -> bool:
        size = os.path.getsize(file_path)
        return self.min <= size <= self.max

# --- COMPOSITE ---
class CompositeFilter(Filter):
    def __init__(self, filters: List[Filter], op: str = "AND"):
        self.filters = filters
        self.op = op.upper()

    def matches(self, file_path: str) -> bool:
        if self.op == "AND":
            return all(f.matches(file_path) for f in self.filters)
        elif self.op == "OR":
            return any(f.matches(file_path) for f in self.filters)
        else:
            raise ValueError(f"Unknown op {self.op}")

# --- SEPARATION OF CONCERNS ---
class DirectorySearcher:
    def __init__(self, root: str):
        self.root = root

    def search(self, filter_obj: Filter) -> List[str]:
        matched = []
        for dirpath, _, filenames in os.walk(self.root):
            for fname in filenames:
                path = os.path.join(dirpath, fname)
                if filter_obj.matches(path):
                    matched.append(path)
        return matched

# Example Usage
if __name__ == "__main__":
    f = CompositeFilter([
        NameFilter(r".*\.log$"),
        SizeFilter(0, 5_000_000)
    ], op="AND")

    searcher = DirectorySearcher("/var/log")
    result = searcher.search(f)
    print(result)

``` 
