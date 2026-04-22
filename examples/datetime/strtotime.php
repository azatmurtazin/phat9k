<?php

$tomorrow = strtotime("+1 day");
echo date("Y-m-d", $tomorrow) . "\n";

$nextWeek = strtotime("+1 week");
echo date("Y-m-d", $nextWeek) . "\n";

$timestamp = mktime(0, 0, 0, 1, 1, 2025);
echo date("Y-m-d", $timestamp) . "\n";
