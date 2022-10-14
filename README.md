# ⚡ godis ⚡
Minimal go redis clone

### Benchmark

I'm using memtier_benchmark (from Redis) to perform benchmark tests

`memtier_benchmark –s $SERVER_IP -t 8 -c 16 --test-time=30 --distinct-client-seed -d 256 --pipeline=30`

#### Machine specs

```
OS: Debian GNU/Linux 11 (bullseye) on Windows 10 x86_64
Kernel: 5.10.102.1-microsoft-standard-WSL2
CPU: AMD Ryzen 5 3600 (12) @ 3.6GHz
GPU: GeForce RTX 2060
Memory: 16GiB
```

#### Run command

`go run main.go`

#### Results

```
============================================================================================================================
Type         Ops/sec     Hits/sec   Misses/sec    Avg. Latency     p50 Latency     p99 Latency   p99.9 Latency       KB/sec
----------------------------------------------------------------------------------------------------------------------------
Sets        65244.72          ---          ---         5.19180         2.86300        29.69500       380.92700     19362.43
Gets       652426.29     61776.51    590649.78         5.16536         2.86300        29.43900       380.92700     41039.76
Waits           0.00          ---          ---             ---             ---             ---             ---
---
Totals     717671.01     61776.51    590649.78         5.16777         2.86300        29.43900       380.92700     60402.20
```
