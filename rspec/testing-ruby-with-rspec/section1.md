What is RSpec?
- RSpec is an open-source Ruby testing library first released in 2005
- RSpec is the most popular Ruby Gem with over 300 million downloads
- The latest version, RSpec 3, was released in June 2014
- RSpec is known as DSL (a domain-specific language)
  - It constructs (i.e its classes and methods) are designed for use in a specific domain (testing)

Why test code?
- Avoid regressions when new features are added
- Provides living documentation for the codebase
- Isolate specific problems or bugs
- Improve the quality of code, especially design
- Cut down developer time

TDD Test Before writing code

The RSpec Ecosystem
- RSpec consists of 3 independent Ruby Gems that are often paired together
  - 'rspec-core' is the base library that organizes and runs the tests
  - 'rspec-expectations' is the matcher library that implements assertions
  - 'rspec-mocks' is a library used to fake behaviour of classes and objects
- RSpec permits integration with other libraries for expectations and mocks
- The 'rspec-rails' Gem integrates RSpec with the Ruby on Rails web framework

Project Structure
- Most projects have a 'spec' directory to house all RSpec test files
- The nested directories inside 'spec' mimic those of the lib folder (Ruby projects) or app folder (Rails projects)
- An RSpec file ends with a '_spec.rb' extension that matches the file it is testing
- Example: A Knight class in 'knight.rb' should have a 'knight_spec.rb' file

---
Test Pyramid
   /                     \
  / E2E Test (UI Testing) \
 / Integration Test        \
/ Unit Test                 \

Types of Tests
- Unit tests focus on individual units of code in the program (a single class, module, object, or method)
  - Elements are tested in isolation; the program is not tested well
  - The specs are easy to write and run fast
- End-to-end (E2E) or acceptance
  - Test focus on a feature and its interaction with the entire system
    - Elements are tested together; the complete program is tested with a good deal of confidence
    - The specs are hard to write, brittle and run slow

---
Test-Driven Development (TDD)
Write the test before write the code
          Red
        /    \
Refactor  --  Green
- Red, write failed test
- Green, make the test pass
- Refactor, make code cleaner



---
- Memoization using Rspec let
- Lazy loading, load when needed

The let method object will persist and be cached within the SAME example instead of example group. RSpec will evaluate the let block again to ensure isolation between test.

Lazy loaded
-> Load resource only when needed

Memoization
-> Caching the result of a computationally expensive operation to reuse it later