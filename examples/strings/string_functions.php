<?php

$str = "Hello, World!";
echo strlen($str) . "\n";

$upper = strtoupper($str);
echo $upper . "\n";

$lower = strtolower($str);
echo $lower . "\n";

$substring = substr($str, 0, 5);
echo $substring . "\n";

$pos = strpos($str, "World");
echo $pos . "\n";
