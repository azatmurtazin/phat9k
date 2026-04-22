<?php

$d = strtotime("2026-04-23");
echo "Tomorrow: " . date("Y-m-d", $d) . "\n";

$w = strtotime("2026-04-27");
echo "Next week: " . date("Y-m-d", $w) . "\n";

$d2 = mktime(0, 0, 0, 1, 1, 2026);
echo "Jan 1: " . date("Y-m-d", $d2) . "\n";