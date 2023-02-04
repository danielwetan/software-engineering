<?php 

namespace App\Controller;

use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

class VinylController
{
  #[Route("/")]
  public function homepage() 
  {
    return new Response("Response from controller");
  }
}

