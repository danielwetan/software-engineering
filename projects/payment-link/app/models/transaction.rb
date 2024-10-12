class Transaction < ApplicationRecord
  belongs_to :product
  has_many :payments, foreign_key: "transaction_id"  # Keeps the original naming
end
