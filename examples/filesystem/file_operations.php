<?php

if (file_exists("/tmp/test.txt")) {
    echo "File exists\n";
} else {
    echo "File not found\n";
}

$dir = __DIR__;
$files = scandir($dir);
echo "Scanned " . count($files) . " items\n";