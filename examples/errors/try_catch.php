<?php

try {
    $result = 10 / 2;
    echo "Result: " . $result . "\n";
} catch (Exception $e) {
    echo "Error: " . $e->getMessage() . "\n";
}

echo "Done\n";