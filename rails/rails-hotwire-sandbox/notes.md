# Hotwire
Hotwire is an alternative approach to building modern web applications without using much JavaScript by sending HTML instead of JSON over the wire.

This makes for fast first-load  pages, keeps development experience in any programmign language, without sacrificing any of the speed or responsiveness associated with a traditional single-page application.

## Turbo
The heart of Hotwire is Turbo.
A set of complementary techniques for speeding up page changes and form submissions, dividing complex pages into components, and stream partial page updates over WebSocket. All without writting any JavaScript at all. And designed from the start to integrate perfectly with native hybrid applications for iOS and Android.

## Stimulus
While Turbo usually takes care of at least 80% of the interactivity that traditionally would have required JavaScript, there are still cases where a dash of custom code is required. Stimulus makes this easy with a HTML-centric approach to state and wiring.

## Strada 
Standardizes the way that web and native parts of a mobile hybrid application talk to each other via HTML bridge attributes. This makes it easy to progressively level-up web interactions with native replacements.

