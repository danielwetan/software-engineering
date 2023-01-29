<?php 

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;


class HelloController extends AbstractController
{

  #[Route('/hello')]
  public function greet(): Response 
  {
    return new Response(
      '<html><body>Hello World</body></html>'
    );
  }

  #[Route('/profile')]
  public function profile(): Response 
  {
    return $this->json([
      'message' => 'Welcome to your new controller!',
      'path' => 'src/Controller/HelloController.php',
  ]);
  }
}
