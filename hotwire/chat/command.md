# Generate new Rails app with name 'chat'

rails new chat --skip-javascript

Install gem bundle

```
bundle
```

Install hotwire

```
rails hotwire:install
```

Scaffold room
rails generate scaffold room name:string

rails g model message room:references


rails g migration add_content_to_messages content:text