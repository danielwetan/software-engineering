<?php 

// Command is behavioral design pattern that converts requests or simple operations into objects.
// The conversion allows deferred or remote execution of commands, storing command history, etc.

// Reference
// https://refactoring.guru/design-patterns/command/php/example#example-1

// The Command pattern is pretty common in PHP code. It’s used for queueing tasks, tracking the history of executed tasks and performing the “undo”.

// In this example, the Command pattern is used to queue web scraping calls to the IMDB website and execute them one by one. The queue itself is kept in a database that helps to preserve commands between script launches.

/**
 * The Command interface declares the main execution method as well as several
 * helper methods for retrieving a command's metadata.
 */
interface Command
{
  public function execute(): void;

  public function getId(): int;

  public function getStatus(): int;
}


abstract class WebScrappingCommand implements Command 
{
  public $id; 

  public $status = 0;

  /**
   * @var string URL for scrapping.
   */
  public $url;

  public function __construct(string $url) 
  {
    $this->url = $url;
  }

  public function getId(): int 
  {
    return $this->id;
  }

  public function getStatus(): int 
  {
    return $this->status;
  }

  public function getURL(): string 
  {
    return $this->url;
  }

  /**
   * Since the execution methods for all web scraping commands are very
   * similar, we can provide a default implementation and let subclasses
   * override them if needed.
   *
   * Psst! An observant reader may spot another behavioral pattern in action
   * here.
   */
  public function execute(): void 
  {
    $html = $this->download();
    $this->parse($html);
    $this->complete();
  }

  public function download(): string 
  {
    $html = file_get_contents($this->getURL());
    echo "WebScrappingCommand: Downloaded {$this->url}\n";

    return $html;
  }

  abstract public function parse(string $html): void;

  public function complete(): void 
  {
    $this->status = 1;
    Queue::get()->completeCommand($this);
  }
}

