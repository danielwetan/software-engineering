class Message < ApplicationRecord
  belongs_to :room

  after_create_commit -> { broadcast_append_to room }
  after_destroy_commit -> { broadcast_remove_to room }
  after_update_commit -> { broadcast_replace_to room }

  # we can replace about hooks to 
  # broadcast_to :room 

  # Reference
  # https://github.com/hotwired/turbo-rails/blob/main/app/models/concerns/turbo/broadcastable.rb
end
