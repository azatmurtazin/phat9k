<?php

try {
    $result = 10 / 0;
} catch (Exception $e) {
    echo "Error: " . $e->getMessage() . "\n";
}

echo "Code continues...\n";
