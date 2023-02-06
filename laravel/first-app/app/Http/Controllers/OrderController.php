<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Events\PodcastProcessed;

class OrderController extends Controller
{
    public function index() {

        PodcastProcessed::dispatch("something happen");
        return "---\n";
    }
}
