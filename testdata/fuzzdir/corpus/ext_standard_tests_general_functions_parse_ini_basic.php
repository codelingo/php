<?php

$ini_file = dirname(__FILE__)."/parse_ini_basic.data";

define('TEST_CONSTANT', 'this_is_test_constant');

var_dump(parse_ini_file($ini_file, 1));

echo "Done.\n";
?>
