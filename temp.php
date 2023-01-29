<?php 

// function box(string $func, mixed ...$args): mixed
// {
//     try {
//         return $func(...$args);
//     } finally {
//         restore_error_handler();
//     }
// }


// box('mkdir', 'abcd', 0777, true);

function box(string $func, $path)
{
  return $func($path);
}

box('mkdir', 'abcd');