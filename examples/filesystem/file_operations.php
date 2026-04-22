<?php

$filename = "test.txt";

if (file_exists($filename)) {
    echo "File exists\n";
    echo file_get_contents($filename);
} else {
    echo "Creating file\n";
    file_put_contents($filename, "Hello, World!");
}

$files = scandir(".");
print_r($files);
