LFSR - Linear Feedback Shift Register
===

A linear feedback shift register is a collection of bits that shifts when triggered, and the next state is a linear function of its previous state. We use [right-shift](https://en.wikipedia.org/wiki/Logical_shift) (`>>`) as the shift operation and [XOR](https://en.wikipedia.org/wiki/Exclusive_or) (`^`) as the linear function to generate the next state of the register.

This repository contains a demonstration of LFSR in

 - generating random bits
 - generating random numbers
 - scrambling and unscrambling a stream of bytes

You can read more about LFSRs on [Pseudorandom Number Generation using LFSR](https://arpitbhayani.me/blogs/lfsr).

# Sample Usage

## Generating Random Bits

The following command generates `10` random bits using a `4`-bit LFSR with
Seed `0b1001`, tap at bit position `1`.

```
$ go run main.go random-bits -s 9 -t 1 -n 4 --count 10
1
0
0
1
1
0
1
0
1
1
```

## Generating Random 8-bit Numbers

The following command generates `10` random `8`-bit numbers using a `4`-bit LFSR with
Seed `0b1001`, tap at bit position `1`.

```
$ go run main.go random-numbers -s 9 -t 1 -n 4 -k 8 --count 10
154
241
53
226
107
196
215
137
175
19
```

## Scrambling a text file

To see scrambling in action let's first create a text file with some
secret content

```
$ echo "british troops entered cuxhaven at 1400 on 6 may - from now on all radio traffic will cease - wishing you all the best. lt kunkel." > secret.txt
```

Now fire the following command that scrambles the file with the default configurations

```
$ go run main.go scramble -f secret.txt
```

The command creates a scrambled file with the name `enc.secret.txt` and if we `cat` the scrambled file, we get the following jargon

```
$ cat enc.secret.txt
D���#РɃ��az�$B���?��9��R���,7�#Vhq�%��\�I>����ӻ��jq_�c`��yp��VM�u/�|���e!�%
```

## Unscrambling a scrambled file

To unscramble the scrambled file and reveal the secret content fire the following command

```
$ go run main.go unscramble -f enc.secret.txt
```

The command unscrambles the scrambled file and generates another file with the name `dec.enc.secret.txt` and if we `cat` it, we get the original content

```
$ cat dec.enc.secret.txt
british troops entered cuxhaven at 1400 on 6 may - from now on all radio traffic will cease - wishing you all the best. lt kunkel.
```
