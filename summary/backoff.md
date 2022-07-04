## 回溯
用于求所有组合

模板

```python
result = []

def backtrack(路径, 选择列表):
    if 满⾜结束条件:
        result.add(路径)
    return

    for 选择 in 选择列表:
    做选择
    backtrack(路径, 选择列表)
    撤销选择

```

