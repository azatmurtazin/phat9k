<?php

function sum(...$numbers) {
    $total = 0;
    foreach ($numbers as $n) {
        $total += $n;
    }
    return $total;
}

echo sum(1, 2, 3, 4, 5) . "\n";
