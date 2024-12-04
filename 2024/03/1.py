import re

print(
    sum(
        int(x) * int(y)
        for match in re.findall(
            r"mul\(\d+,\d+\)|do\(\)|don't\(\)", open("input.txt").read()
        )
        if (
            flag := (match == "do()")
            or (match != "don't()" and globals().get("flag", True))
        )
        and match.startswith("mul(")
        for x, y in [match[4:-1].split(",")]
    )
)
