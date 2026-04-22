<?php

$numbers = [1, 2, 3, 4, 5];

echo "Sum: " . array_sum($numbers) . "\n";
echo "Min: " . min($numbers) . "\n";
echo "Max: " . max($numbers) . "\n";
echo "Has 3: " . (in_array(3, $numbers) ? "yes" : "no") . "\n";
echo "Keys: " . implode(", ", array_keys(["a" => 1, "b" => 2])) . "\n";