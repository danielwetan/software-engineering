class Room < ApplicationRecord
  has_many :messages

  # since room identifies the Stream, we could just use broadcasts 
  broadcasts
end

