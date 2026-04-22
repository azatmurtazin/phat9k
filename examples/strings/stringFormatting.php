<?php

$name = "John";
$age = 30;

echo sprintf("My name is %s and I am %d years old\n", $name, $age);

$formatted = number_format(1234567.89123, 2);
echo $formatted . "\n";

$trimmed = trim("  hello  ");
echo ":" . $trimmed . ":\n";
