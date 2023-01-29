<?php 

namespace App\Controller;

use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;


class HelloController 
{

  #[Route('/hello')]
  public function greet(): Response 
  {
    return new Response(
      '<html><body>Hello World</body></html>'
    );
  }
}
