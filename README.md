# truerandom

This is a simple go package for generating true random numbers, provided by random.org.

It creates integer and float64 random numbers within a given range.

## Usage

you first need to create a generator object:

```gen, err := NewIntGenerator(min, max, num)```

or

```gen, err := NewFloatGenerator(min, max, num)```

### Parameters

```min```: the minimum random value created

```max```: the maximum random value created

```num```: the amount of numbers created internally

### example

if ```min``` is 4 and ```max``` is 45, the smallest random numer is 4 and the largest is 45.

the ```num``` parameters describes how many random numbers are fetched from random.org in one go. If you ask for more numbers then initially fetched, the package fetches the next ```num``` numbers.
It is planned to do that asynchonously but this is not yet implemented.
