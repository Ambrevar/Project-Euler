# Project Euler Solutions

The [Project Euler](https://projecteuler.net/) is a website hosting a serie of
computational problems. This repository also mirrors all the problem statements
(with some delay).

## Statement fetcher

Use `fetch.sh` to fetch all the problem statements locally.
See `fetch.sh -h` for more details.

## Objectives

The main objective is to be smart and to grasp a deep understanding of every
problem.

* Write the fastest possible code with a variable limit condition (embodied in
the 'limit' variable). The variable limit is to make clear that nothing is
precomputed.

* The only output should be the result, except for exercises that can be derived
with "pen and paper" and that displays some nice theoretical value.

* Code should be clear and short.

* Only the algorithm is worked out to be optimal, avoid language
implementation-specific optmization. Hardware optimizations are usually
discarded, although some of them at the core of computing can be used (e.g.
binary operations).

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
shot. Indeed, some wrong algorithm can output the right result.

* For some exercises requiring big numbers, assisting the resolution with a big
number library makes the resolution trivial and pointless: the goal is to either
avoid going through big numbers manipulation, or to find an efficient way to
store them.

## Note

A lot of solutions available out there on the Internet are plain brute force: I
think the authors missed the point. Many problems suggested on Project Euler are
rather simple to implement using a naive approach. The real purpose of the
project is to make you think of a fast and elegant resolution.

## Contribution

Feedback _on solved problems_ is very welcome. Please do not send me anything on
problems yet to be solved.
