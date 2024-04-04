# HyperLogLog

This is a toy program used to verify my understanding of the HyperLogLog algorithm.

The implementation is almost as the original paper of HyperLogLog.
However, in order to get better results, a few changes are made to the original algorithm.
For example, I compute the `j`(the index of the register) using the rightmost `b` bits of the hash value instead of
the leftmost `b` bits. The reason is that when I was debugging,
I noticed that the non-carefully selected hash function may cause
the first `b` bits of the hash to be usually only 1 to 2 bits, which has a large impact on the result.

The functions used to generate the test data were written by ChatGPT, readers may choose other methods to get the test
dataset.

I also compared this program with Redis' HyperLogLog using the same test data, see the `hyperloglog_test.go`.

## References

- Flajolet, Philippe, et al. "Hyperloglog: the analysis of a near-optimal cardinality estimation algorithm." Discrete mathematics & theoretical computer science Proceedings (2007).
- Heule, Stefan, Marc Nunkesser, and Alexander Hall. "Hyperloglog in practice: Algorithmic engineering of a state of the art cardinality estimation algorithm." Proceedings of the 16th International Conference on Extending Database Technology. 2013.