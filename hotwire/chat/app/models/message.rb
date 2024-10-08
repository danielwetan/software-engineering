class Message < ApplicationRecord
  belongs_to :room

  # we can replace below callback to
  broadcasts_to :room

  # after_create_commit -> { broadcast_append_to room }
  # after_destroy_commit -> { broadcast_remove_to room }
  # after_update_commit -> { broadcast_replace_to room }
end


