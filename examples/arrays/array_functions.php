<?php

$numbers = [1, 2, 3, 4, 5];

$squared = array_map(function($n) {
    return $n * $n;
}, $numbers);

print_r($squared);

$evens = array_filter($numbers, function($n) {
    return $n % 2 == 0;
});

print_r($evens);
