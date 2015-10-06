# Project Euler Solutions

The [Project Euler](https://projecteuler.net/) is a website hosting a series of
computational problems. This repository also mirrors all the problem statements
(with some delay).

## Statement fetcher

Use `fetch.sh` to fetch all the problem statements locally.
See `fetch.sh -h` for more details.

## Objectives

The main objective is to be smart and to grasp a deep understanding of every
problem.

* Write the fastest possible code under a limit condition. The 'limit' variable
is to make clear that nothing is precomputed.

* The only output should be the result.

* Some problem solution are not an algorithm but merely a mathematical
derivation (also known as a "pen and paper" solution).

* Code should be clear and short.

* Focus on the algorithm, avoid language implementation-specific optmization.
Hardware optimizations are usually discarded, although some of them at the core
of computing can be used (e.g. binary operations).

* Every implementation should run in less than a minute on any decent computer
as suggested on the home page. Strive to keep running-time below a second, or
even 10 ms on a modern computer. Compiler optimization is _disabled_.

* Explain the main algorithm in head comment if needed. Point out the tricks and
pitfalls, explain why some approaches would be slow.

* If possible or reasonable, write portable code, use the standard of the
language only, without external library nor external file. Every problem should
be straighforward to compile and execute.

* There might exist various solutions where complexity is unclear or where
performance depends on the input. There are usually trade-offs in memory/speed
or speed/code length. These implementations are kept together for comparison.

* Code has to be proven and should not be working just because of some lucky
shot. Indeed, some wrong algorithms can output the right result.

* Some problem evolve around big numbers, thus assisting the resolution with a
big number library makes the resolution trivial and pointless. The goal is to
either avoid going through big numbers manipulation, or to find an efficient way
to store them.

## Note

A lot of solutions available on the forum and out there on the Internet are
plain brute force: I think the authors missed the point. Many problems suggested
on Project Euler are rather simple to implement using a naive approach. The real
purpose of the project is to make you think of a fast and elegant resolution.

## Contribution

Feedback _on solved problems_ is very welcome. Please do not send me anything on
problems yet to be solved.
