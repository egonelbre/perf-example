# A session on low-level optimizations

* [measuring time](./measure_time)
* [writing benchmarks](./writing_benchmarks)
* [inspecting assembly](./inspecting_assembly)
* [basic profiling](./profiling)
* [branch cost](./branch)
* [call cost](./call)
* [dispatch call cost](./dispatch)
* [unintentional copy](./unintentional_copy)
* [pointer cost](./pointers)
* [memory cost](./memory)
* [bounds check cost](./bounds_checks)
* [unrolling](./unrolling)
* [loop alignment](./loop_alignment)

## Additional Videos / Articles

* Intuitive Performance https://www.youtube.com/watch?v=51ZIFNqgCkA
* https://egonelbre.com/a-tale-of-bfs/
* https://egonelbre.com/a-tale-of-bfs-going-parallel/

## Recommended

* https://github.com/dgryski/go-perfbook
* https://www.computerenhance.com/

## Tools

* Benchmark statistical analysis (https://golang.org/x/perf/cmd/benchstat)
* Assembly and Code viewer (https://github.com/loov/lensm)
* Visualizing bounds checks (https://github.com/loov/view-annotated-file)
* AMD μProf (https://www.amd.com/en/developer/uprof.html)
* Intel VTune Profiler (https://www.intel.com/content/www/us/en/developer/tools/oneapi/vtune-profiler.html)
* Apple Instruments (https://help.apple.com/instruments/mac/current/#)