<?php

$array = [1, 2, 3, 4, 5];
echo "Count: " . count($array) . "\n";

$associative = ["name" => "John", "age" => 30];
echo "Name: " . $associative["name"] . "\n";
echo "Age: " . $associative["age"] . "\n";

$array[] = 6;
echo "Last: " . end($array) . "\n";