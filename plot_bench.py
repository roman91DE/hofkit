import os
import re
import sys
import platform
import subprocess
from datetime import datetime

import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

def run_benchmarks(output_path: str):
    print(f"Running benchmarks... saving to: {output_path}")
    os.makedirs(os.path.dirname(output_path), exist_ok=True)
    with open(output_path, "w") as f:
        subprocess.run(["go", "test", "-bench=.", "-benchmem"], stdout=f, check=True)

def parse_benchmark_output(filepath: str) -> pd.DataFrame:
    pattern = re.compile(
        r"^Benchmark(\w+)/(Hofkit|Native)-\d+\s+\d+\s+([\d.]+)\sns/op\s+([\d.]+)\sB/op\s+([\d.]+)\sallocs/op"
    )
    records = []

    with open(filepath) as f:
        for line in f:
            match = pattern.match(line)
            if match:
                func, impl, ns_op, b_op, allocs = match.groups()
                records.append({
                    "Function": func,
                    "Implementation": impl,
                    "ns/op": float(ns_op),
                    "B/op": float(b_op),
                    "allocs/op": float(allocs),
                })

    return pd.DataFrame(records)

def metric_title(metric: str) -> str:
    return {
        "ns/op": "Time per Operation (ns/op)",
        "B/op": "Memory Used per Operation (Bytes)",
        "allocs/op": "Memory Allocations per Operation",
    }.get(metric, metric)

def plot(df: pd.DataFrame, metric: str, output_dir: str):
    sns.set(style="whitegrid")
    plt.figure(figsize=(12, 6))
    sns.barplot(data=df, x="Function", y=metric, hue="Implementation")
    plt.title(f"Go Benchmark – {metric_title(metric)}")
    plt.ylabel(metric_title(metric))
    plt.tight_layout()
    plot_path = os.path.join(output_dir, f"benchmark_{metric.replace('/', '_')}.png")
    plt.savefig(plot_path)
    print(f"Saved plot: {plot_path}")
    plt.show()

def sanitize(s: str) -> str:
    return re.sub(r"[^\w]+", "_", s)

def main():
    goos = platform.system().lower()
    goarch = platform.machine().lower()
    cpu = sanitize(platform.processor() or platform.machine())
    timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
    out_dir = os.path.join("benchmark", f"{goos}_{goarch}_{cpu}_{timestamp}")
    bench_file = os.path.join(out_dir, "bench.txt")

    run_benchmarks(bench_file)
    df = parse_benchmark_output(bench_file)

    if df.empty:
        print("No benchmark results parsed.")
        sys.exit(1)

    print(df)

    for metric in ["ns/op", "B/op", "allocs/op"]:
        plot(df, metric, out_dir)

if __name__ == "__main__":
    main()